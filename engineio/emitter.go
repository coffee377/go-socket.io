package engineio

import (
	"reflect"
	"sync"
)

// Listener 定义事件监听器接口
type Listener interface {
	call(args ...any)
}

// onceListener 定义一次性事件监听器结构体
type onceListener struct {
	emitter Emitter
	event   string
	fn      Listener
}

func (once onceListener) call(args ...any) {
	once.emitter.OffListener(once.event, once)
	once.fn.call(args)
}

type funcListener func(args ...any)

func (f funcListener) call(args ...any) {
	f(args...)
}

func WithListener(fn func(args ...any)) Listener {
	var fl funcListener = func(args ...any) {
		fn(args...)
	}
	return &fl
}

type Emitter interface {
	On(event string, listener Listener) Emitter
	Once(event string, listener Listener) Emitter
	OffListener(event string, listener Listener) Emitter
	Off(event string) Emitter
	OffAll() Emitter
	Emit(event string, args ...any) Emitter
	Listeners(event string) []Listener
	HasListeners(event string) bool
}

type innerEmitter struct {
	callbacks map[string][]Listener
	mutex     sync.Mutex
}

// NewEmitter 创建一个新的 innerEmitter 实例
func NewEmitter() Emitter {
	return &innerEmitter{
		callbacks: make(map[string][]Listener),
	}
}

func (e *innerEmitter) On(event string, listener Listener) Emitter {
	if e.mutex.TryLock() {
		defer e.mutex.Unlock()
	}
	if _, ok := e.callbacks[event]; !ok {
		e.callbacks[event] = []Listener{}
	}
	e.callbacks[event] = append(e.callbacks[event], listener)
	return e
}

func (e *innerEmitter) Once(event string, listener Listener) Emitter {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.On(event, onceListener{
		emitter: e,
		event:   event,
		fn:      listener,
	})
	return e
}

func (e *innerEmitter) OffListener(event string, listener Listener) Emitter {
	if e.mutex.TryLock() {
		defer e.mutex.Unlock()
	}
	listeners, ok := e.callbacks[event]
	if ok {
		for i, call := range listeners {
			if sameAs(listener, call) {
				listeners = append(listeners[:i], listeners[i+1:]...)
				e.callbacks[event] = listeners
				break
			}
		}
	}
	return e
}

func (e *innerEmitter) Off(event string) Emitter {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	delete(e.callbacks, event)
	return e
}

func (e *innerEmitter) OffAll() Emitter {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.callbacks = make(map[string][]Listener)
	return e
}

func (e *innerEmitter) Emit(event string, args ...any) Emitter {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if listeners, ok := e.callbacks[event]; ok {
		for _, listener := range listeners {
			listener.call(args...)
		}
	}
	return e
}

func (e *innerEmitter) Listeners(event string) []Listener {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	var res []Listener
	if listeners, ok := e.callbacks[event]; ok {
		for _, listener := range listeners {
			res = append(res, listener)
		}
	}
	return res
}

func (e *innerEmitter) HasListeners(event string) bool {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if listeners, ok := e.callbacks[event]; ok {
		return len(listeners) > 0
	}
	return false
}

func sameAs(fn Listener, internal Listener) bool {
	return reflect.DeepEqual(fn, internal)
}

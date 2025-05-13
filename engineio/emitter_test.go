package engineio

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleListener(t *testing.T) {
	emitter := NewEmitter()
	emitter.On("event", WithListener(func(args ...any) {
		assert.Equal(t, 2, len(args))
		assert.Equal(t, "Engine.IO", args[0])
		assert.Equal(t, 1, args[1])
	}))
	emitter.Emit("event", "Engine.IO", 1)
	assert.True(t, emitter.HasListeners("event"))
}

func TestMultipleListenersSameEvent(t *testing.T) {
	emitter := NewEmitter()
	var result []string
	listeners := [3]Listener{}
	for i := 0; i < len(listeners); i++ {
		listeners[i] = WithListener(func(args ...any) {
			result = append(result, fmt.Sprintf("event %d %v", i+1, args[0]))
		})
		emitter.On("event", listeners[i])
	}
	emitter.Emit("event", "Engine.IO")
	assert.True(t, emitter.HasListeners("event"))
	assert.Equal(t, 3, len(emitter.Listeners("event")))
}

func TestMultipleListenersDifferentEvents(t *testing.T) {
	emitter := NewEmitter()
	listenerExecuteCounts := make(map[string]int)
	listeners := [4]Listener{}
	for i := 0; i < len(listeners); i++ {
		key := fmt.Sprintf("listener%d", i)
		listenerExecuteCounts[key] = 0
		listeners[i] = WithListener(func(args ...any) {
			v := listenerExecuteCounts[key]
			v++
			listenerExecuteCounts[key] = v
		})
	}

	emitter.On("event0", listeners[0])
	emitter.On("event1", listeners[0])
	emitter.On("event2", listeners[0])
	emitter.On("event1", listeners[1])
	emitter.On("event2", listeners[2])

	emitter.Emit("event0")
	emitter.Emit("event1")
	emitter.Emit("event2")
	emitter.Emit("event3")

	assert.Equal(t, 3, listenerExecuteCounts["listener0"])
	assert.Equal(t, 1, listenerExecuteCounts["listener1"])
	assert.Equal(t, 1, listenerExecuteCounts["listener2"])
	assert.Equal(t, 0, listenerExecuteCounts["listener3"])
}

func TestOnceListener(t *testing.T) {
	emitter := NewEmitter()
	listenerExecuteCounts := map[string]int{
		"listener1": 0,
		"listener2": 0,
	}
	listener1 := WithListener(func(args ...any) {
		v := listenerExecuteCounts["listener1"]
		v++
		listenerExecuteCounts["listener1"] = v
	})
	listener2 := WithListener(func(args ...any) {
		v := listenerExecuteCounts["listener2"]
		v++
		listenerExecuteCounts["listener2"] = v
	})

	emitter.Once("event", listener1)
	emitter.On("event", listener2)

	emitter.Emit("event")
	emitter.Emit("event")

	assert.Equal(t, 1, listenerExecuteCounts["listener1"])
	assert.Equal(t, 2, listenerExecuteCounts["listener2"])
}

func TestOffAll(t *testing.T) {
	emitter := NewEmitter()
	counter := 0
	listener := WithListener(func(args ...any) {
		counter++
	})

	emitter.On("event0", listener)
	emitter.On("event1", listener)

	emitter.Emit("event0")
	emitter.Emit("event1")
	assert.Equal(t, 2, counter)

	emitter.OffAll()
	emitter.Emit("event0")
	emitter.Emit("event1")
	assert.Equal(t, 2, counter)
}

func TestEventOff(t *testing.T) {
	emitter := NewEmitter()
	counter := 0
	listener := WithListener(func(args ...any) {
		counter++
	})

	emitter.On("event0", listener)
	emitter.On("event1", listener)

	assert.True(t, emitter.HasListeners("event0"))
	assert.True(t, emitter.HasListeners("event1"))

	emitter.Emit("event0")
	emitter.Emit("event1")

	assert.Equal(t, 2, counter)

	emitter.Off("event0")

	assert.False(t, emitter.HasListeners("event0"))
	assert.True(t, emitter.HasListeners("event1"))

	emitter.Emit("event0")
	emitter.Emit("event1")

	assert.Equal(t, 3, counter)
}

func TestListenerOff(t *testing.T) {
	emitter := NewEmitter()
	counter1 := 0
	counter2 := 0
	listener1 := WithListener(func(args ...any) {
		counter1++
	})
	listener2 := WithListener(func(args ...any) {
		counter2++
	})

	emitter.On("event1", listener1)
	emitter.On("event2", listener1)
	emitter.On("event1", listener2)
	emitter.On("event2", listener2)

	emitter.Emit("event1")
	emitter.Emit("event2")

	assert.Equal(t, 2, counter1)
	assert.Equal(t, 2, counter2)

	emitter.OffListener("event1", listener2)
	emitter.Emit("event1")
	emitter.Emit("event2")

	assert.Equal(t, 4, counter1)
	assert.Equal(t, 3, counter2)
}

func TestListenerError(t *testing.T) {
	emitter := NewEmitter()
	counter1 := 0
	counter2 := 0
	listener1 := WithListener(func(args ...any) {
		counter1++
	})
	listener2 := WithListener(func(args ...any) {
		counter2++
	})

	emitter.On("event1", listener1)
	emitter.On("event2", listener1)
	emitter.On("event1", listener2)
	emitter.On("event2", listener2)

	emitter.Emit("event1")
	emitter.Emit("event2")

	assert.Equal(t, 2, counter1)
	assert.Equal(t, 2, counter2)
}

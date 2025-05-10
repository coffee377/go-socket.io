package main

import (
	"go-socket.io/engineio"
	"go-socket.io/logger"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
)

// Channel 模拟Netty的Channel
type Channel interface {
	Read() <-chan []byte
	Write(data []byte)
	Close() error
}

// DefaultChannel 实现Channel接口
type DefaultChannel struct {
	conn     net.Conn
	inBound  chan []byte
	outBound chan []byte
	closed   chan struct{}
	once     sync.Once
}

// NewChannel 创建新的Channel
func NewChannel(conn net.Conn) Channel {
	ch := &DefaultChannel{
		conn:     conn,
		inBound:  make(chan []byte, 100),
		outBound: make(chan []byte, 100),
		closed:   make(chan struct{}),
	}
	go ch.readLoop()
	go ch.writeLoop()
	return ch
}

// Read 从inBound通道读取数据
func (ch *DefaultChannel) Read() <-chan []byte {
	return ch.inBound
}

// Write 向outBound通道写入数据
func (ch *DefaultChannel) Write(data []byte) {
	select {
	case ch.outBound <- data:
	case <-ch.closed:
	}
}

// Close 关闭Channel
func (ch *DefaultChannel) Close() error {
	var err error
	ch.once.Do(func() {
		close(ch.closed)
		err = ch.conn.Close()
	})
	return err
}

// readLoop 循环读取网络数据并发送到inBound通道
func (ch *DefaultChannel) readLoop() {
	defer ch.Close()
	buf := make([]byte, 4096)
	for {
		n, err := ch.conn.Read(buf)
		if n > 0 {
			data := make([]byte, n)
			copy(data, buf[:n])
			ch.inBound <- data
		}
		if err != nil {
			if err != io.EOF {
				log.Printf("read error: %v", err)
			}
			break
		}
	}
}

// writeLoop 循环从outBound通道读取数据并写入网络
func (ch *DefaultChannel) writeLoop() {
	defer ch.Close()
	for {
		select {
		case data, ok := <-ch.outBound:
			if !ok {
				return
			}
			_, err := ch.conn.Write(data)
			if err != nil {
				log.Printf("write error: %v", err)
				return
			}
		case <-ch.closed:
			return
		}
	}
}

// ChannelHandler 模拟Netty的ChannelHandler
type ChannelHandler interface {
	ChannelRead(ctx ChannelHandlerContext, data []byte)
	Write(ctx ChannelHandlerContext, data []byte)
}

// ChannelHandlerContext 模拟Netty的ChannelHandlerContext
type ChannelHandlerContext interface {
	Channel() Channel
	FireChannelRead(data []byte)
	Write(data []byte)
	Next() ChannelHandlerContext
}

// DefaultChannelHandlerContext 实现ChannelHandlerContext接口
type DefaultChannelHandlerContext struct {
	channel  Channel
	index    int
	handlers []ChannelHandler
}

// Channel 获得Channel
func (ctx *DefaultChannelHandlerContext) Channel() Channel {
	return ctx.channel
}

// FireChannelRead 触发读事件
func (ctx *DefaultChannelHandlerContext) FireChannelRead(data []byte) {
	if ctx.index < len(ctx.handlers) {
		ctx.handlers[ctx.index].ChannelRead(ctx, data)
	}
}

// Write 写入数据
func (ctx *DefaultChannelHandlerContext) Write(data []byte) {
	ctx.channel.Write(data)
}

// Next 调用下一个处理器
func (ctx *DefaultChannelHandlerContext) Next() ChannelHandlerContext {
	return &DefaultChannelHandlerContext{
		channel:  ctx.channel,
		index:    ctx.index + 1,
		handlers: ctx.handlers,
	}
}

// ChannelPipeline 模拟Netty的ChannelPipeline
type ChannelPipeline interface {
	AddLast(handler ChannelHandler)
	FireChannelRead(data []byte)
}

// DefaultChannelPipeline 实现ChannelPipeline接口
type DefaultChannelPipeline struct {
	handlers []ChannelHandler
	channel  Channel
}

// AddLast 添加处理器到Pipeline
func (pipeline *DefaultChannelPipeline) AddLast(handler ChannelHandler) {
	pipeline.handlers = append(pipeline.handlers, handler)
}

// FireChannelRead 触发读事件
func (pipeline *DefaultChannelPipeline) FireChannelRead(data []byte) {
	ctx := &DefaultChannelHandlerContext{
		channel:  pipeline.channel,
		index:    0,
		handlers: pipeline.handlers,
	}
	ctx.FireChannelRead(data)
}

// EchoHandler 模拟Netty的EchoHandler
type EchoHandler struct{}

// ChannelRead 处理读事件
func (h *EchoHandler) ChannelRead(ctx ChannelHandlerContext, data []byte) {
	ctx.Write(data)
	ctx.Next().FireChannelRead(data)
}

// Write 处理写事件
func (h *EchoHandler) Write(ctx ChannelHandlerContext, data []byte) {
	ctx.Write(data)
}

func main() {
	s := engineio.NewServer(nil)
	defer s.Shutdown()
	http.Handle("/socket.io/", s)
	http.Handle("/", http.FileServer(http.Dir("./asset")))

	logger.Info("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

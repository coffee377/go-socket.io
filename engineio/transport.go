package engineio

import (
	"go-socket.io/engineio/parser"
	"go-socket.io/engineio/protocol"
	"net/http"
)

type EmitterTransport interface {
	Parser() parser.Parser
	OnError(reason, description string)
	OnPacket(packet protocol.EnginePacket)
	OnData(data any)
	OnClose()
	Emitter
}

type Transport interface {
	Name() string
	InitialQuery() map[string]string
	InitialHeaders() http.Header
	Send(packets []protocol.EnginePacket)
	isWritable() bool
	Close()
	EmitterTransport
	http.Handler
}

type emitterTransport struct {
	parser     parser.Parser
	readyState ReadyState
	Emitter
}

func newTransport(parser parser.Parser) EmitterTransport {
	return &emitterTransport{
		parser:     parser,
		readyState: ReadyStateOpen,
		Emitter:    NewEmitter(),
	}
}

func (t *emitterTransport) Parser() parser.Parser {
	return t.parser
}

func (t *emitterTransport) Close() {
	if t.readyState != ReadyStateClosed && t.readyState != ReadyStateClosing {
		t.readyState = ReadyStateClosing
		//doClose();
	}
}

func (t *emitterTransport) OnError(reason, description string) {
	if t.HasListeners("error") {
		t.Emit("error", reason, description)
	}
}

func (t *emitterTransport) OnPacket(packet protocol.EnginePacket) {
	t.Emit("packet", packet)
}

func (t *emitterTransport) OnData(data any) {
	packet := t.parser.DecodePacket(data)
	t.OnPacket(packet)
}

func (t *emitterTransport) OnClose() {
	t.readyState = ReadyStateClosed
	t.Emit("close")
}

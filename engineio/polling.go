package engineio

import (
	"go-socket.io/engineio/protocol"
	"net/http"
	"strings"
	"sync"
)

const (
	NAME               = "polling"
	CLOSE_PACKET       = "close"
	NOOP_PACKET        = "noop"
	CONTENT_TYPE       = "text/plain; charset=UTF-8"
	JSONP_CONTENT_TYPE = "text/javascript; charset=UTF-8"
)

type Polling struct {
	EmitterTransport

	locker sync.Mutex

	writable    bool
	shouldClose bool
	query       map[string]string
	headers     http.Header
}

// Default is the default transport.
var Default = &Polling{
	writable:         false,
	shouldClose:      false,
	query:            make(map[string]string),
	EmitterTransport: &emitterTransport{},
}

func (p *Polling) Name() string {
	return NAME
}

func (p *Polling) InitialQuery() map[string]string {
	return p.query
}

func (p *Polling) InitialHeaders() http.Header {
	return p.headers
}

func (p *Polling) Send(packets []protocol.EnginePacket) {

}

func (p *Polling) isWritable() bool {
	return p.writable
}

func (p *Polling) Close() {
	p.locker.Lock()
	defer p.locker.Unlock()

	if p.writable {
		p.Send([]protocol.EnginePacket{protocol.Packet(protocol.WithPacketType(protocol.PacketClose))})
		p.OnClose()
	} else {
		p.shouldClose = true
	}
}

func (p *Polling) OnClose() {
	if p.writable {
		p.Send([]protocol.EnginePacket{protocol.Packet(protocol.WithPacketType(protocol.PacketNoop))})
	}
	p.EmitterTransport.OnClose()
}

func (p *Polling) OnData(data any) {
	packets := p.Parser().DecodePayload(data)
	for _, packet := range packets {
		if packet.GetType() == protocol.PacketClose {
			p.OnClose()
		} else {
			p.OnPacket(packet)
		}
	}
}

func (p *Polling) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	p.locker.Lock()
	defer p.locker.Unlock()
	if len(p.query) == 0 {
		for key, values := range request.URL.Query() {
			p.query[key] = values[0]
		}
	}
	if p.headers == nil {
		p.headers = request.Header
	}

	switch strings.ToUpper(request.Method) {
	case "GET":
		p.onPollRequest(writer, request)
	case "POST":
		p.onDataRequest(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (p *Polling) onPollRequest(writer http.ResponseWriter, request *http.Request) {

}

func (p *Polling) onDataRequest(writer http.ResponseWriter, request *http.Request) {

}

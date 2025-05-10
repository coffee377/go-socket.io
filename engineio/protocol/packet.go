package protocol

import "encoding/base64"

type EnginePacket interface {
	GetType() PacketType
	GetPayload() any
	PacketType(packetType PacketType) EnginePacket
	Payload(payload any) EnginePacket
}

type packet struct {
	packetType PacketType
	payload    any
}

func (p *packet) GetType() PacketType {
	return p.packetType
}

func (p *packet) GetPayload() any {
	return p.payload
}

func (p *packet) PacketType(packetType PacketType) EnginePacket {
	p.packetType = packetType
	return p
}

func (p *packet) Payload(payload any) EnginePacket {
	p.payload = payload
	return p
}

type PacketOption func(*packet)

func Packet(opts ...PacketOption) EnginePacket {
	p := &packet{
		packetType: PacketMessage,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func WithPacketType(packetType PacketType) PacketOption {
	return func(p *packet) {
		p.packetType = packetType
	}
}

func WithPayload(data any) PacketOption {
	return func(p *packet) {
		p.payload = data
	}
}

func FromBytes(bytes []byte, payloadType PayloadType) PacketOption {
	var (
		packetType = PacketMessage
		data       []byte
	)

	// base64 编码的二进制数据
	if bytes[0] == 'b' {
		data, _ = base64.StdEncoding.DecodeString(string(bytes[1:]))
	} else if bytes[0] >= PacketOpen.Byte() && bytes[0] <= PacketNoop.Byte() {
		packetType = ByteToPacketType(bytes[0])
		payload := bytes[1:]
		if len(payload) > 0 {
			data = payload
		}
	} else {
		if len(bytes) > 0 {
			data = bytes
		}
	}
	return func(p *packet) {
		p.packetType = packetType
		if payloadType == PayloadBinary {
			p.payload = data
		} else {
			p.payload = string(data)
		}
	}
}

package parser

import "go-socket.io/engineio/packet"

type EncodeCallback[T string | []byte | any] = func(data T)

type DecodePayloadCallback[T string | []byte | any] = func(packet packet.Packet[T], index, total int) bool

type Parser interface {
	GetProtocolVersion() int

	EncodePacket(pack packet.Packet[any], supportsBinary bool, callback EncodeCallback[any])

	EncodePayload(packets []packet.Packet[any], supportsBinary bool, callback EncodeCallback[any])

	DecodePacket(data any) packet.Packet[any]

	DecodePayload(data any, callback DecodePayloadCallback[any])
}

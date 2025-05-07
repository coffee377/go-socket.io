package parser

import (
	"encoding/base64"
	"fmt"
	"go-socket.io/engineio/frame"
	"go-socket.io/engineio/packet"
	"strings"
)

const (
	RecordSeparator = "\x1e"
)

var ProtocolV4 = v4{}

type v4 struct {
}

func (p v4) GetProtocolVersion() int {
	return 4
}

func (p v4) EncodePacket(pack packet.Packet[any], supportsBinary bool, callback EncodeCallback[any]) {
	switch data := pack.Data.(type) {
	case []byte:
		p.encodeByteArray(packet.Packet[[]byte]{Type: pack.Type, Data: data}, supportsBinary, callback)
	default:
		encoded := fmt.Sprintf("%v", pack.Type.BinaryByte())
		if pack.Data != nil {
			encoded += fmt.Sprintf("%v", pack.Data)
		}
		callback(encoded)
	}
}

func (p v4) EncodePayload(packets []packet.Packet[any], supportsBinary bool, callback EncodeCallback[any]) {
	encodedPackets := make([]string, len(packets))
	for i, pack := range packets {
		// force base64 encoding for binary packets
		p.EncodePacket(pack, false, func(data interface{}) {
			encodedPackets[i] = data.(string)
		})
	}
	callback(strings.Join(encodedPackets, RecordSeparator))
}

func (p v4) DecodePacket(data any) packet.Packet[any] {
	result := packet.Packet[any]{Type: packet.MESSAGE}
	if data == nil {
		result.Data = data
		return result
	}
	switch d := data.(type) {
	case string:
		if d[0] == 'b' {
			decodeString, _ := base64.StdEncoding.DecodeString(d[1:])
			result.Data = string(decodeString)
		} else {
			result.Type = packet.ByteToPacketType(d[0], frame.String)
			result.Data = d[1:]
		}
	case []byte:
		result.Data = data
	default:
		panic("Invalid type for data: " + fmt.Sprintf("%T", data))
	}
	return result
}

func (p v4) DecodePayload(data any, callback DecodePayloadCallback[any]) {
	packets := make([]packet.Packet[any], 0)
	if strData, ok := data.(string); ok {
		encodedPackets := strings.Split(strData, RecordSeparator)
		for _, encodedPacket := range encodedPackets {
			pack := p.DecodePacket(encodedPacket)
			packets = append(packets, pack)
			//if packet.Type == ERROR {
			//	break
			//}
		}
	} else {
		panic("data must be a string")
	}

	for i, pack := range packets {
		if !callback(pack, i, len(packets)) {
			return
		}
	}
}

func (p v4) encodeByteArray(pack packet.Packet[[]byte], binary bool, callback EncodeCallback[any]) {
	if binary {
		callback(pack.Data)
	} else {
		callback(fmt.Sprintf("b%s", base64.StdEncoding.EncodeToString(pack.Data)))
	}
}

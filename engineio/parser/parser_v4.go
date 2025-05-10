package parser

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-socket.io/engineio/protocol"
)

const (
	V4RecordSeparator     = 0x1e
	RecordSeparatorString = "\x1e"
)

var ProtocolV4 = v4{
	buf: &bytes.Buffer{},
}

type v4 struct {
	buf *bytes.Buffer
}

func (p v4) GetProtocolVersion() int {
	return protocol.V4.Value()
}

func (p v4) EncodePacket(packet protocol.EnginePacket, supportsBinary bool) []byte {
	encode(p.buf, packet, supportsBinary)
	result := p.buf.Bytes()
	p.buf.Reset()
	return result
}

func (p v4) EncodePayload(packets []protocol.EnginePacket, supportsBinary bool) bytes.Buffer {
	buf := bytes.Buffer{}
	for i, packet := range packets {
		// v4 force base64 encoding for binary packets
		// supportsBinary = false
		encode(&buf, packet, supportsBinary)
		if i < len(packets)-1 {
			buf.WriteByte(V4RecordSeparator)
		}
	}
	return buf
}

func (p v4) DecodePacket(data any, opts ...DecodeOptions) protocol.EnginePacket {
	if data == nil {
		return nil
	}

	var byteData []byte
	switch d := data.(type) {
	case string:
		byteData = []byte(d)
	case []byte:
		byteData = d
	default:
		panic("Invalid type for data: " + fmt.Sprintf("%T", data))
	}
	decodeOpt := &decodeOpts{
		payloadType: protocol.PayloadBinary,
	}
	for _, opt := range opts {
		opt(decodeOpt)
	}
	packet := protocol.Packet(protocol.FromBytes(byteData, decodeOpt.payloadType))
	return packet
}

func (p v4) DecodePayload(data any, opts ...DecodeOptions) []protocol.EnginePacket {
	var bytesPackets [][]byte
	switch data.(type) {
	case string:
		bytesPackets = splitData([]byte(data.(string)), V4RecordSeparator)
	case []byte:
		bytesPackets = splitData(data.([]byte), V4RecordSeparator)
	default:
		panic("data must be a string or []byte")
	}

	packets := make([]protocol.EnginePacket, 0)
	for _, bytesPacket := range bytesPackets {
		packet := p.DecodePacket(bytesPacket, opts...)
		packets = append(packets, packet)
	}

	return packets
}

func splitData(data []byte, sep byte) [][]byte {
	// 计算分隔符数量
	count := 0
	for _, b := range data {
		if b == sep {
			count++
		}
	}

	// 预分配结果切片
	result := make([][]byte, 0, count+1)
	start := 0

	for i, b := range data {
		if b == sep {
			if len(data[start:i]) > 0 {
				result = append(result, data[start:i])
			}
			start = i + 1
		}
	}

	return result
}

func encode(buf *bytes.Buffer, packet protocol.EnginePacket, supportsBinary bool) {
	switch data := packet.GetPayload().(type) {
	case string:
		buf.WriteByte(packet.GetType().StringByte())
		buf.WriteString(data)
	case []byte:
		if supportsBinary {
			buf.Write(data)
		} else {
			buf.WriteByte('b')
			buf.WriteString(base64.StdEncoding.EncodeToString(data))
		}
	default:
		buf.WriteByte(packet.GetType().StringByte())
		if packet.GetPayload() != nil {
			err := json.NewEncoder(buf).Encode(packet.GetPayload())
			if err != nil {
				panic(err)
			}
		}
	}
}

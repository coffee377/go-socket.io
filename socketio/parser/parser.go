package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-socket.io/socketio"
	"reflect"
)

type Parser interface {
}

type EncodeCallback = func(data []any)
type DecodeCallback = func(packet socketio.Packet)

type Encoder interface {
	Encode(packet socketio.Packet) (bytes.Buffer, [][]byte)
}

type Decoder[T string | []byte] interface {
	Add(obj T)
	Destroy()
	OnDecoded(callback DecodeCallback)
}

var IOParser = ioParser{}

type ioParser struct {
}

const (
	bufferTypeName = "Buffer"
)

func (p ioParser) Encode(packet socketio.Packet) (bytes.Buffer, [][]byte) {
	//num := uint64(0)
	// 获取二进制数据
	_ = packet.Deconstruct()
	buffers := packet.GetBuffers()
	//buffers, _ := p.attachBuffer(reflect.ValueOf(packet.Data), &num)
	//packet.Attachments = len(buffers)

	if len(buffers) > 0 && (packet.Type == socketio.Event || packet.Type == socketio.Ack) {
		packet.Type += 3
	}

	if packet.Type == socketio.BinaryEvent || packet.Type == socketio.BinaryAck {
		return p.encodeAsBinary(packet)
	}
	encoding := p.encodeAsString(packet)
	return encoding, buffers
}

// <packet type>[<# of binary attachments>-][<namespace>,][<acknowledgment id>][JSON-stringified payload without binary]
//
// + binary attachments extracted
func (p ioParser) encodeAsString(packet socketio.Packet) bytes.Buffer {
	str := bytes.Buffer{}
	// 1. 写入数据包类型
	//str.WriteByte(byte(packet.Type + '0'))
	str.WriteString(fmt.Sprintf("%d", packet.Type))

	// 2. 写入二进制数据包数量
	if packet.Type == socketio.BinaryEvent || packet.Type == socketio.BinaryAck {
		//str.WriteByte(byte(packet.Attachments + '0'))
		//str.WriteByte('-')
		str.WriteString(fmt.Sprintf("%d-", packet.Attachments))
	}

	// 3. 写入命名空间
	if packet.Namespace != "" && packet.Namespace != "/" {
		str.WriteString(fmt.Sprintf("%s,", packet.Namespace))
	}

	// 4. 写入数据包id
	if packet.Id > 0 {
		str.WriteString(fmt.Sprintf("%d", packet.Id))
	}

	// 5. 写入数据包数据
	if packet.Data != nil {
		res, err := json.Marshal(packet.Data)
		if err != nil {
			fmt.Printf("failed to marshal data: %v\n", err)
		}
		if string(res) != "null" {
			str.Write(res)
		}
	}

	fmt.Printf("encoded %s => %s\n", packet, str.String())
	return str
}

func (p ioParser) encodeAsBinary(packet socketio.Packet) (bytes.Buffer, [][]byte) {
	return bytes.Buffer{}, packet.GetBuffers()
}

func (p ioParser) hasBinary(packet socketio.Packet) bool {
	return false
}

func (p ioParser) attachBuffer(v reflect.Value, index *uint64) ([][]byte, error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	var result [][]byte
	switch v.Kind() {
	case reflect.Struct:
		if v.Type().Name() == bufferTypeName {
			if !v.CanAddr() {
				return nil, errFailedBufferAddress
			}
			buffer := v.Addr().Interface().(*Buffer)
			buffer.num = *index
			buffer.isBinary = true
			result = append(result, buffer.Data)
			*index++
		} else {
			for i := 0; i < v.NumField(); i++ {
				b, err := p.attachBuffer(v.Field(i), index)
				if err != nil {
					return nil, err
				}
				result = append(result, b...)
			}
		}

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			b, err := p.attachBuffer(v.Index(i), index)
			if err != nil {
				return nil, err
			}

			result = append(result, b...)
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			b, err := p.attachBuffer(v.MapIndex(key), index)
			if err != nil {
				return nil, err
			}

			result = append(result, b...)
		}
	default:
	}

	return result, nil
}

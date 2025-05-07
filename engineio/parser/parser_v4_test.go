package parser

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-socket.io/engineio/packet"
	"testing"
)

func TestV4_GetProtocolVersion(t *testing.T) {
	assert.Equal(t, 4, ProtocolV4.GetProtocolVersion())
}

func TestEncodePacketString(t *testing.T) {
	pack := packet.Packet[any]{Type: packet.MESSAGE}

	pack.Data = "Hello World"
	ProtocolV4.EncodePacket(pack, false, func(data interface{}) {
		expected := fmt.Sprintf("%d%s", packet.MESSAGE, "Hello World")
		assert.Equal(t, expected, data)
	})

	pack.Data = "Engine.IO"
	ProtocolV4.EncodePacket(pack, false, func(data interface{}) {
		expected := fmt.Sprintf("%d%s", packet.MESSAGE, "Engine.IO")
		assert.Equal(t, expected, data)
	})
}

func TestEncodePacketBinary(t *testing.T) {
	pack := packet.Packet[any]{Type: packet.MESSAGE}
	// 72 101 108 108 111 44 32 87 111 114 108 100 33
	tData := []byte("Hello, World!")
	pack.Data = tData
	ProtocolV4.EncodePacket(pack, true, func(data interface{}) {
		assert.Equal(t, tData, data)
	})
}

func TestEncodePacketBase64(t *testing.T) {
	pack := packet.Packet[any]{Type: packet.MESSAGE}
	tData := []byte("Hello, World!")
	pack.Data = tData
	ProtocolV4.EncodePacket(pack, false, func(data interface{}) {
		expected := fmt.Sprintf("b%s", base64.StdEncoding.EncodeToString(tData))
		assert.Equal(t, expected, data)
	})
}

func TestEncodePayloadStringEmpty(t *testing.T) {
	ProtocolV4.EncodePayload([]packet.Packet[any]{}, false, func(data interface{}) {
		assert.Equal(t, "", data)
	})
}

func TestEncodePayloadString(t *testing.T) {
	packets := []packet.Packet[any]{
		{Type: packet.MESSAGE, Data: "Engine.IO"},
		{Type: packet.MESSAGE, Data: "Test.Data"},
	}
	var expected bytes.Buffer
	for i, p := range packets {
		expected.WriteByte(p.Type.StringByte())
		switch a := p.Data.(type) {
		case string:
			expected.WriteString(a)
		case []byte:
			expected.WriteString(string(a))
		}
		if i < len(packets)-1 {
			expected.WriteString(RecordSeparator)
		}
	}

	ProtocolV4.EncodePayload(packets, false, func(data interface{}) {
		s := expected.String()
		assert.Equal(t, s, data)
	})
}

func TestEncodePayloadBinary(t *testing.T) {

}

func TestEncodePayloadBinaryMixed(t *testing.T) {

}

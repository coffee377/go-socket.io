package parser

import (
	"bytes"
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"go-socket.io/engineio/frame"
	"go-socket.io/engineio/packet"
	"testing"
)

func expectedResult(supportsBinary bool, packets ...packet.Packet[any]) bytes.Buffer {
	var expected bytes.Buffer
	for i, p := range packets {
		switch d := p.Data.(type) {
		case string:
			expected.WriteByte(p.Type.StringByte())
			expected.WriteString(d)
		case []byte:
			if supportsBinary {
				expected.Write(d)
			} else {
				expected.WriteString("b")
				expected.WriteString(base64.StdEncoding.EncodeToString(d))
			}
		}
		if i < len(packets)-1 {
			expected.WriteString(RecordSeparator)
		}
	}
	return expected
}

func TestV4_GetProtocolVersion(t *testing.T) {
	assert.Equal(t, 4, ProtocolV4.GetProtocolVersion())
}

func TestV4_EncodePacketString(t *testing.T) {
	pack := packet.Packet[any]{Type: packet.MESSAGE}

	pack.Data = "Hello World"
	ProtocolV4.EncodePacket(pack, false, func(data interface{}) {
		expected := expectedResult(false, pack)
		assert.Equal(t, expected.String(), data)
	})

	pack.Data = "Engine.IO"
	ProtocolV4.EncodePacket(pack, false, func(data interface{}) {
		expected := expectedResult(false, pack)
		assert.Equal(t, expected.String(), data)
	})
}

func TestV4_EncodePacketBinary(t *testing.T) {
	// 72 101 108 108 111 44 32 87 111 114 108 100 33
	pack := packet.Packet[any]{Type: packet.MESSAGE, Data: []byte("Hello, World!")}
	ProtocolV4.EncodePacket(pack, true, func(data interface{}) {
		expected := expectedResult(true, pack)
		assert.Equal(t, expected.Bytes(), data)
	})
}

func TestV4_EncodePacketBase64(t *testing.T) {
	pack := packet.Packet[any]{Type: packet.MESSAGE, Data: []byte("Hello, World!")}
	ProtocolV4.EncodePacket(pack, false, func(data interface{}) {
		expected := expectedResult(false, pack)
		assert.Equal(t, expected.String(), data)
	})
}

func TestV4_EncodePayloadStringEmpty(t *testing.T) {
	var packets []packet.Packet[any]
	ProtocolV4.EncodePayload(packets, false, func(data interface{}) {
		expected := expectedResult(false, packets...)
		assert.Equal(t, expected.String(), data)
	})
}

func TestV4_EncodePayloadString(t *testing.T) {
	packets := []packet.Packet[any]{
		{Type: packet.MESSAGE, Data: "Engine.IO"},
		{Type: packet.MESSAGE, Data: "Test.Data"},
	}
	expected := expectedResult(false, packets...)
	ProtocolV4.EncodePayload(packets, false, func(data interface{}) {
		assert.Equal(t, expected.String(), data)
	})
}

func TestV4_EncodePayloadBinary(t *testing.T) {
	packets := []packet.Packet[any]{
		{Type: packet.MESSAGE, Data: []byte("Engine.IO")},
		{Type: packet.MESSAGE, Data: []byte("Test.Data")},
	}
	expected := expectedResult(true, packets...)

	ProtocolV4.EncodePayload(packets, true, func(data interface{}) {
		assert.Equal(t, expected.Bytes(), data)
	})
}

func TestV4_EncodePayloadBinaryMixed(t *testing.T) {
	packets := []packet.Packet[any]{
		{Type: packet.MESSAGE, Data: []byte("Engine.IO")},
		{Type: packet.MESSAGE, Data: "Test.Data"},
	}
	expected := expectedResult(false, packets...)

	ProtocolV4.EncodePayload(packets, false, func(data interface{}) {
		assert.Equal(t, expected.String(), data)
	})
}

func TestV4_EncodePayloadBase64(t *testing.T) {
	packets := []packet.Packet[any]{
		{Type: packet.MESSAGE, Data: []byte("Engine.IO")},
		{Type: packet.MESSAGE, Data: []byte("Test.Data")},
	}
	expected := expectedResult(false, packets...)

	ProtocolV4.EncodePayload(packets, false, func(data interface{}) {
		assert.Equal(t, expected.String(), data)
	})
}

func TestV4_EncodePayloadBase64Mixed(t *testing.T) {
	packets := []packet.Packet[any]{
		{Type: packet.MESSAGE, Data: []byte("Engine.IO")},
		{Type: packet.MESSAGE, Data: "Test.Data"},
	}
	expected := expectedResult(false, packets...)

	ProtocolV4.EncodePayload(packets, false, func(data interface{}) {
		assert.Equal(t, expected.String(), data)
	})
}

func TestV4_DecodePacket(t *testing.T) {
	//packetOriginal := packet.Packet[any]{Type: packet.MESSAGE, Data: "Engine.IO"}
	//ProtocolV4.EncodePacket(packetOriginal, false, func(data interface{}) {
	//	packetDecoded := ProtocolV4.DecodePacket(data)
	//	assert.Equal(t, packetOriginal.Type, packetDecoded.Type)
	//	assert.Equal(t, packetOriginal.Data, packetDecoded.Data)
	//})
	types := []string{"0", "1", "2", "3", "4", "5", "6"}
	for _, ty := range types {
		packetDecoded := ProtocolV4.DecodePacket(ty)
		assert.Equal(t, packet.ByteToPacketType(ty[0], frame.String), packetDecoded.Type)
	}
}

func TestV4_DecodePacketNull(t *testing.T) {
	decodePacket := ProtocolV4.DecodePacket(nil)
	assert.Equal(t, packet.MESSAGE, decodePacket.Type)
	assert.Nil(t, decodePacket.Data)
}

func TestV4_DecodePacketString(t *testing.T) {
	packetOriginal := packet.Packet[any]{Type: packet.MESSAGE, Data: "Engine.IO"}
	ProtocolV4.EncodePacket(packetOriginal, false, func(data interface{}) {
		packetDecoded := ProtocolV4.DecodePacket(data)
		assert.Equal(t, packetOriginal.Type, packetDecoded.Type)
		assert.Equal(t, packetOriginal.Data, packetDecoded.Data)
	})
}

func TestV4_DecodePacketBinary(t *testing.T) {
	packetOriginal := packet.Packet[any]{Type: packet.MESSAGE, Data: []byte("Engine.IO")}
	ProtocolV4.EncodePacket(packetOriginal, true, func(data interface{}) {
		packetDecoded := ProtocolV4.DecodePacket(data)
		assert.Equal(t, packetOriginal.Type, packetDecoded.Type)
		assert.Equal(t, packetOriginal.Data, packetDecoded.Data)
	})
}

func TestV4_DecodePacketBase64(t *testing.T) {
	packetOriginal := packet.Packet[any]{Type: packet.MESSAGE, Data: []byte("Engine.IO")}
	ProtocolV4.EncodePacket(packetOriginal, false, func(data interface{}) {
		packetDecoded := ProtocolV4.DecodePacket(data)
		assert.Equal(t, packetOriginal.Type, packetDecoded.Type)
		assert.Equal(t, "Engine.IO", packetDecoded.Data)
	})
}

package parser

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-socket.io/engineio/protocol"
	"testing"
)

func expectedResult(supportsBinary bool, packets ...protocol.EnginePacket) bytes.Buffer {
	var buf bytes.Buffer
	for i, p := range packets {
		encode(&buf, p, supportsBinary)
		if i < len(packets)-1 {
			buf.WriteByte(V4RecordSeparator)
		}
	}
	return buf
}

func TestV4_GetProtocolVersion(t *testing.T) {
	assert.Equal(t, protocol.V4.Value(), ProtocolV4.GetProtocolVersion())
}

func TestV4_EncodePacketString(t *testing.T) {
	packet := protocol.Packet()

	packet.Payload("Hello World")
	expected := expectedResult(false, packet)
	data := ProtocolV4.EncodePacket(packet, false)
	assert.Equal(t, expected.String(), string(data))

	packet.Payload("Engine.IO")
	data = ProtocolV4.EncodePacket(packet, false)
	expected = expectedResult(false, packet)
	assert.Equal(t, expected.String(), string(data))
}

func TestV4_EncodePacketBinary(t *testing.T) {
	// 72 101 108 108 111 44 32 87 111 114 108 100 33
	packet := protocol.Packet(protocol.WithPayload([]byte("Hello, World!")))
	encodePacket := ProtocolV4.EncodePacket(packet, true)
	expected := expectedResult(true, packet)
	assert.Equal(t, expected.Bytes(), encodePacket)
}

func TestV4_EncodePacketBase64(t *testing.T) {
	packet := protocol.Packet(protocol.WithPayload([]byte("Hello, World!")))
	encodePacket := ProtocolV4.EncodePacket(packet, false)
	expected := expectedResult(false, packet)
	assert.Equal(t, expected.Bytes(), encodePacket)
}

func TestV4_EncodePayloadStringEmpty(t *testing.T) {
	var packets []protocol.EnginePacket
	payload := ProtocolV4.EncodePayload(packets, false)
	expected := expectedResult(false, packets...)
	assert.Equal(t, expected.String(), payload.String())
	assert.Equal(t, expected.Bytes(), payload.Bytes())
}

func TestV4_EncodePayloadString(t *testing.T) {
	packets := []protocol.EnginePacket{
		protocol.Packet(protocol.WithPayload("Engine.IO")),
		protocol.Packet(protocol.WithPayload("Test.Data")),
	}
	payload := ProtocolV4.EncodePayload(packets, false)
	expected := expectedResult(false, packets...)
	assert.Equal(t, expected.String(), payload.String())
	assert.Equal(t, expected.Bytes(), payload.Bytes())
}

func TestV4_EncodePayloadBinary(t *testing.T) {
	packets := []protocol.EnginePacket{
		protocol.Packet(protocol.WithPayload([]byte("Engine.IO"))),
		protocol.Packet(protocol.WithPayload([]byte("Test.Data"))),
	}
	expected := expectedResult(true, packets...)
	payload := ProtocolV4.EncodePayload(packets, true)
	assert.Equal(t, expected.Bytes(), payload.Bytes())
	assert.Equal(t, expected.String(), payload.String())
}

func TestV4_EncodePayloadBinaryMixed(t *testing.T) {
	packets := []protocol.EnginePacket{
		protocol.Packet(protocol.WithPayload([]byte("Engine.IO"))),
		protocol.Packet(protocol.WithPayload("Test.Data")),
	}
	expected := expectedResult(true, packets...)
	payload := ProtocolV4.EncodePayload(packets, true)
	assert.Equal(t, expected.String(), payload.String())
	assert.Equal(t, expected.Bytes(), payload.Bytes())
}

func TestV4_EncodePayloadBase64(t *testing.T) {
	packets := []protocol.EnginePacket{
		protocol.Packet(protocol.WithPayload([]byte("Engine.IO"))),
		protocol.Packet(protocol.WithPayload([]byte("Test.Data"))),
	}
	expected := expectedResult(false, packets...)
	payload := ProtocolV4.EncodePayload(packets, false)
	assert.Equal(t, expected.String(), payload.String())
	assert.Equal(t, expected.Bytes(), payload.Bytes())
}

func TestV4_EncodePayloadBase64Mixed(t *testing.T) {
	packets := []protocol.EnginePacket{
		protocol.Packet(protocol.WithPayload([]byte("Engine.IO"))),
		protocol.Packet(protocol.WithPayload("Test.Data")),
	}
	expected := expectedResult(false, packets...)
	payload := ProtocolV4.EncodePayload(packets, false)
	assert.Equal(t, expected.String(), payload.String())
	assert.Equal(t, expected.Bytes(), payload.Bytes())
}

func TestV4_DecodePacketNull(t *testing.T) {
	packet := ProtocolV4.DecodePacket(nil)
	assert.Nil(t, packet)
}

func TestV4_DecodePacket(t *testing.T) {
	pts := []protocol.PacketType{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6}
	for _, pt := range pts {
		data := fmt.Sprintf("%d", pt)
		decodePacket := ProtocolV4.DecodePacket(data)
		assert.Equal(t, protocol.ByteToPacketType(pt.Byte()), decodePacket.GetType())
		p := decodePacket.GetPayload()
		if p == nil {
			assert.Nil(t, p)
			continue
		}
	}
}

func TestV4_DecodePacketString(t *testing.T) {
	packetOriginal := protocol.Packet(protocol.WithPayload("Engine.IO"))
	encoded := ProtocolV4.EncodePacket(packetOriginal, false)
	packet := ProtocolV4.DecodePacket(encoded)
	assert.Equal(t, packetOriginal.GetType(), packet.GetType())

	payload := packet.GetPayload()
	if payload == nil {
		assert.Nil(t, payload)
	}
	switch payload.(type) {
	case string:
		assert.Equal(t, packetOriginal.GetPayload(), payload.(string))
	case []byte:
		assert.Equal(t, packetOriginal.GetPayload(), string(payload.([]byte)))
	}
}

func TestV4_DecodePacketBinary(t *testing.T) {
	packetOriginal := protocol.Packet(protocol.WithPayload([]byte("Engine.IO")))
	encoded := ProtocolV4.EncodePacket(packetOriginal, false)
	packet := ProtocolV4.DecodePacket(encoded)
	assert.Equal(t, packetOriginal.GetType(), packet.GetType())

	payload := packet.GetPayload()
	if payload == nil {
		assert.Nil(t, payload)
	}
	switch payload.(type) {
	case string:
		assert.Equal(t, packetOriginal.GetPayload(), payload.(string))
	case []byte:
		assert.Equal(t, packetOriginal.GetPayload(), payload.([]byte))
	}
}

func TestV4_DecodePacketBase64(t *testing.T) {
	data := "Engine.IO"
	packetOriginal := protocol.Packet(protocol.WithPayload(data))
	encoded := ProtocolV4.EncodePacket(packetOriginal, false)
	packet := ProtocolV4.DecodePacket(encoded, WithPlaintextDecode())
	assert.Equal(t, packetOriginal.GetType(), packet.GetType())
	assert.Equal(t, data, packet.GetPayload())
}

func Test_split(t *testing.T) {
	testData := []byte{30, 1, 2, 30, 30}
	rds := splitData(testData, V4RecordSeparator)
	for i, r := range rds {
		fmt.Printf("%d => %v\n", i, r)
	}
	assert.Equal(t, []byte{1, 2}, rds[0])
	assert.Equal(t, cap(rds), 4)
	assert.Equal(t, len(rds), 1)
}

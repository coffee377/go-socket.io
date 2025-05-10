package protocol

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPacketType(t *testing.T) {
	var tests = []struct {
		b           byte
		payloadType PayloadType
		packetType  PacketType
		strByte     byte
		binByte     byte
		str         string
	}{
		{0, PayloadBinary, PacketOpen, '0', 0, "open"},
		{1, PayloadBinary, PacketClose, '1', 1, "close"},
		{2, PayloadBinary, PacketPing, '2', 2, "ping"},
		{3, PayloadBinary, PacketPong, '3', 3, "pong"},
		{4, PayloadBinary, PacketMessage, '4', 4, "message"},
		{5, PayloadBinary, PacketUpgrade, '5', 5, "upgrade"},
		{6, PayloadBinary, PacketNoop, '6', 6, "noop"},

		{'0', PayloadPlaintext, PacketOpen, '0', 0, "open"},
		{'1', PayloadPlaintext, PacketClose, '1', 1, "close"},
		{'2', PayloadPlaintext, PacketPing, '2', 2, "ping"},
		{'3', PayloadPlaintext, PacketPong, '3', 3, "pong"},
		{'4', PayloadPlaintext, PacketMessage, '4', 4, "message"},
		{'5', PayloadPlaintext, PacketUpgrade, '5', 5, "upgrade"},
		{'6', PayloadPlaintext, PacketNoop, '6', 6, "noop"},
	}

	for i, test := range tests {
		typ := ByteToPacketType(test.b, test.payloadType)

		require.Equal(t, test.packetType, typ, fmt.Sprintf(`types not equal by case: %d`, i))

		assert.Equal(t, test.strByte, typ.StringByte(), fmt.Sprintf(`string byte not equal by case: %d`, i))
		assert.Equal(t, test.binByte, typ.BinaryByte(), fmt.Sprintf(`bytes not equal by case: %d`, i))
		assert.Equal(t, test.str, typ.String(), fmt.Sprintf(`strings not equal by case: %d`, i))
	}
}

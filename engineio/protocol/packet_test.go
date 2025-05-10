package protocol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPacket(t *testing.T) {
	enginePacket := Packet()
	assert.Equal(t, PacketMessage, enginePacket.GetType())
	assert.Nil(t, enginePacket.GetPayload())
}

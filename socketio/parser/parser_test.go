package parser

import (
	"github.com/stretchr/testify/assert"
	"go-socket.io/socketio"
	"testing"
)

type testData struct {
	name      string
	data      socketio.Packet
	expected  string
	bufferLen int
}

func TestIOParser_ConnectionToNamespace(t *testing.T) {

	tests := []testData{
		{"main namespace", socketio.CreateDataPacket(socketio.Connect, "", ""), "0", 0},
		{"custom namespace", socketio.CreateDataPacket(socketio.Connect, "/admin", "", map[string]any{
			"sid": "test",
		}), "0/admin,{\"sid\":\"test\"}", 0},
		{"in case the connection is refused", socketio.CreateDataPacket(socketio.ConnectError, "", "", map[string]any{
			"message": "Not authorized",
		}), "4{\"message\":\"Not authorized\"}", 0},
	}

	for _, test := range tests {
		encode, b := IOParser.Encode(test.data)
		assert.Equal(t, test.expected, encode.String())
		assert.Equal(t, test.bufferLen, len(b))
	}

}

func TestIoParser_SendingAndReceivingData(t *testing.T) {
	tests := []testData{
		//{"with the main namespace", socketio.CreateDataPacket(socketio.Event, "", "", []string{"foo"}), "2[\"foo\"]", 0},
		//{"with a custom namespace", socketio.CreateDataPacket(socketio.Event, "/admin", "", []string{"bar"}), "2/admin,[\"bar\"]", 0},
		{"with binary data", socketio.CreateDataPacket(socketio.BinaryEvent, "", "", []byte{01, 02, 03, 04}), "51-[\"baz\",{\"_placeholder\":true,\"num\":0}]\n\n+ <Buffer <01 02 03 04>>", 0},
	}

	for _, test := range tests {
		encode, b := IOParser.Encode(test.data)
		assert.Equal(t, test.expected, encode.String())
		assert.Equal(t, test.bufferLen, len(b))
	}
}

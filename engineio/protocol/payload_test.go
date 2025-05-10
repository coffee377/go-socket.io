package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayloadType(t *testing.T) {
	at := assert.New(t)
	tests := []struct {
		b   byte
		typ PayloadType
		out byte
	}{
		{0, PayloadPlaintext, 0},
		{1, PayloadBinary, 1},
	}

	for _, test := range tests {
		typ := ByteToPayloadType(test.b)
		at.Equal(test.typ, typ)
		at.Equal(test.out, typ.Byte())
	}
}

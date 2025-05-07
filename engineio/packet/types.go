package packet

import "go-socket.io/engineio/frame"

type (
	Frame struct {
		FType frame.Type
		Data  []byte
	}

	Packet[T string | []byte | any] struct {
		Type Type
		Data T
	}
)

package socketio

import (
	"bytes"
	"encoding/json"
)

type Packet struct {
	Type        PacketType `json:"type"`
	Id          uint64     `json:"id" `
	Namespace   string     `json:"namespace"`
	Data        any        `json:"data,omitempty"`
	Attachments int        `json:"attachments"`
	buffers     [][]byte
}

func (p Packet) Deconstruct() *Packet {
	// todo
	return &p
}

func (p Packet) GetBuffers() [][]byte {
	return p.buffers
}

func (p Packet) String() string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	_ = encoder.Encode(p)
	return buffer.String()
}

func CreateDataPacket(packetType PacketType, namespace, event string, args ...any) Packet {
	var array []any
	if event != "" {
		array = append(array, event)
	}
	if len(args) > 0 {
		array = append(array, args...)
	}

	var data any = array
	if len(array) == 1 {
		data = array[0]
	}

	return Packet{
		Type:      packetType,
		Namespace: namespace,
		Data:      data,
	}
}

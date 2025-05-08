package socketio

// PacketType Type of packet.
type PacketType byte

const (
	// Connect type
	Connect PacketType = iota
	// Disconnect type
	Disconnect
	// Event type
	Event
	// Ack type
	Ack
	// ConnectError type
	ConnectError

	// BinaryEvent type
	BinaryEvent
	// BinaryAck type
	BinaryAck
)

func (pt PacketType) String() string {
	switch pt {
	case Connect:
		return "CONNECT"
	case Disconnect:
		return "DISCONNECT"
	case Event:
		return "EVENT"
	case Ack:
		return "pong"
	case ConnectError:
		return "CONNECT_ERROR"
	case BinaryEvent:
		return "BINARY_EVENT"
	case BinaryAck:
		return "BINARY_ACK"
	default:
		return "UNKNOWN"
	}
}

package protocol

// PacketType is the type of engine.io packet
type PacketType int

const (

	// PacketOpen is sent from the server when new transport is opened (recheck).
	PacketOpen PacketType = iota
	// PacketClose is request the close of this transport but does not shut down the
	// connection itself.
	PacketClose
	// PacketPing is sent by the client. Server should answer with a pong packet
	// containing the same data.
	PacketPing
	// PacketPong is sent by the server to respond to ping packets.
	PacketPong
	// PacketMessage is actual message, client and server should call their callbacks
	// with the data.
	PacketMessage
	// PacketUpgrade is sent before engine.io switches transport to test if server
	// and client can communicate over this transport. If this test succeed,
	// the client sends an upgrade packets which requests the server to flush
	// its cache on the old transport and switch to the new transport.
	PacketUpgrade
	// PacketNoop is a noop packet. Used primarily to force a poll cycle when an
	// incoming websocket connection is received.
	PacketNoop
)

func (id PacketType) String() string {
	switch id {
	case PacketOpen:
		return "open"
	case PacketClose:
		return "close"
	case PacketPing:
		return "ping"
	case PacketPong:
		return "pong"
	case PacketMessage:
		return "message"
	case PacketUpgrade:
		return "upgrade"
	case PacketNoop:
		return "noop"
	default:
		return "unknown"
	}
}

// Byte converts a PacketType to byte.
func (id PacketType) Byte() byte {
	return id.BinaryByte() + '0'
}

// BinaryByte converts a PacketType to byte in binary.
func (id PacketType) BinaryByte() byte {
	return byte(id) & 0x0F // 保留低4位
}

// StringByte converts a PacketType to byte in string.
func (id PacketType) StringByte() byte {
	return id.Byte()
}

// ByteToPacketType converts a byte to PacketType.
func ByteToPacketType(b byte) PacketType {
	return PacketType(b - '0')
}

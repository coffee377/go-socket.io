package protocol

type PayloadType byte

const (
	PayloadPlaintext PayloadType = iota
	PayloadBinary
)

// ByteToPayloadType converts a byte to PayloadType.
func ByteToPayloadType(b byte) PayloadType {
	return PayloadType(b)
}

// Byte returns type in byte.
func (t PayloadType) Byte() byte {
	return byte(t)
}

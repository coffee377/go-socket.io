package engineio

import (
	"bytes"
)

type ServerError interface {
	Code() int
	Message() string
}

type ErrorCode int

const (
	UnknownTransport ErrorCode = iota
	UnknownSid
	BadHandshakeMethod
	BadRequest
	Forbidden
	UnsupportedProtocolVersion
)

func (e ErrorCode) Code() int {
	return int(e)
}

func (e ErrorCode) Message() string {
	switch e {
	case UnknownTransport:
		return "Transport unknown"
	case UnknownSid:
		return "Session ID unknown"
	case BadHandshakeMethod:
		return "Bad handshake method"
	case BadRequest:
		return "Bad request"
	case Forbidden:
		return "Forbidden"
	case UnsupportedProtocolVersion:
		return "Unsupported protocol version"
	default:
		return "Unknown error"
	}
}

func (e ErrorCode) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	buffer.WriteByte('{')
	buffer.WriteByte('"')
	buffer.WriteString("code")
	buffer.WriteByte('"')
	buffer.WriteByte(':')
	buffer.WriteByte(byte(e.Code() + '0'))
	if e.Message() != "" {
		buffer.WriteByte(',')
		buffer.WriteByte('"')
		buffer.WriteString("message")
		buffer.WriteByte('"')
		buffer.WriteByte(':')
		buffer.WriteByte('"')
		buffer.WriteString(e.Message())
		buffer.WriteByte('"')
	}
	buffer.WriteByte('}')
	return buffer.Bytes(), nil
}

package transport

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

// Conn is a transport connection.
type Conn interface {
	io.Closer
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

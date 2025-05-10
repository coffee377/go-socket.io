package engineio

import (
	"net"
	"net/http"
	"net/url"
)

// Conn is connection by client session
type Conn interface {
	ID() string
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	SetContext(v interface{})
	Context() interface{}
}

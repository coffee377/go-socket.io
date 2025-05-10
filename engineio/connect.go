package engineio

import (
	"go-socket.io/engineio/session"
	"io"
	"net"
	"net/http"
	"net/url"
)

// Conn is connection by client session
type Conn interface {
	ID() string
	NextReader() (session.FrameType, io.ReadCloser, error)
	NextWriter(fType session.FrameType) (io.WriteCloser, error)
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	SetContext(v interface{})
	Context() interface{}
}

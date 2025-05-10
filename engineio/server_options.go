package engineio

import (
	"go-socket.io/engineio/session"
	"go-socket.io/engineio/transport"
	"net/http"
	"time"
)

// Options is options to create a server.
type Options struct {
	locked       bool
	PingTimeout  time.Duration
	PingInterval time.Duration

	Transports         []transport.Transport
	SessionIDGenerator session.IDGenerator
	RequestChecker     CheckerFunc
	ConnInitiator      ConnInitiatorFunc
}

func (c *Options) getRequestChecker() CheckerFunc {
	if c != nil && c.RequestChecker != nil {
		return c.RequestChecker
	}
	return defaultChecker
}

func (c *Options) getConnConnInitiator() ConnInitiatorFunc {
	if c != nil && c.ConnInitiator != nil {
		return c.ConnInitiator
	}
	return defaultConnInitiator
}

func (c *Options) getPingTimeout() time.Duration {
	if c != nil && c.PingTimeout != 0 {
		return c.PingTimeout
	}
	return time.Minute
}

func (c *Options) getPingInterval() time.Duration {
	if c != nil && c.PingInterval != 0 {
		return c.PingInterval
	}
	return time.Second * 20
}

func (c *Options) getTransport() []transport.Transport {
	if c != nil && len(c.Transports) != 0 {
		return c.Transports
	}
	return []transport.Transport{
		//polling.Default,
		//websocket.Default,
	}
}

func (c *Options) getSessionIDGenerator() session.IDGenerator {
	if c != nil && c.SessionIDGenerator != nil {
		return c.SessionIDGenerator
	}
	return &session.DefaultIDGenerator{}
}

func defaultChecker(*http.Request) (http.Header, error) {
	return nil, nil
}

func defaultConnInitiator(*http.Request, Conn) {}

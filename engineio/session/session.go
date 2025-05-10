package session

import (
	"go-socket.io/engineio/transport"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Pauser is connection which can be paused and resumes.
type Pauser interface {
	Pause()
	Resume()
}

type Session struct {
	conn      transport.Conn
	params    transport.ConnParameters
	transport string

	context interface{}

	upgradeLocker sync.RWMutex
}

func New(conn transport.Conn, sid, transport string, params transport.ConnParameters) (*Session, error) {
	params.SID = sid

	ses := &Session{
		transport: transport,
		conn:      conn,
		params:    params,
	}

	//if err := ses.setDeadline(); err != nil {
	//	if closeErr := ses.Close(); closeErr != nil {
	//		logger.Error("session close:", closeErr)
	//	}
	//
	//	return nil, err
	//}

	return ses, nil
}

func (s *Session) SetContext(v interface{}) {
	s.context = v
}

func (s *Session) Context() interface{} {
	return s.context
}

func (s *Session) ID() string {
	return s.params.SID
}

func (s *Session) Transport() string {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()

	return s.transport
}

func (s *Session) Close() error {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()

	return s.conn.Close()
}

func (s *Session) URL() url.URL {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()

	return s.conn.URL()
}

func (s *Session) LocalAddr() net.Addr {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()

	return s.conn.LocalAddr()
}

func (s *Session) RemoteAddr() net.Addr {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()

	return s.conn.RemoteAddr()
}

func (s *Session) RemoteHeader() http.Header {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()

	return s.conn.RemoteHeader()
}

func (s *Session) Upgrade(transport string, conn transport.Conn) {
	//go s.upgrading(transport, conn)
}

func (s *Session) InitSession() error {
	//w, err := s.nextWriter(frame.String, packet.OPEN)
	//if err != nil {
	//	if closeErr := s.Close(); closeErr != nil {
	//		logger.Error("close session with string frame and packet open:", closeErr)
	//	}
	//
	//	return err
	//}
	//
	//if _, err := s.params.WriteTo(w); err != nil {
	//	if closeErr := w.Close(); closeErr != nil {
	//		logger.Error("close writer:", closeErr)
	//	}
	//
	//	if closeErr := s.Close(); closeErr != nil {
	//		logger.Error("close session:", closeErr)
	//	}
	//
	//	return err
	//}
	//
	//if err := w.Close(); err != nil {
	//	if closeErr := s.Close(); closeErr != nil {
	//		logger.Error("close session:", closeErr)
	//	}
	//
	//	return err
	//}

	return nil
}

func (s *Session) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.upgradeLocker.RLock()
	conn := s.conn
	s.upgradeLocker.RUnlock()

	if h, ok := conn.(http.Handler); ok {
		h.ServeHTTP(w, r)
	}
}

func (s *Session) setDeadline() error {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()

	deadline := time.Now().Add(s.params.PingTimeout)

	err := s.conn.SetReadDeadline(deadline)
	if err != nil {
		return err
	}

	return s.conn.SetWriteDeadline(deadline)
}

func (s *Session) upgrading(t string, conn transport.Conn) {
	// Read a ping from the client.
	//err := conn.SetReadDeadline(time.Now().Add(s.params.PingTimeout))
	//if err != nil {
	//	logger.Error("set read deadline:", err)
	//
	//	if closeErr := conn.Close(); closeErr != nil {
	//		logger.Error("close connect after set read deadline:", closeErr)
	//	}
	//
	//	return
	//}
	//
	//ft, pt, r, err := conn.NextReader()
	//if err != nil {
	//	logger.Error("get next reader:", err)
	//
	//	if closeErr := conn.Close(); closeErr != nil {
	//		logger.Error("close connect after get next reader:", closeErr)
	//	}
	//
	//	return
	//}
	//
	//if pt != packet.PING {
	//	if err := r.Close(); err != nil {
	//		logger.Error("close reade:", err)
	//	}
	//
	//	if err := conn.Close(); err != nil {
	//		logger.Error("close connect:", err)
	//	}
	//
	//	return
	//}
	//// Wait to close the reader until after data is read and echoed in the reply.
	//
	//// Sent a pong in reply.
	//err = conn.SetWriteDeadline(time.Now().Add(s.params.PingTimeout))
	//if err != nil {
	//	logger.Error("set write deadline:", err)
	//
	//	if closeErr := r.Close(); closeErr != nil {
	//		logger.Error("close reader:", closeErr)
	//	}
	//
	//	if closeErr := conn.Close(); closeErr != nil {
	//		logger.Error("close connect:", closeErr)
	//	}
	//
	//	return
	//}
	//
	//w, err := conn.NextWriter(ft, packet.PONG)
	//if err != nil {
	//	logger.Error("get next writer with pong packet:", err)
	//
	//	if closeErr := r.Close(); closeErr != nil {
	//		logger.Error("close reader:", closeErr)
	//	}
	//
	//	if closeErr := conn.Close(); closeErr != nil {
	//		logger.Error("close connect:", closeErr)
	//	}
	//
	//	return
	//}
	//
	//// echo
	//if _, err = io.Copy(w, r); err != nil {
	//	logger.Error("copy from reader to writer:", err)
	//
	//	if closeErr := w.Close(); closeErr != nil {
	//		logger.Error("close writer:", closeErr)
	//	}
	//
	//	if closeErr := r.Close(); closeErr != nil {
	//		logger.Error("close reader:", closeErr)
	//	}
	//
	//	if closeErr := conn.Close(); closeErr != nil {
	//		logger.Error("close connect:", closeErr)
	//	}
	//
	//	return
	//}
	//
	//if err = r.Close(); err != nil {
	//	logger.Error("close reader:", err)
	//
	//	if closeErr := w.Close(); closeErr != nil {
	//		logger.Error("close writer:", closeErr)
	//	}
	//
	//	if closeErr := conn.Close(); closeErr != nil {
	//		logger.Error("close connect:", closeErr)
	//	}
	//
	//	return
	//}
	//
	//if err = w.Close(); err != nil {
	//	logger.Error("close writer:", err)
	//
	//	if closeErr := conn.Close(); closeErr != nil {
	//		logger.Error("close connect:", closeErr)
	//	}
	//
	//	return
	//}
	//
	//// Pause the old connection.
	//s.upgradeLocker.RLock()
	//old := s.conn
	//s.upgradeLocker.RUnlock()
	//
	//p, ok := old.(Pauser)
	//if !ok {
	//	// old transport doesn't support upgrading
	//	if closeErr := conn.Close(); closeErr != nil {
	//		logger.Error("close connect after get pauser:", closeErr)
	//	}
	//
	//	return
	//}
	//
	//p.Pause()
	//
	//// Prepare to resume the connection if upgrade fails.
	//defer func() {
	//	if p != nil {
	//		p.Resume()
	//	}
	//}()
	//
	//// Check for upgrade packet from the client.
	//_, pt, r, err = conn.NextReader()
	//if err != nil {
	//	logger.Error("get next reader:", err)
	//
	//	if closeErr := conn.Close(); closeErr != nil {
	//		logger.Error("close connect:", closeErr)
	//	}
	//
	//	return
	//}
	//
	//if pt != packet.UPGRADE {
	//	if closeErr := r.Close(); closeErr != nil {
	//		logger.Error("close reader:", closeErr)
	//	}
	//
	//	if closeErr := conn.Close(); closeErr != nil {
	//		logger.Error("close connect:", closeErr)
	//	}
	//
	//	return
	//}
	//
	//if err = r.Close(); err != nil {
	//	logger.Error("close reader:", err)
	//
	//	if closeErr := conn.Close(); closeErr != nil {
	//		logger.Error("close connect:", closeErr)
	//	}
	//
	//	return
	//}
	//
	//// Successful upgrade.
	//s.upgradeLocker.Lock()
	//s.conn = conn
	//s.transport = t
	//s.upgradeLocker.Unlock()
	//
	//p = nil
	//
	//if closeErr := old.Close(); closeErr != nil {
	//	logger.Error("close old connection:", closeErr)
	//}
}

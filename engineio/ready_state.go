package engineio

// ReadyState The possible states a connection can be in at any given time.
// These states also apply to transports.
type ReadyState int

const (
	// ReadyStateOpening The connection is opening. No IO operation can be performed.
	ReadyStateOpening ReadyState = iota
	// ReadyStateOpen The connection is open. IO operations can be performed on it.
	ReadyStateOpen
	// ReadyStateClosing The connection is closing. No IO operation can be performed.
	ReadyStateClosing
	// ReadyStateClosed The connection is closed. No IO operation can be performed.
	ReadyStateClosed
)

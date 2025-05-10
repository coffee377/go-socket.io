package protocol

// EngineIOVersion represents the Engine.IO protocol version
type EngineIOVersion int

const (
	// UNKNOWN represents an unknown version
	UNKNOWN EngineIOVersion = 0
	// V2 represents Engine.IO version 2
	V2 EngineIOVersion = 2
	// V3 represents Engine.IO version 3
	V3 EngineIOVersion = 3
	// V4 represents Engine.IO version 4
	V4 EngineIOVersion = 4
)

func (v EngineIOVersion) Value() int {
	return int(v)
}

const EngineIOVersionParamName = "EIO"

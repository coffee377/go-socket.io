package parser

import "bytes"

var ProtocolV3 = v3{
	buf: &bytes.Buffer{},
}

type v3 struct {
	buf *bytes.Buffer
}

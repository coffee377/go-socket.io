package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decode(t *testing.T) {
	u1 := []string{"XRsZ", "HUkJ", "T&g", "Q&w", "GRz1", "JYx3Gb#8Pb5R", "VdgJHc#wJb59", "4LJITMx8UMcA"}
	u2 := []string{"Evl", "Prd", "N", "C", "Dms", "location", "protocol", "127.0.0.1"}
	for i := 0; i < len(u1); i++ {
		s1 := decode(u1[i])
		assert.Equal(t, u2[i], string(s1))
	}
}

func Test_encode(t *testing.T) {
	sjs := NewSpreadJSLicense()
	buf := &bytes.Buffer{}
	err := sjs.Output(buf)
	assert.Nil(t, err)
	assert.Equal(t, lic, buf.String())

	buf.Reset()
	sjs.Data.Evaluation = false
	err = sjs.Output(buf)
	assert.Nil(t, err)
	sjs2 := sjs.Read(buf.String())
	assert.Equal(t, sjs, sjs2)
	assert.Equal(t, false, sjs2.Data.Evaluation)
	println(buf.String())

}

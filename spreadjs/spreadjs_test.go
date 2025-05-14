package main

import (
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

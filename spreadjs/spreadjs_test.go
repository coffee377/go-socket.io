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
	sjs.HexHash = "49B4E07C"
	sjs.Data.Evaluation = false
	sjs.Data.Expiration = "20260101"
	sjs.Data.Domains = "127.0.0.1,10.1.40.93"
	err = sjs.Output(buf)
	assert.Nil(t, err)
	sjs2 := sjs.Read(buf.String())
	assert.Equal(t, sjs, sjs2)
	assert.Equal(t, false, sjs2.Data.Evaluation)
	println(buf.String())
}

func Test_swapCaseAndOffsetDigit(t *testing.T) {
	digits := "0123456789"
	expected := map[int]string{
		-9: "1234567890",
		-8: "2345678901",
		-7: "3456789012",
		-6: "4567890123",
		-5: "5678901234",
		-4: "6789012345",
		-3: "7890123456",
		-2: "8901234567",
		-1: "9012345678",
		0:  "0123456789",
		1:  "1234567890",
		2:  "2345678901",
		3:  "3456789012",
		4:  "4567890123",
		5:  "5678901234",
		6:  "6789012345",
		7:  "7890123456",
		8:  "8901234567",
		9:  "9012345678",
	}
	for i, d := range digits {
		for offset, s := range expected {
			assert.Equal(t, rune(s[i]), characterConversion(d, offset))
		}
	}
}

func Test_hexLic(t *testing.T) {
	// 615274881    24AC5981
	var raw = `E879948536774266#B1{"Anl":{"dsr":false,"flg":["ReportSheet","DataChart"]},"Id":"879948536774266","Evl":true,"CNa":"安徽晶奇网络科技股份有限公司","Dms":"127.0.0.1","Exp":"20250606","Crt":"20250507 032315","Prd":[{"N":"Spread JS v.18","C":"BJIH"}]}`
	d, h := A(raw)
	assert.Equal(t, 615274881, d)
	assert.Equal(t, "24AC5981", h)
}

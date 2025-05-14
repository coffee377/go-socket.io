package main

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
)

type Product struct {
	Name string `json:"N"`
	Code string `json:"C"`
}

type Data struct {
	Anl        Anl       `json:"Anl"`
	Id         string    `json:"Id"`
	Evaluation bool      `json:"Evl"`
	CNa        string    `json:"CNa"`
	Dms        string    `json:"Dms"`
	Exp        string    `json:"Exp"`
	CreateTime string    `json:"Crt"`
	Products   []Product `json:"Prd"`
}

type Anl struct {
	Dsr bool     `json:"dsr"`
	Flg []string `json:"flg"`
}

type License interface {
	Output(writer io.Writer) error
	Read(lic, sep, prefix string) *SpreadJSLicense
}

type SpreadJSLicense struct {
	ac        string
	pk        string
	R         int    `json:"_r"`
	H         string `json:"H"`
	Signature string `json:"S"`
	Data      Data   `json:"D"`
}

func (s *SpreadJSLicense) Output(writer io.Writer) error {
	bytes, err := json.Marshal(s)
	encoded := encode(bytes)
	if err != nil {
		return err
	}
	if len(encoded) > 0 {
		_, err = writer.Write(encoded)
	}
	return err
}

func (s *SpreadJSLicense) Read(lic, sep, prefix string) *SpreadJSLicense {
	if prefix == "" {
		prefix = "E"
	}
	regStr := fmt.Sprintf("(%s)(%s)#(%s)(%s)", prefix, ".*", sep, ".*")
	all := regexp.MustCompile(regStr).FindStringSubmatch(lic)
	s.Data.Id = all[2]
	licData := all[4]

	bytes := decode(licData)
	_ = json.Unmarshal(bytes, s)
	return s
}

func NewSpreadJSLicense() SpreadJSLicense {
	return SpreadJSLicense{
		R:         1332046125,
		H:         "24AC5981",
		Signature: "N++NtKxSFV4lGqBTqdu2D94fbq/BuExoKTHFOWS0R6X28SaPAak29Y7chZPlHcD/owaQy1kU4dT3gI281yta1tpIxrKgNXYrLazMw4wTceDyKGSHXrm7csAltd3YTxJu/wLXJZS6ZABjQ7W0jF5skv8ZndxwgeDjuATtOVPvv3v3qSAxRlK9uKKpyaRZ+cJwZ2fuv56vHBLq5KyJAAO2E2tm8kx1bggegCc2Kh8yTvIq2kCWma2dSFoZowPDlWd8bhQrFT5N2eyhyuxD3oB3W4lD3iLkc/r0pxcK8gb2Xrv+aCE6rsTe4QQD/DSoYWI0tR7NvWXXhyOZVIsv2lBTjQ==",
		Data: Data{Anl{false, []string{"ReportSheet", "DataChart"}},
			"879948536774266", true, "安徽晶奇网络科技股份有限公司", "127.0.0.1",
			"20250606", "20250507 032315", []Product{
				{"Spread JS v.18", "BJIH"},
			},
		},
	}
}

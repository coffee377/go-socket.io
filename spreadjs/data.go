package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
)

type Product struct {
	Name string `json:"N"` // 产品名称
	Code string `json:"C"` // 产品代码
}

type Data struct {
	Annual     Annual    `json:"Anl"`           // 授权信息
	Id         string    `json:"Id,omitempty"`  // 授权标识
	Evaluation bool      `json:"Evl,omitempty"` // 是否为评估使用版
	CNa        string    `json:"CNa,omitempty"` // 客户名称
	Domains    string    `json:"Dms,omitempty"` // 域名
	Ips        string    `json:"Ips,omitempty"` // IP 地址
	Expiration string    `json:"Exp,omitempty"` // 过期时间
	CreateTime string    `json:"Crt,omitempty"` // 创建时间
	Products   []Product `json:"Prd"`           // 产品信息
}

type Annual struct {
	Distribution bool     `json:"dsr"` // 是否为分发版
	PluginFlags  []string `json:"flg"` // 插件标志
}

type License interface {
	Read(lic string) *SpreadJSLicense
	Output(writer io.Writer) error
}

type SpreadJSLicense struct {
	sep      string
	prefix   string                 // 前缀
	prefixFn func(data Data) string // 前缀计算函数
	licData  string                 // 加密数据

	R         int    `json:"_r,omitempty"`
	HexHash   string `json:"H,omitempty"` // 授权数据 16 进制
	Signature string `json:"S,omitempty"` // 签名
	Data      Data   `json:"D,omitempty"` // 授权数据
}

func (s *SpreadJSLicense) Read(lic string) *SpreadJSLicense {
	if s.sep == "" {
		s.sep = "B1"
	}
	regStr := fmt.Sprintf("(%s)#(%s)(%s)", ".*", s.sep, ".*")
	all := regexp.MustCompile(regStr).FindStringSubmatch(lic)
	s.prefix = all[1]
	s.licData = all[3]

	bData := decode(s.licData)
	_ = json.Unmarshal(bData, s)
	return s
}

func (s *SpreadJSLicense) GetPrefix() string {
	return s.prefix
}

func (s *SpreadJSLicense) Output(writer io.Writer) error {
	enc, _ := json.Marshal(s.Data)
	prefix := s.prefixFn(s.Data)
	d := fmt.Sprintf("%s#%s%s", prefix, s.sep, string(enc))

	_, hexCal := A(d)
	if hexCal != "" {
		s.HexHash = hexCal
	}

	byt, err := json.Marshal(s)
	encoded := encode(byt)
	if err != nil {
		return err
	}

	if len(encoded) > 0 {
		buf := bytes.Buffer{}
		buf.WriteString(fmt.Sprintf("%s", prefix))
		buf.WriteString(fmt.Sprintf("#%s", s.sep))
		buf.Write(encoded)
		_, err = writer.Write(buf.Bytes())
	}
	return err
}

func NewSpreadJSLicense() *SpreadJSLicense {
	return &SpreadJSLicense{
		sep: "B1",
		prefixFn: func(data Data) string {
			return fmt.Sprintf("%s%s", "E", data.Id)
		},
		R:         1332046125,
		Signature: "N++NtKxSFV4lGqBTqdu2D94fbq/BuExoKTHFOWS0R6X28SaPAak29Y7chZPlHcD/owaQy1kU4dT3gI281yta1tpIxrKgNXYrLazMw4wTceDyKGSHXrm7csAltd3YTxJu/wLXJZS6ZABjQ7W0jF5skv8ZndxwgeDjuATtOVPvv3v3qSAxRlK9uKKpyaRZ+cJwZ2fuv56vHBLq5KyJAAO2E2tm8kx1bggegCc2Kh8yTvIq2kCWma2dSFoZowPDlWd8bhQrFT5N2eyhyuxD3oB3W4lD3iLkc/r0pxcK8gb2Xrv+aCE6rsTe4QQD/DSoYWI0tR7NvWXXhyOZVIsv2lBTjQ==",
		Data: Data{Annual{false, []string{"ReportSheet", "DataChart"}},
			"879948536774266", true, "安徽晶奇网络科技股份有限公司",
			"127.0.0.1", "", "20250606", "20250507 032315", []Product{
				{"Spread JS v.18", "BJIH"},
			},
		},
	}
}

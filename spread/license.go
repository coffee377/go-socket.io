package spread

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"
)

// Output(writer io.Writer) error // 输出授权文件

type License interface {
	//GetSeparator() string // 分隔符
	//GetPrefix() string    // 前缀
	GetData() *Data // 授权数据

	R() int          //
	Sign() string    // 签名
	HexHash() string // 十六进制哈希

	Read(lic string) License // 读取解析授权文件
	// ReadFromFile(file os.File) License // 从文件读取解析授权文件

	PrefixGenerate(data Data) string // 生成前缀
	Output(writer io.Writer) error   // 输出授权文件

	json.Marshaler
	json.Unmarshaler
}

type options struct {
	sep       string
	createdAt time.Time     // 创建时间
	duration  time.Duration // 有效时长
	domains   []string      // 域名
	ips       []string      // IP
	prefixFn  func(data Data) string
}

type license struct {
	opts options

	prefix  string // 前缀
	licData string // 加密数据

	r         int    // _r
	hash      string // H 授权数据 16 进制
	signature string // S 签名
	data      *Data  // D 授权数据
}

type jsonLic struct {
	R int    `json:"_r,omitempty"` //
	H string `json:"H,omitempty"`  // H 授权数据 Hash 16 进制
	S string `json:"S,omitempty"`  // S 签名
	D Data   `json:"D,omitempty"`  // D 授权数据
}

//func (l *license) GetSeparator() string {
//	return l.sep
//}
//
//func (l *license) GetPrefix() string {
//	return l.prefix
//}

func (l *license) GetData() *Data {
	return l.data
}

func (l *license) R() int {
	return l.r
}

func (l *license) Sign() string {
	return l.signature
}

func (l *license) HexHash() string {
	return l.hash
}

func (l *license) Read(lic string) License {
	regStr := fmt.Sprintf("(%s)#(%s)(%s)", ".*", l.opts.sep, ".*")
	all := regexp.MustCompile(regStr).FindStringSubmatch(lic)
	l.prefix = all[1]
	l.licData = all[3]

	bData := decode(l.licData)
	_ = json.Unmarshal(bData, l)
	// 显式类型转换（可选）
	return l
}

func (l *license) MarshalJSON() ([]byte, error) {
	jl := &jsonLic{}
	jl.R = l.r
	jl.H = l.hash
	jl.S = l.signature
	jl.D = *l.data
	return json.Marshal(jl)
}

func (l *license) UnmarshalJSON(data []byte) error {
	jl := &jsonLic{}
	err := json.Unmarshal(data, jl)
	l.r = jl.R
	l.hash = jl.H
	l.signature = jl.S
	l.data = &jl.D
	return err
}

func (l *license) PrefixGenerate(data Data) string {
	if l.opts.prefixFn != nil {
		return l.opts.prefixFn(data)
	}
	return fmt.Sprintf("%s%s", "E", data.Id)
}

func (l *license) Output(writer io.Writer) error {
	enc, _ := json.Marshal(l.data)
	prefix := l.PrefixGenerate(*l.data)
	d := fmt.Sprintf("%s#%s%s", prefix, l.opts.sep, string(enc))

	h := hexHash(d)
	if l.hash != h {
		l.hash = h
	}

	// todo 计算签名
	//l.signature = ""

	byt, err := json.Marshal(l)
	encoded := encode(byt)
	if err != nil {
		return err
	}

	if len(encoded) > 0 {
		buf := bytes.Buffer{}
		buf.WriteString(fmt.Sprintf("%s", prefix))
		buf.WriteString(fmt.Sprintf("#%s", l.opts.sep))
		buf.Write(encoded)
		_, err = writer.Write(buf.Bytes())
	}
	return err
}

type Options func(opts *options)

func NewSpreadJSLicense(opts ...Options) License {
	o := &options{sep: "B1"}
	for _, opt := range opts {
		opt(o)
	}
	return &license{
		opts: *o,
	}
	//annual := &Annual{false, []string{"ReportSheet", "DataChart"}}
	//products := &[]Product{
	//	{"Spread JS v.18", "BJIH"},
	//}
	//return &license{
	//	sep: "B1",
	//	prefixFn: func(data Data) string {
	//		return fmt.Sprintf("%s%s", "E", data.Id)
	//	},
	//	//R: 1332046125,
	//	//Signature: "N++NtKxSFV4lGqBTqdu2D94fbq/BuExoKTHFOWS0R6X28SaPAak29Y7chZPlHcD/owaQy1kU4dT3gI281yta1tpIxrKgNXYrLazMw4wTceDyKGSHXrm7csAltd3YTxJu/wLXJZS6ZABjQ7W0jF5skv8ZndxwgeDjuATtOVPvv3v3qSAxRlK9uKKpyaRZ+cJwZ2fuv56vHBLq5KyJAAO2E2tm8kx1bggegCc2Kh8yTvIq2kCWma2dSFoZowPDlWd8bhQrFT5N2eyhyuxD3oB3W4lD3iLkc/r0pxcK8gb2Xrv+aCE6rsTe4QQD/DSoYWI0tR7NvWXXhyOZVIsv2lBTjQ==",
	//	Data: &Data{
	//		annual,
	//		"879948536774266", true, "安徽晶奇网络科技股份有限公司",
	//		"127.0.0.1", "", "20250606", "20250507 032315",
	//		products,
	//	},
	//}
}

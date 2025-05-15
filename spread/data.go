package spread

type Product struct {
	Name string `json:"N"` // 产品名称
	Code string `json:"C"` // 产品代码
}

type Data struct {
	Annual     *Annual    `json:"Anl,omitempty"` // 授权信息
	Id         string     `json:"Id,omitempty"`  // 授权标识
	Evaluation bool       `json:"Evl,omitempty"` // 是否为评估使用版
	CNa        string     `json:"CNa,omitempty"` // 客户名称
	Domains    string     `json:"Dms,omitempty"` // 域名
	Ips        string     `json:"Ips,omitempty"` // IP 地址
	Expiration string     `json:"Exp,omitempty"` // 过期时间
	CreateTime string     `json:"Crt,omitempty"` // 创建时间
	Products   []*Product `json:"Prd,omitempty"` // 产品信息
}

type Annual struct {
	Distribution bool     `json:"dsr"` // 是否为分发版
	PluginFlags  []string `json:"flg"` // 插件标志
}

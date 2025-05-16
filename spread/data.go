package spread

import "fmt"

type Product struct {
	Name     string `json:"N"` // 产品名称
	Code     string `json:"C"` // 产品代码
	designer bool
}

var prods = map[int][]*Product{
	18: {
		&Product{"Spread JS v.18", "BJIH", false},
		&Product{"SpreadJS-Designer-Addon v.18", "33Y9", true},
	},
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

// Plugin 在线表格编辑器 Designer
// 报表 / ReportSheet
// 数据透视表 / PivotTable
// 集算表 / TableSheet
// 甘特图 / GanttSheet
// 数据图表 / DataChart
type Plugin int

const (
	PluginDesigner    Plugin = 1 << 0
	PluginReportSheet Plugin = 1 << 1
	PluginPivotTable  Plugin = 1 << 2
	PluginTableSheet  Plugin = 1 << 3
	PluginGanttSheet  Plugin = 1 << 4
)

func (p Plugin) String() string {
	switch p {
	case PluginDesigner:
		return "Designer"
	case PluginReportSheet:
		return "ReportSheet"
	case PluginPivotTable:
		return "PivotTable"
	case PluginTableSheet:
		return "TableSheet"
	case PluginGanttSheet:
		return "GanttSheet"

	}
	return fmt.Sprintf("%d", p)
}

func PluginsFrom(mask int) []string {
	plugins := make([]Plugin, 0)
	add(&plugins, mask, PluginReportSheet)
	add(&plugins, mask, PluginPivotTable)
	add(&plugins, mask, PluginTableSheet)
	add(&plugins, mask, PluginGanttSheet)
	result := make([]string, len(plugins))
	for i, plugin := range plugins {
		result[i] = plugin.String()
	}
	return result
}

func add(plugins *[]Plugin, mask int, plugin Plugin) {
	if mask&int(plugin) == int(plugin) {
		*plugins = append(*plugins, plugin)
	}
}

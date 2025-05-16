package spread

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewSpreadJSLicense(t *testing.T) {
	opts := make([]Options, 0)
	opts = append(opts, WithSeparator("B1"))
	//opts = append(opts, WithCreateTime(""))
	opts = append(opts, WithDeadline("24h"))
	opts = append(opts, WithDomain("designer/0.0.0.0"))
	//opts = append(opts, WithIP("10.1.82.8", "*"))
	opts = append(opts, WithPrefix(func(data Data) string {
		return fmt.Sprintf("%s%s", "E", data.Id)
	}))
	//opts = append(opts, WithoutTrial())
	//opts = append(opts, WithPlugin(8))
	//opts = append(opts, WithDesigner())
	sjs := NewSpreadJSLicense(opts...)
	//raw := `E879948536774266#B1{"Anl":{"dsr":false,"flg":["ReportSheet","DataChart"]},"Id":"879948536774266","Evl":true,"CNa":"安徽晶奇网络科技股份有限公司","Dms":"127.0.0.1","Exp":"20250606","Crt":"20250507 032315","Prd":[{"N":"Spread JS v.18","C":"BJIH"}]}`
	//enc, _ := json.Marshal(sLic.GetData())
	//res := fmt.Sprintf("E%s#%s%s", sLic.GetData().Id, sLic.GetSeparator(), string(enc))
	//assert.Equal(t, raw, res)
	assert.NotNil(t, sjs)
	assert.NotNil(t, sjs.GetData())
	_ = sjs.Output(os.Stdout)
	println()
}

func Test_license_GetData(t *testing.T) {

}

func Test_license_HexHash(t *testing.T) {

}

func Test_license_MarshalJSON(t *testing.T) {

}

func Test_license_Output(t *testing.T) {

}

func Test_license_PrefixGenerate(t *testing.T) {

}

func Test_license_R(t *testing.T) {

}

func Test_license_Read(t *testing.T) {

}

func Test_license_Sign(t *testing.T) {

}

func Test_license_UnmarshalJSON(t *testing.T) {

}

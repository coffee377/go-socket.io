package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewSpreadJSLicense(t *testing.T) {
	license := NewSpreadJSLicense()
	//pK := "l6/zrbWoSbcLFwEetFh38rH3ErBZE9H+Cqix3R+wTlfA1wD5B+lUcCQn+EJ60I4RGrm0x1sFjkiLWwB0jAn6BWZv0W4WbqAKriOdeoivxDp1Wmjs3qkEDhvbsjPtfvwx2BHil6o+/tDrdMJQSGs18WZm2PoQLQuL+9VhZ4FNRHUQU3Jtioke/OZEGHJOdYVwvCGalzBad6QFOiVbDBQPePpS3++GJzOxN8SN/7lyS5/IdKiy3WJRaVGkB370+HbN6hKraDfUgReLX26yxRaKC/5aWnGAJ2NnWLoGyAGRcwT9dVjo4bcAZNrrA0U9JVKQxaSskhdv2p49XzJkltXx5w=="
	aC := "B1"

	raw := `E879948536774266#B1{"Anl":{"dsr":false,"flg":["ReportSheet","DataChart"]},"Id":"879948536774266","Evl":true,"CNa":"安徽晶奇网络科技股份有限公司","Dms":"127.0.0.1","Exp":"20250606","Crt":"20250507 032315","Prd":[{"N":"Spread JS v.18","C":"BJIH"}]}`
	enc, _ := json.Marshal(license.Data)
	res := fmt.Sprintf("E%s#%s%s", license.Data.Id, aC, string(enc))
	assert.Equal(t, raw, res)
}

func TestSpreadJSLicense_Decrypt(t *testing.T) {
	// SpreadJS 授权
	lic := "E879948536774266#B12c3MC36MEhXd9hWelJjT5QlRyFFaihDZXxGRQd7ba3mRTRmMh56VDtmMxlkdUlHOotkMjN4Zld6ZiFDerhTb4JTRy2UQBpUeLVTcMJES6ZTN6VnZyo5dKN6KaJVY9B7SLVXOLxmU8F4UxNjdzYndQZ5T4RVQ5pGRld6d8RmbahjdrNXNGpGMXdTUqJUQaZzUapEWMd7L5pEeUl5MkRHbBN7Y74mcYh4UHtUeEV6YUdHN75kehxkcZhlTntkc8lEc4FTY4lXM8ITSnNDVkRTVrFTeRF6dv3CRjhEbQpFajdTW9IzahFEUhNFOyglNSBzUX3kRIR5SvhXR5J4LxJmZ4kDRyUHZxRlQxdEb4YlRTh7S494Kr8kI0IyUiwiIxgTO5MUQ4IjI0ICSiwSNyEjN4AjMzMTM0IicfJye#4Xfd5nIIlkSCJiOiMkIsICOx8idgMlSgQWYlJHcTJiOi8kI1tlOiQmcQJCLiUTMzIzMwAyNwUDM5IDMyIiOiQncDJCLiYDM6ATNyAjMiojIwhXRiwiIx8CMuAjL7ITMiojIz5GRiwiI8+Y9sWY9QmZ0Jyp93uL9hKI0Aqo9Re09cu19R619HWa96mp930b9J0a9iojIh94QiwSZ5JHd0ICb6VkIsIiN6IDN7cjNzUDO4kTO7gjI0ICZJJCL35lI4JXYoNUY4FGRiwiI4VWZoNFdy3GclJlIbpjInxmZiwSZzxWYmpjIyNHZisnOiwmbBJye0ICRiwiI34TUqRlQsJjdzlkVa3UeohFWXZnT7IFdwk4VZ36UE3CRRFFNlR5cyZTRDF6K6JHWyI6Z8s4Y8BHMy3yYrxUaeQZb"
	//	{
	//	   "_r": 1332046125,
	//	   "H": "24AC5981",
	//	   "S": "N++NtKxSFV4lGqBTqdu2D94fbq/BuExoKTHFOWS0R6X28SaPAak29Y7chZPlHcD/owaQy1kU4dT3gI281yta1tpIxrKgNXYrLazMw4wTceDyKGSHXrm7csAltd3YTxJu/wLXJZS6ZABjQ7W0jF5skv8ZndxwgeDjuATtOVPvv3v3qSAxRlK9uKKpyaRZ+cJwZ2fuv56vHBLq5KyJAAO2E2tm8kx1bggegCc2Kh8yTvIq2kCWma2dSFoZowPDlWd8bhQrFT5N2eyhyuxD3oB3W4lD3iLkc/r0pxcK8gb2Xrv+aCE6rsTe4QQD/DSoYWI0tR7NvWXXhyOZVIsv2lBTjQ==",
	//	   "D": {
	//	       "Anl": {
	//	           "dsr": false,
	//	           "flg": [
	//	               "ReportSheet",
	//	               "DataChart"
	//	           ]
	//	       },
	//	       "Id": "879948536774266",
	//	       "Evl": true,
	//	       "CNa": "安徽晶奇网络科技股份有限公司",
	//	       "Dms": "127.0.0.1",
	//	       "Exp": "20250606",
	//	       "Crt": "20250507 032315",
	//	       "Prd": [
	//	           {
	//	               "N": "Spread JS v.18",
	//	               "C": "BJIH"
	//	           }
	//	       ]
	//	   }
	//	}
	sjs := &SpreadJSLicense{}
	sjs.Read(lic, "B1", "")

	assert.Equal(t, 1332046125, sjs.R)
	assert.Equal(t, "24AC5981", sjs.H)
	assert.Equal(t, "N++NtKxSFV4lGqBTqdu2D94fbq/BuExoKTHFOWS0R6X28SaPAak29Y7chZPlHcD/owaQy1kU4dT3gI281yta1tpIxrKgNXYrLazMw4wTceDyKGSHXrm7csAltd3YTxJu/wLXJZS6ZABjQ7W0jF5skv8ZndxwgeDjuATtOVPvv3v3qSAxRlK9uKKpyaRZ+cJwZ2fuv56vHBLq5KyJAAO2E2tm8kx1bggegCc2Kh8yTvIq2kCWma2dSFoZowPDlWd8bhQrFT5N2eyhyuxD3oB3W4lD3iLkc/r0pxcK8gb2Xrv+aCE6rsTe4QQD/DSoYWI0tR7NvWXXhyOZVIsv2lBTjQ==", sjs.Signature)
	assert.Equal(t, false, sjs.Data.Anl.Dsr)
	assert.Equal(t, []string{"ReportSheet", "DataChart"}, sjs.Data.Anl.Flg)
	assert.Equal(t, "879948536774266", sjs.Data.Id)
	assert.Equal(t, true, sjs.Data.Evaluation)
	assert.Equal(t, "安徽晶奇网络科技股份有限公司", sjs.Data.CNa)
	assert.Equal(t, "127.0.0.1", sjs.Data.Dms)
	assert.Equal(t, "20250606", sjs.Data.Exp)
	assert.Equal(t, "20250507 032315", sjs.Data.CreateTime)
	assert.Equal(t, 1, len(sjs.Data.Products))
	assert.Equal(t, "Spread JS v.18", sjs.Data.Products[0].Name)
	assert.Equal(t, "BJIH", sjs.Data.Products[0].Code)

	_ = sjs.Output(os.Stdout)
}

package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// SpreadJS 授权
	lic := "E879948536774266#B12c3MC36MEhXd9hWelJjT5QlRyFFaihDZXxGRQd7ba3mRTRmMh56VDtmMxlkdUlHOotkMjN4Zld6ZiFDerhTb4JTRy2UQBpUeLVTcMJES6ZTN6VnZyo5dKN6KaJVY9B7SLVXOLxmU8F4UxNjdzYndQZ5T4RVQ5pGRld6d8RmbahjdrNXNGpGMXdTUqJUQaZzUapEWMd7L5pEeUl5MkRHbBN7Y74mcYh4UHtUeEV6YUdHN75kehxkcZhlTntkc8lEc4FTY4lXM8ITSnNDVkRTVrFTeRF6dv3CRjhEbQpFajdTW9IzahFEUhNFOyglNSBzUX3kRIR5SvhXR5J4LxJmZ4kDRyUHZxRlQxdEb4YlRTh7S494Kr8kI0IyUiwiIxgTO5MUQ4IjI0ICSiwSNyEjN4AjMzMTM0IicfJye#4Xfd5nIIlkSCJiOiMkIsICOx8idgMlSgQWYlJHcTJiOi8kI1tlOiQmcQJCLiUTMzIzMwAyNwUDM5IDMyIiOiQncDJCLiYDM6ATNyAjMiojIwhXRiwiIx8CMuAjL7ITMiojIz5GRiwiI8+Y9sWY9QmZ0Jyp93uL9hKI0Aqo9Re09cu19R619HWa96mp930b9J0a9iojIh94QiwSZ5JHd0ICb6VkIsIiN6IDN7cjNzUDO4kTO7gjI0ICZJJCL35lI4JXYoNUY4FGRiwiI4VWZoNFdy3GclJlIbpjInxmZiwSZzxWYmpjIyNHZisnOiwmbBJye0ICRiwiI34TUqRlQsJjdzlkVa3UeohFWXZnT7IFdwk4VZ36UE3CRRFFNlR5cyZTRDF6K6JHWyI6Z8s4Y8BHMy3yYrxUaeQZb"

	ss := strings.Split(lic, "#B1")

	prefix := ss[0]
	licData := ss[1]
	log.Print(prefix)
	log.Print(licData)

	//lk := []string{"Lcnee", "iesKy"}
	//pK := "l6/zrbWoSbcLFwEetFh38rH3ErBZE9H+Cqix3R+wTlfA1wD5B+lUcCQn+EJ60I4RGrm0x1sFjkiLWwB0jAn6BWZv0W4WbqAKriOdeoivxDp1Wmjs3qkEDhvbsjPtfvwx2BHil6o+/tDrdMJQSGs18WZm2PoQLQuL+9VhZ4FNRHUQU3Jtioke/OZEGHJOdYVwvCGalzBad6QFOiVbDBQPePpS3++GJzOxN8SN/7lyS5/IdKiy3WJRaVGkB370+HbN6hKraDfUgReLX26yxRaKC/5aWnGAJ2NnWLoGyAGRcwT9dVjo4bcAZNrrA0U9JVKQxaSskhdv2p49XzJkltXx5w=="
	//aC := "B1"

	// 'E879948536774266#B1{"Anl":{"dsr":false,"flg":["ReportSheet","DataChart"]},"Id":"879948536774266","Evl":true,"CNa":"安徽晶奇网络科技股份有限公司","Dms":"127.0.0.1","Exp":"20250606","Crt":"20250507 032315","Prd":[{"N":"Spread JS v.18","C":"BJIH"}]}'
	log.Println()
	//log.Println(lk)
	//log.Println(pK)
	//log.Println(aC)

	u1 := []string{"XRsZ", "HUkJ", "T&g", "Q&w", "GRz1", "JYx3Gb#8Pb5R", "VdgJHc#wJb59", "4LJITMx8UMcA"}
	u2 := []string{"Evl", "Prd", "N", "C", "Dms", "location", "protocol", "127.0.0.1"}
	for i := 0; i < len(u1); i++ {
		if i == 0 {
			break
		}
		s1 := M(u1[i])
		fmt.Printf("%s => %s %t\n", u1[i], s1, s1 == u2[i])
	}

	//licStr := "2c3MC36MEhXd9hWelJjT5QlRyFFaihDZXxGRQd7ba3mRTRmMh56VDtmMxlkdUlHOotkMjN4Zld6ZiFDerhTb4JTRy2UQBpUeLVTcMJES6ZTN6VnZyo5dKN6KaJVY9B7SLVXOLxmU8F4UxNjdzYndQZ5T4RVQ5pGRld6d8RmbahjdrNXNGpGMXdTUqJUQaZzUapEWMd7L5pEeUl5MkRHbBN7Y74mcYh4UHtUeEV6YUdHN75kehxkcZhlTntkc8lEc4FTY4lXM8ITSnNDVkRTVrFTeRF6dv3CRjhEbQpFajdTW9IzahFEUhNFOyglNSBzUX3kRIR5SvhXR5J4LxJmZ4kDRyUHZxRlQxdEb4YlRTh7S494Kr8kI0IyUiwiIxgTO5MUQ4IjI0ICSiwSNyEjN4AjMzMTM0IicfJye#4Xfd5nIIlkSCJiOiMkIsICOx8idgMlSgQWYlJHcTJiOi8kI1tlOiQmcQJCLiUTMzIzMwAyNwUDM5IDMyIiOiQncDJCLiYDM6ATNyAjMiojIwhXRiwiIx8CMuAjL7ITMiojIz5GRiwiI8+Y9sWY9QmZ0Jyp93uL9hKI0Aqo9Re09cu19R619HWa96mp930b9J0a9iojIh94QiwSZ5JHd0ICb6VkIsIiN6IDN7cjNzUDO4kTO7gjI0ICZJJCL35lI4JXYoNUY4FGRiwiI4VWZoNFdy3GclJlIbpjInxmZiwSZzxWYmpjIyNHZisnOiwmbBJye0ICRiwiI34TUqRlQsJjdzlkVa3UeohFWXZnT7IFdwk4VZ36UE3CRRFFNlR5cyZTRDF6K6JHWyI6Z8s4Y8BHMy3yYrxUaeQZb"
	// {
	//    "_r": 1332046125,
	//    "H": "24AC5981",
	//    "S": "N++NtKxSFV4lGqBTqdu2D94fbq/BuExoKTHFOWS0R6X28SaPAak29Y7chZPlHcD/owaQy1kU4dT3gI281yta1tpIxrKgNXYrLazMw4wTceDyKGSHXrm7csAltd3YTxJu/wLXJZS6ZABjQ7W0jF5skv8ZndxwgeDjuATtOVPvv3v3qSAxRlK9uKKpyaRZ+cJwZ2fuv56vHBLq5KyJAAO2E2tm8kx1bggegCc2Kh8yTvIq2kCWma2dSFoZowPDlWd8bhQrFT5N2eyhyuxD3oB3W4lD3iLkc/r0pxcK8gb2Xrv+aCE6rsTe4QQD/DSoYWI0tR7NvWXXhyOZVIsv2lBTjQ==",
	//    "D": {
	//        "Anl": {
	//            "dsr": false,
	//            "flg": [
	//                "ReportSheet",
	//                "DataChart"
	//            ]
	//        },
	//        "Id": "879948536774266",
	//        "Evl": true,
	//        "CNa": "安徽晶奇网络科技股份有限公司",
	//        "Dms": "127.0.0.1",
	//        "Exp": "20250606",
	//        "Crt": "20250507 032315",
	//        "Prd": [
	//            {
	//                "N": "Spread JS v.18",
	//                "C": "BJIH"
	//            }
	//        ]
	//    }
	//}
	ls := M(licData)
	fmt.Println(ls)
}

# jd-api

京东 api 客户端

## for developers

本地调试可以在 go.mod 中添加 

```bash
replace github.com/one-piece-official/jd-api => /Users/xxx/Documents/jd-api
```

## example

```go
package main

import (
	"fmt"

	"github.com/one-piece-official/jd-api"
	dto "github.com/one-piece-official/jd-api/dto"
)

const (
	AppID     = ""
	AppSecret = ""
)

func main() {
	fmt.Println("Hello")

	client := jingdong.NewClient(AppID, AppSecret)

	var res dto.JdUnionOpenUserRegisterValidateResponse

	userStateReq := dto.UserStateReq{
		UserID:     "869706036987313",
		UserIDType: "8",
	}

	if err := client.Query(&dto.UserAccountDeviceInfoRequest{UserStateReq: userStateReq}, &res); err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Registered())
}
```

### 参数说明

userIdType 说明：

* 8(安卓移动设备Imei); 
* 16(苹果移动设备Openudid)；
* 32(苹果移动设备idfa); 
* 64(安卓移动设备imei的md5编码，32位，大写，匹配率略低);
* 128(苹果移动设备idfa的md5编码，32位，大写，匹配率略低); 
* 32768(安卓移动设备oaid)；

## 参考文档

* [京东宙斯开发文档](https://help.jd.com/jos/question-580.html#A6)
* [联盟API接入文档](https://union.jd.com/helpcenter/13246-13247-46301)
* [京东联盟 api 加密 封装](https://github.com/woddp/go-jingdong)
* [jop-go](https://github.com/zhanglianxin/jop-go)
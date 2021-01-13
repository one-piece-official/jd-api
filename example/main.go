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

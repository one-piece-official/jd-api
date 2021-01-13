package dto

type ResponseBody struct {
	Data struct {
		Result string `json:"result"`
		Code   string `json:"code"`
	} `json:"jd_union_open_user_register_validate_response"`
}

type JdUnionOpenUserRegisterValidateResponse struct {
	Code int `json:"code"`
	Data struct {
		UserRes UserRes
	} `json:"data"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

type UserRes struct {
	JdUser int `json:"jdUser"`
}

// 1、京东注册用户 ，2、非京东注册用户（包括未查询到身份的用户）.
func (resp *JdUnionOpenUserRegisterValidateResponse) Available() bool {
	return resp.Data.UserRes.JdUser == 1
}

func (resp *ResponseBody) Success() bool {
	return resp.Data.Code == "0"
}

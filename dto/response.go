package dto

type ResponseBody struct {
	JdUnionOpenUserRegisterValidateResponse JdUnionOpenUserRegisterValidateResponse `json:"jd_union_open_user_register_validate_response"`
}

type JdUnionOpenUserRegisterValidateResponse struct {
	Result string `json:"result"`
	Code   string `json:"code"`
}

type ValidateResult struct {
	Code int `json:"code"`
	Data struct {
		UserResp struct {
			JdUser int `json:"jdUser"`
		} `json:"userResp"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

func (r *JdUnionOpenUserRegisterValidateResponse) Success() bool {
	return r.Code == "0"
}
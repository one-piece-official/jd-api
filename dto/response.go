package dto

type ResponseBody struct {
	JdUnionOpenUserRegisterValidateResponse JdUnionOpenUserRegisterValidateResponse `json:"jd_union_open_user_register_validate_response"`
}

type JdUnionOpenUserRegisterValidateResponse struct {
	ValidateResult ValidateResult `json:"validateResult"`
}

type ValidateResult struct {
	Message string `json:"message"`
	Data    struct {
		UserResp struct {
			JdUser string `json:"jdUser"`
		} `json:"userResp"`
	} `json:"data"`
	Code string `json:"code"`
}

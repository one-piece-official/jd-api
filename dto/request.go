package dto

type RequestBody struct {
	// 公共参数
	AppID     string `json:"app_key" url:"app_key" structs:"app_key"`             // 联盟分配 app key
	Format    string `json:"format" url:"format" structs:"format"`                // 格式
	Method    string `json:"method" url:"method" structs:"method"`                // 接口名称
	SignType  string `json:"sign_method" url:"sign_method" structs:"sign_method"` // 生成签名的算法
	Sign      string `json:"sign,omitempty" url:"sign" structs:"sign,omitempty"`  // 签名
	Timestamp string `json:"timestamp" url:"timestamp" structs:"timestamp"`       // 时间，格式为 yyyy-MM-dd HH:mm:ss，时区为 GMT+8
	Version   string `json:"v" url:"v" structs:"v"`             // 版本
	// 业务参数
	ParamJSON string `json:"param_json" url:"param_json" structs:"param_json"`
}

type UserAccountDeviceInfoRequest struct {
	UserStateReq UserStateReq `json:"userStateReq"` // 请求对象
}

type UserStateReq struct {
	UserIDType string `json:"userIdType"` // 设备类型
	UserID     string `json:"userId"`     // 设备号
}

func (req *UserAccountDeviceInfoRequest) GetMethod() string {
	return "jd.union.open.user.register.validate"
}

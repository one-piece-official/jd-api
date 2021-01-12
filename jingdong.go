package jingdong

const (
	version    = "1.0"
	signMethod = "md5"
	format     = "json"
	gatewayURL = "https://union.jd.com/openplatform/api"
)

type QueryRequest interface {
	GetMethod() string
}

type Client struct {
	appKey string
}

// NewClient 初始化京东客户端.
func NewClient(key string) *Client {
	return &Client{
		appKey: key,
	}
}

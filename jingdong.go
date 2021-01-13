package jingdong

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/google/go-querystring/query"
	"github.com/one-piece-official/jd-api/dto"
	"github.com/one-piece-official/jd-api/utils"
)

const (
	version    = "1.0"
	signType   = "md5"
	format     = "json"
	gatewayURL = "https://router.jd.com/api"
)

var errJDRtaRespFailed = errors.New("jd rta response is failed")

type QueryRequest interface {
	GetMethod() string
}

type Client struct {
	appID     string
	appSecret string
}

// NewClient 初始化京东客户端.
func NewClient(key, secret string) *Client {
	return &Client{
		appID:     key,
		appSecret: secret,
	}
}

func (c *Client) Query(req QueryRequest, response interface{}) (err error) {
	bizContentStr, err := buildBizContent(req)
	if err != nil {
		return err
	}

	// fmt.Printf("bizContentStr is:%s\n", bizContentStr)

	params := dto.RequestBody{
		AppID:     c.appID,
		Format:    format,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		// Timestamp: "11",
		Version:   version,
		SignType:  signType,
		Method:    req.GetMethod(),
		ParamJSON: bizContentStr,
	}

	signString := c.appSecret + composeParameterString(params) + c.appSecret

	// fmt.Printf("sign string is:%s\n\n", signString)
	sign := strings.ToUpper(utils.MD5(signString))
	params.Sign = sign

	// fmt.Printf("sign is:%s\n\n", sign)

	httpClient := http.DefaultClient

	v, _ := query.Values(&params)
	requestURL := fmt.Sprintf("%s?%s", gatewayURL, v.Encode())

	// fmt.Printf("requestURL is %s\n", requestURL)
	resp, err := utils.MakeHTTPClientGet(httpClient, requestURL, bytes.NewBuffer(nil))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	resByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Println(bytes.NewBuffer(resByte))

	var res dto.ResponseBody
	if err = json.Unmarshal(resByte, &res); err != nil {
		return fmt.Errorf("query json unmarshal failed: %w", err)
	}

	if !res.Success() {
		return fmt.Errorf("%w", errJDRtaRespFailed)
	}

	if err = json.Unmarshal([]byte(res.Data.Result), &response); err != nil {
		return fmt.Errorf("query json unmarshal failed: %w", err)
	}

	return nil
}

// composeParameterString 拼装签名用的字符串.
func composeParameterString(params dto.RequestBody) (signString string) {
	// Convert Structs to Map
	requestDataMap := structs.Map(params)

	// 遍历 map 将 key 取出来并按照 ascii 排序
	keys := make([]string, 0, len(requestDataMap))
	for k, v := range requestDataMap {
		keys = append(keys, fmt.Sprintf(`%s%s`, k, v))
	}

	sort.Strings(keys)
	signString = strings.Join(keys, "")

	return signString
}

// buildBizContent 构造请求参数的集合.
func buildBizContent(req QueryRequest) (biz string, err error) {
	byteBiz, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("biz content marshal failed: %w", err)
	}

	biz = string(byteBiz)

	return biz, nil
}

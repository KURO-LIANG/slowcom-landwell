package http

import (
	"fmt"
	"github.com/ddliu/go-httpclient"
	netUrl "net/url"
)

type LandwellClient struct {
	BaseUrl string
}

const (
	USERAGENT       = "slowcom_agent"
	TIMEOUT         = 30
	CONNECT_TIMEOUT = 10
)

// 生成一个http请求客户端
func buildHttpClient() *httpclient.HttpClient {
	return httpclient.NewHttpClient().Defaults(httpclient.Map{
		"opt_useragent":      USERAGENT,
		"opt_timeout":        TIMEOUT,
		"Accept-Encoding":    "gzip,deflate,sdch",
		"opt_connecttimeout": CONNECT_TIMEOUT,
		"OPT_DEBUG":          true,
	})
}

// PostJson json请求
func (s *LandwellClient) PostJson(url string, data interface{}) (response *LandwellResponse, err error) {
	res, err := buildHttpClient().PostJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}

// Post post普通请求
func (s *LandwellClient) Post(url, data interface{}) (response *LandwellResponse, err error) {
	res, err := buildHttpClient().WithHeader("Content-Type", "application/x-www-form-urlencoded").Post(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}

// PutJson json请求
func (s *LandwellClient) PutJson(url, data interface{}) (response *LandwellResponse, err error) {
	res, err := buildHttpClient().PutJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}

// Get get请求
func (s *LandwellClient) Get(url string) (response *LandwellResponse, err error) {
	res, err := buildHttpClient().Get(fmt.Sprintf("%s%s", s.BaseUrl, url), netUrl.Values{})
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}

// Delete 删除请求
func (s *LandwellClient) Delete(url string) (response *LandwellResponse, err error) {
	res, err := buildHttpClient().Delete(fmt.Sprintf("%s%s", s.BaseUrl, url), netUrl.Values{})
	if err != nil {
		return
	}
	response, err = checkResponse(res)
	return
}

package http

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"github.com/kuro-liang/slowcom-landwell/serror"
)

type LandwellBaseResponse struct {
	Success int    `json:"success" description:"错误码"`
	Message string `json:"message" description:"错误信息"`
}
type LandwellResponse struct {
	LandwellBaseResponse
	Data interface{} `json:"data" description:"数据"`
}

// checkResponse 校验请求
func checkResponse(res *httpclient.Response, requestError error) (landwellResponse *LandwellResponse, err error) {
	if requestError != nil {
		return nil, serror.New(405, "请求服务异常：请求超时")
	}
	if res.Response.StatusCode != 200 {
		return nil, serror.New(res.Response.StatusCode, fmt.Sprint(res.Response.StatusCode, " 请求服务异常"))
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return nil, serror.ErrIs数据解析异常
	}
	err = json.Unmarshal(bytes, &landwellResponse)
	if err != nil {
		return nil, serror.ErrIs数据解析异常
	}
	if landwellResponse.Success == 1 {
		return
	} else {
		return landwellResponse, serror.New(landwellResponse.Success, landwellResponse.Message)
	}
}

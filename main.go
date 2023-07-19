package main

import (
	"fmt"
	"slowcom-landwell/app/business/entity"
	"slowcom-landwell/app/business/service"
	"slowcom-landwell/config"
	"slowcom-landwell/http"
)

func main() {
	// 调用demo
	client := &http.LandwellClient{BaseUrl: config.BaseUrl}
	request := service.PlanEvaluationRequest{LandwellClient: client}
	list, err := request.ListByPage(&entity.PlanEvaluationReq{
		DeptNumber:  "",
		StartTime:   "2023-07-10 00:00:00",
		EndTime:     "2023-07-10 23:59:59",
		LineName:    "",
		UserName:    "",
		TeamName:    "",
		DeviceName:  "",
		PlaceName:   "",
		PlanMode:    0,
		PlanState:   "",
		SearchChild: false,
		PageIndex:   1,
		PageSize:    10,
		OnlyCount:   false,
	})
	if err != nil {
		fmt.Println("请求错误：", err)
		return
	}
	fmt.Println("请求成功 <==", list)
}

package main

import (
	"github.com/kuro-liang/slowcom-landwell/app/business/entity"
	"github.com/kuro-liang/slowcom-landwell/app/business/service"
	"github.com/kuro-liang/slowcom-landwell/config"
	"github.com/kuro-liang/slowcom-landwell/http"
)

func main() {
	// 调用demo
	client := &http.LandwellClient{BaseUrl: config.BaseUrl}
	request := service.PlanEvaluationRequest{LandwellClient: client}
	request.ListByPage(&entity.PlanEvaluationReq{
		DeptNumber:  "",
		StartTime:   "2023-07-10 00:00:00",
		EndTime:     "2023-07-10 23:59:59",
		LineName:    "",
		UserName:    "",
		TeamName:    "",
		DeviceName:  "",
		PlaceName:   "",
		PlanMode:    "0",
		PlanState:   "",
		SearchChild: false,
		PageIndex:   1,
		PageSize:    10,
		OnlyCount:   false,
	})
	//request := service.UserModelRequest{LandwellClient: client}
	//user := new(entity.UserModelUpdateReq)
	//user.Id = 112918
	//user.CheckpointName = "吕雪婷1"
	//user.CheckpointCard = "133222366"
	//user.DepartmentNumber = "01"
	//err := request.Edit(user)
	//if err != nil {
	//	fmt.Println("请求错误：", err)
	//	return
	//}
}

package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/kuro-liang/slowcom-landwell/app/business/entity"
	"github.com/kuro-liang/slowcom-landwell/http"
)

type PlanModelRequest struct {
	LandwellClient *http.LandwellClient
}

// List 查询巡更计划列表
func (s *PlanModelRequest) List(req *entity.PlanModelReq) (list []*entity.PlanModelRes, err error) {
	values, _ := query.Values(req)
	queryStr := values.Encode()
	res, err := s.LandwellClient.Get(fmt.Sprintf(`/api/plan?%s`, queryStr))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &list)
	return
}

// GetInfo 查询巡更计划详情
func (s *PlanModelRequest) GetInfo(id int) (planModel *entity.PlanModelRes, err error) {
	res, err := s.LandwellClient.Get(fmt.Sprintf(`/api/plan/%d`, id))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &planModel)
	return
}

// Add 新增巡更计划
func (s *PlanModelRequest) Add(req *entity.PlanModelAddReq) (resData *entity.PlanModelAddRes, err error) {
	res, err := s.LandwellClient.PostJson("/api/plan", req)
	if err != nil {
		return
	}
	resData = new(entity.PlanModelAddRes)
	resData.Id = res.Data
	return
}

// Edit 修改巡更计划
func (s *PlanModelRequest) Edit(req *entity.PlanModelUpdateReq) (err error) {
	_, err = s.LandwellClient.PutJson("/api/plan", req)
	return
}

// Delete 删除巡更计划
func (s *PlanModelRequest) Delete(id int) (err error) {
	_, err = s.LandwellClient.Delete(fmt.Sprintf(`/api/plan/%d`, id))
	return
}

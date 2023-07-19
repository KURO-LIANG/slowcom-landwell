package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/kuro-liang/slowcom-landwell/app/business/entity"
	"github.com/kuro-liang/slowcom-landwell/http"
)

type PlaceModelRequest struct {
	LandwellClient *http.LandwellClient
}

// List 查询巡更地点列表
func (s *PlaceModelRequest) List(req *entity.PlaceModelReq) (list []*entity.PlaceModelRes, err error) {
	values, _ := query.Values(req)
	queryStr := values.Encode()
	res, err := s.LandwellClient.Get(fmt.Sprintf(`/api/place?%s`, queryStr))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &list)
	return
}

// GetInfo 查询巡更地点详情
func (s *PlaceModelRequest) GetInfo(id int) (planModel *entity.PlaceModelRes, err error) {
	res, err := s.LandwellClient.Get(fmt.Sprintf(`/api/place/%d`, id))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &planModel)
	return
}

// Add 新增巡更地点
func (s *PlaceModelRequest) Add(req *entity.PlaceModelAddReq) (resData *entity.PlaceModelAddRes, err error) {
	res, err := s.LandwellClient.PostJson("/api/place", req)
	if err != nil {
		return
	}
	resData.Id = res.Data
	return
}

// Edit 修改巡更地点
func (s *PlaceModelRequest) Edit(req *entity.PlaceModelUpdateReq) (err error) {
	_, err = s.LandwellClient.PutJson("/api/place", req)
	return
}

// Delete 删除巡更地点
func (s *PlaceModelRequest) Delete(id int) (err error) {
	_, err = s.LandwellClient.Delete(fmt.Sprintf(`/api/place/%d`, id))
	return
}

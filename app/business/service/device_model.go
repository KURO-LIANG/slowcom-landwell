package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"slowcom-landwell/app/business/entity"
	"slowcom-landwell/http"
)

type DeviceModelRequest struct {
	LandwellClient *http.LandwellClient
}

// List 查询巡检器列表
func (s *DeviceModelRequest) List(req *entity.DeviceModelReq) (list []*entity.DeviceModelRes, err error) {
	values, _ := query.Values(req)
	queryStr := values.Encode()
	res, err := s.LandwellClient.Get(fmt.Sprintf(`/api/device?%s`, queryStr))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &list)
	return
}

// GetInfo 查询巡检器详情
func (s *DeviceModelRequest) GetInfo(id int) (planModel *entity.DeviceModelRes, err error) {
	res, err := s.LandwellClient.Get(fmt.Sprintf(`/api/device/%d`, id))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &planModel)
	return
}

// Add 新增巡检器
func (s *DeviceModelRequest) Add(req *entity.DeviceModelAddReq) (resData *entity.DeviceModelAddRes, err error) {
	res, err := s.LandwellClient.PostJson("/api/device", req)
	if err != nil {
		return
	}
	resData.Id = res.Data
	return
}

// Edit 修改巡检器
func (s *DeviceModelRequest) Edit(req *entity.DeviceModelUpdateReq) (err error) {
	_, err = s.LandwellClient.PutJson("/api/device", req)
	return
}

// Delete 删除巡检器
func (s *DeviceModelRequest) Delete(id int) (err error) {
	_, err = s.LandwellClient.Delete(fmt.Sprintf(`/api/device/%d`, id))
	return
}

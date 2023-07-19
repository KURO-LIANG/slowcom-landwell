package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"slowcom-landwell/app/business/entity"
	"slowcom-landwell/http"
)

type UserModelRequest struct {
	LandwellClient *http.LandwellClient
}

// List 查询人员列表
func (s *UserModelRequest) List(req *entity.UserModelReq) (list []*entity.UserModelRes, err error) {
	values, _ := query.Values(req)
	queryStr := values.Encode()
	res, err := s.LandwellClient.Get(fmt.Sprintf(`/api/user?%s`, queryStr))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &list)
	return
}

// GetInfo 查询人员详情
func (s *UserModelRequest) GetInfo(id int) (planModel *entity.UserModelRes, err error) {
	res, err := s.LandwellClient.Get(fmt.Sprintf(`/api/user/%d`, id))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &planModel)
	return
}

// Add 新增人员
func (s *UserModelRequest) Add(req *entity.UserModelAddReq) (resData *entity.UserModelAddRes, err error) {
	res, err := s.LandwellClient.PostJson("/api/user", req)
	if err != nil {
		return
	}
	resData.Id = res.Data
	return
}

// Edit 修改人员
func (s *UserModelRequest) Edit(req *entity.UserModelUpdateReq) (err error) {
	_, err = s.LandwellClient.PutJson("/api/user", req)
	return
}

// Delete 删除人员
func (s *UserModelRequest) Delete(id int) (err error) {
	_, err = s.LandwellClient.Delete(fmt.Sprintf(`/api/user/%d`, id))
	return
}

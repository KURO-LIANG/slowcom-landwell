package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/kuro-liang/slowcom-landwell/app/business/entity"
	"github.com/kuro-liang/slowcom-landwell/http"
)

type LineModelRequest struct {
	LandwellClient *http.LandwellClient
}

// List 查询线路列表
func (s *LineModelRequest) List(req *entity.LineModelReq) (list []*entity.LineModelRes, err error) {
	values, _ := query.Values(req)
	queryStr := values.Encode()
	url := "/api/line"
	if req.SearchPlace {
		url = "/api/lineextd"
	}

	res, err := s.LandwellClient.Get(fmt.Sprintf(`%s?%s`, url, queryStr))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &list)
	return
}

// GetInfo 查询线路详情 searchPlace:是否需要路线地点列表
func (s *LineModelRequest) GetInfo(id int, searchPlace bool) (planModel *entity.LineModelRes, err error) {
	url := "/api/line"
	if searchPlace {
		url = "/api/lineextd"
	}
	res, err := s.LandwellClient.Get(fmt.Sprintf(`%s/%d`, url, id))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &planModel)
	return
}

// Add 新增线路
func (s *LineModelRequest) Add(req *entity.LineModelAddReq) (resData *entity.LineModelAddRes, err error) {
	res, err := s.LandwellClient.PostJson("/api/line", req)
	if err != nil {
		return
	}
	resData.Id = res.Data
	return
}

// AddExtd 新增线路详情
func (s *LineModelRequest) AddExtd(req *entity.LineExtdAddReq) (resData *entity.LineModelAddRes, err error) {
	res, err := s.LandwellClient.PostJson("/api/lineextd", req)
	if err != nil {
		return
	}
	resData.Id = res.Data
	return
}

// Edit 修改线路
func (s *LineModelRequest) Edit(req *entity.LineModelUpdateReq) (err error) {
	_, err = s.LandwellClient.PutJson("/api/line", req)
	return
}

// EditExtd 修改线路详情
func (s *LineModelRequest) EditExtd(req *entity.LineExtdUpdateReq) (err error) {
	_, err = s.LandwellClient.PutJson("/api/lineextd", req)
	return
}

// Delete 删除线路详情
func (s *LineModelRequest) Delete(id int) (err error) {
	_, err = s.LandwellClient.Delete(fmt.Sprintf(`/api/line/%d`, id))
	return
}

// DeleteExtd 删除线路详情
func (s *LineModelRequest) DeleteExtd(id int) (err error) {
	_, err = s.LandwellClient.Delete(fmt.Sprintf(`/api/lineextd/%d`, id))
	return
}

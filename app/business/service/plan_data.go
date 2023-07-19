package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"slowcom-landwell/app/business/entity"
	"slowcom-landwell/http"
)

type PlanDataRequest struct {
	LandwellClient *http.LandwellClient
}

// ListByPage 分页查询巡更历史记录
func (s *PlanDataRequest) ListByPage(req *entity.PatrolDataReq) (list []*entity.PatrolDataRes, err error) {
	values, _ := query.Values(req)
	queryStr := values.Encode()
	res, err := s.LandwellClient.Get(fmt.Sprintf(`/api/data?%s`, queryStr))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &list)
	return
}

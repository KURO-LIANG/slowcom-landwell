package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"slowcom-landwell/app/business/entity"
	"slowcom-landwell/http"
)

type PlanEvaluationRequest struct {
	LandwellClient *http.LandwellClient
}

// ListByPage 分页查询巡更考核结果
func (s *PlanEvaluationRequest) ListByPage(req *entity.PlanEvaluationReq) (list []*entity.PlanEvaluationRes, err error) {
	values, _ := query.Values(req)
	queryStr := values.Encode()
	res, err := s.LandwellClient.Get(fmt.Sprintf(`/api/evaluation?%s`, queryStr))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &list)
	return
}

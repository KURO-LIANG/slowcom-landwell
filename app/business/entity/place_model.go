package entity

// PlaceModelReq 查询巡更地点列表参数
type PlaceModelReq struct {
	DeptNumber  string `json:"deptNumber,omitempty"`  // 部门编号，不填则查询所有部门
	PlaceCard   string `json:"placeCard,omitempty"`   // 地点卡号，非模糊查询
	PlaceName   string `json:"placeName,omitempty"`   // 地点名称，非模糊查询
	SearchChild bool   `json:"searchChild,omitempty"` // 是否查询子部门，true为是，false为否
}

// PlaceModelRes 巡更地点列表返回
type PlaceModelRes struct {
	ID               int    `json:"id"`                // 自增长ID
	CheckPointNumber int    `json:"CheckPoint_Number"` // 地点编号，每个部门从1开始累加
	CheckPointName   string `json:"CheckPoint_Name"`   // 地点名称
	CheckPointCard   string `json:"CheckPoint_Card"`   // 地点卡号，8位16进制
	DepartmentName   string `json:"DepartMent_Name"`   // 部门名称
	DepartmentNumber string `json:"DepartMent_Number"` // 部门编号
}

// PlaceModelAddReq 添加巡更地点参数
type PlaceModelAddReq struct {
	CheckpointName   string `json:"Checkpoint_Name,omitempty"`   // 地点名称（必填）
	CheckpointCard   string `json:"Checkpoint_Card,omitempty"`   // 地点钮号（必填）
	DepartmentNumber string `json:"Department_Number,omitempty"` // 部门编码（必填）
}

// PlaceModelAddRes 添加巡更地点返回
type PlaceModelAddRes struct {
	Id interface{} `json:"id"` // 巡更地点id
}

// PlaceModelUpdateReq 修改巡更地点参数
type PlaceModelUpdateReq struct {
	Id int `json:"id"` // 巡更地点id（必填）
	PlaceModelAddReq
}

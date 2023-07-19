package entity

// UserModelReq 查询人员列表参数
type UserModelReq struct {
	DeptNumber string `json:"deptNumber,omitempty"` // 部门编号，不填则查询所有部门
	UserCard   string `json:"userCard,omitempty"`   // 人员卡号，非模糊查询
	UserName   string `json:"userName,omitempty"`   // 人员名称，非模糊查询
}

// UserModelRes 人员列表返回
type UserModelRes struct {
	ID               int    `json:"id"`                // 自增长ID
	CheckPointNumber int    `json:"CheckPoint_Number"` // 人员编号，每个部门从1开始累加
	CheckPointName   string `json:"CheckPoint_Name"`   // 人员名称
	CheckPointCard   string `json:"CheckPoint_Card"`   // 人员卡号
	DepartmentName   string `json:"DepartMent_Name"`   // 部门名称
	DepartmentNumber string `json:"DepartMent_Number"` // 部门编号
}

// UserModelAddReq 添加人员参数
type UserModelAddReq struct {
	CheckpointName   string `json:"Checkpoint_Name,omitempty"`   // 人员名称（必填）
	CheckpointCard   string `json:"Checkpoint_Card,omitempty"`   // 人员卡号（必填）
	DepartmentNumber string `json:"Department_Number,omitempty"` // 部门编码（必填）
}

// UserModelAddRes 添加人员返回
type UserModelAddRes struct {
	Id interface{} `json:"id"` // 人员id
}

// UserModelUpdateReq 修改人员参数
type UserModelUpdateReq struct {
	Id int `json:"id"` // 人员id（必填）
	UserModelAddReq
}

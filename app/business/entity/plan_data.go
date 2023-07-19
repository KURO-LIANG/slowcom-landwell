package entity

// PatrolDataReq 历史记录请求参数
type PatrolDataReq struct {
	DeptNumber  string `json:"deptNumber,omitempty"`  // 部门编号，不填则查询所有部门
	StartTime   string `json:"startTime"`             // 查询开始时间段yyyy-MM-dd HH:mm:ss（必填）
	EndTime     string `json:"endTime"`               // 查询结束时间段yyyy-MM-dd HH:mm:ss（必填）
	UserName    string `json:"userName,omitempty"`    // 人员名称，非模糊查询
	TeamName    string `json:"teamName,omitempty"`    // 班组名称，非模糊查询
	DeviceName  string `json:"deviceName,omitempty"`  // 设备名称，非模糊查询
	PlaceName   string `json:"placeName,omitempty"`   // 地点名称，非模糊查询
	SearchChild bool   `json:"searchChild,omitempty"` // 是否查询子部门，true为是，false为否
	PageIndex   int    `json:"pageIndex,omitempty"`   // 分页参数，第几页（从1开始）
	PageSize    int    `json:"pageSize,omitempty"`    // 分页参数，每页几条（最少为1）
	OnlyCount   bool   `json:"onlyCount,omitempty"`   // 是否只查询结果数量，一般统计数量时设置为True，默认推荐False
}

// PatrolDataRes 历史记录返回
type PatrolDataRes struct {
	ID               int    `json:"id"`                // 自增长ID
	PatrolTime       string `json:"Patrol_Time"`       // 巡检时间
	UserName         string `json:"User_Name"`         // 巡检人员
	TeamName         string `json:"Team_Name"`         // 所属班组
	PlaceName        string `json:"Place_Name"`        // 巡检地点
	DeviceName       string `json:"Device_Name"`       // 巡检设备
	DepartmentName   string `json:"DepartMent_Name"`   // 部门名称
	DepartmentNumber string `json:"DepartMent_Number"` // 部门编号
	EventCount       int    `json:"Event_Count"`       // 事件个数
}

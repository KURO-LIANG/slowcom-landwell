package entity

// PlanEvaluationReq 巡更计划考核查询
type PlanEvaluationReq struct {
	DeptNumber  string `json:"deptNumber"`  // 部门编号，不填则查询所有部门
	StartTime   string `json:"startTime"`   // 查询开始时间段，格式为 yyyy-MM-dd HH:mm:ss（必填）
	EndTime     string `json:"endTime"`     // 查询结束时间段，格式为 yyyy-MM-dd HH:mm:ss（必填）
	LineName    string `json:"lineName"`    // 线路名称，非模糊查询
	UserName    string `json:"userName"`    // 人员名称，非模糊查询
	TeamName    string `json:"teamName"`    // 班组名称，非模糊查询
	DeviceName  string `json:"deviceName"`  // 设备名称，非模糊查询
	PlaceName   string `json:"placeName"`   // 地点名称，非模糊查询
	PlanMode    string `json:"planMode"`    // 计划模式，0常规1排班2星期3周期
	PlanState   string `json:"planState"`   // 考核状态，（准时、未到、错序）
	SearchChild bool   `json:"searchChild"` // 是否查询子部门，true为是，false为否
	PageIndex   int    `json:"pageIndex"`   // 分页参数，第几页（从1开始）
	PageSize    int    `json:"pageSize"`    // 分页参数，每页几条（最少为1）
	OnlyCount   bool   `json:"onlyCount"`   // 是否只查询结果数量，一般统计数量时设置为True，默认推荐False
}

// PlanEvaluationRes 巡更计划考核查询结果
type PlanEvaluationRes struct {
	ID               int    `json:"id"`                // 自增长ID
	PlanStartTime    string `json:"Plan_StartTime"`    // 计划开始时间，格式yyyy-MM-dd HH:mm:ss
	PlanEndTime      string `json:"Plan_EndTime"`      // 计划结束时间，格式yyyy-MM-dd HH:mm:ss
	PatrolTime       string `json:"Patrol_Time"`       // 巡检时间，格式yyyy-MM-dd HH:mm:ss
	LineName         string `json:"Line_Name"`         // 线路名称
	PlaceName        string `json:"Place_Name"`        // 地点名称
	BindUser         string `json:"Bind_User"`         // 计划绑定人员
	UserName         string `json:"User_Name"`         // 实际巡检人员
	BindDevice       string `json:"Bind_Device"`       // 计划绑定设备
	DeviceName       string `json:"Device_Name"`       // 实际巡检设备
	DataID           string `json:"Data_ID"`           // 参与考核原始记录ID
	PlanState        string `json:"Plan_State"`        // 考核状态（准时，未到，错序）
	BindTeam         string `json:"Bind_Team"`         // 排班计划，绑定班组
	TeamName         string `json:"Team_Name"`         // 实际巡检班组
	PlanType         string `json:"Plan_Type"`         // 计划类型（有序，无序）
	LineOrder        string `json:"Line_Order"`        // 地点在线路中的顺序
	PlanMode         string `json:"Plan_Mode"`         // 计划模式，0常规，1排班，2星期，3周期
	PlanID           string `json:"Plan_ID"`           // 生成该考核信息的计划ID
	ErrorTime        string `json:"Error_Time"`        // 误差时间
	DepartmentName   string `json:"DepartMent_Name"`   // 部门名称
	DepartmentNumber string `json:"DepartMent_Number"` // 部门编号
	EventCount       int    `json:"Event_Count"`       // 事件个数
}

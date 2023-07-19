package entity

// PlanModelReq 查询巡更计划列表参数
type PlanModelReq struct {
	DeptNumber  string `json:"deptNumber,omitempty"`  // 部门编号，不填则查询所有部门
	LineName    string `json:"lineName,omitempty"`    // 线路名称，非模糊查询
	PlanMode    int    `json:"planMode,omitempty"`    // 计划模式，0常规计划 1排班计划（未启用） 2星期计划 3周期计划 4每周计划 5每月计划
	SearchChild bool   `json:"searchChild,omitempty"` // 是否查询子部门，true为是，false为否
}

// PlanModelRes 巡更计划列表返回
type PlanModelRes struct {
	ID               int    `json:"id"`                // 自增长ID
	LineID           int    `json:"Line_ID"`           // 线路ID
	LineName         string `json:"Line_Name"`         // 线路名称
	PlanStartTime    string `json:"Plan_StartTime"`    // 计划开始时间 格式HH:mm:ss
	PlanEndTime      string `json:"Plan_EndTime"`      // 计划结束时间 格式HH:mm:ss
	PlanType         string `json:"Plan_Type"`         // 计划类型（有序，无序）
	ClassID          *int   `json:"Class_ID"`          // 班次ID（未启用）
	ClassName        string `json:"Class_Name"`        // Plan_Mode为1时有用，班次名（早班，中班，晚班）（未启用）
	ExcuteWeek       string `json:"Excute_Week"`       // Plan_Mode为2时有用，执行星期(格式0,1,2,3,4,5,6代表周日到周六都执行本计划）
	ExcuteCycle      string `json:"Excute_Cycle"`      // Plan_Mode为3时有用，执行周期，连续多少天执行本计划
	RestCycle        string `json:"Rest_Cycle"`        // Plan_Mode为3时有用，休息周期，连续多少天不执行本计划
	ErrorTime        string `json:"Error_Time"`        // 误差时间，允许计划开始结束时间有一定误差分钟
	UserID           *int   `json:"User_ID"`           // 计划绑定人员ID
	UserName         string `json:"User_Name"`         // 计划绑定人员
	DeviceID         *int   `json:"Device_ID"`         // 计划绑定设备ID
	DeviceName       string `json:"Device_Name"`       // 计划绑定设备
	StartDate        string `json:"Start_Date"`        // Plan_Mode为3时有用，周期计划开始日期
	PlanMode         string `json:"Plan_Mode"`         // 计划模式，0常规，1排班，2星期，3周期
	DepartmentNumber string `json:"DepartMent_Number"` // 部门编号
	DepartmentName   string `json:"DepartMent_Name"`   // 部门名称
	ExcuteTimes      int    `json:"Excute_Times"`      // 执行次数
}

// PlanModelAddReq 添加巡更计划参数
type PlanModelAddReq struct {
	LineID           int    `json:"Line_ID"`                // 线路ID（必填）
	PlanStartTime    string `json:"Plan_StartTime"`         // 计划开始时间HH:mm（必填）
	PlanEndTime      string `json:"Plan_EndTime"`           // 计划结束时间HH:mm（必填）
	PlanType         string `json:"Plan_Type"`              // 计划类型，“有序”或“无序”（必填）
	ClassID          *int   `json:"Class_ID,omitempty"`     // 班次ID（未启用）
	ErrorTime        int    `json:"Error_Time"`             // 误差时间分钟数，默认填0即可（必填）
	DepartmentNumber string `json:"Department_Number"`      // 部门编码（必填）
	ExcuteWeek       string `json:"Excute_Week,omitempty"`  // 执行星期，例如每天都执行则为0,1,2,3,4,5,6（当Plan_Mode为2时必填，其余情况填""）
	ExcuteCycle      int    `json:"Excute_Cycle,omitempty"` // 执行周期，例如上三休二，填3（当Plan_Mode为3时必填，其余情况填0）
	RestCycle        int    `json:"Rest_Cycle,omitempty"`   // 休息周期，例如上三休二，填2（当Plan_Mode为3时必填，其余情况填0）
	UserID           *int   `json:"User_ID,omitempty"`      // 计划绑定人员ID，不绑定人员则填null或-1
	DeviceID         *int   `json:"Device_ID,omitempty"`    // 计划绑定设备ID，不绑定设备则填null或-1
	StartDate        string `json:"Start_Date,omitempty"`   // 周期开始日期yyyy-MM-dd（当Plan_Mode为3时必填，其余情况填""）
	PlanMode         int    `json:"Plan_Mode"`              // 计划模式，0常规计划1排班计划（未启用）2星期计划3周期计划4每周计划5每月计划
	ExcuteTimes      *int   `json:"Excute_Times,omitempty"` // 执行次数（当Plan_Mode为4或5时必填并启用，其余情况填任何数字均为1）
}

// PlanModelAddRes 添加巡更计划返回
type PlanModelAddRes struct {
	Id interface{} `json:"id"` // 计划id
}

// PlanModelUpdateReq 修改巡更计划参数
type PlanModelUpdateReq struct {
	Id int `json:"id"` // 线路计划ID（必填）
	PlanModelAddReq
}

package entity

// LineModelReq 查询线路列表参数
type LineModelReq struct {
	DeptNumber  string `json:"deptNumber,omitempty"`  // 部门编号，不填则查询所有部门
	LineName    string `json:"lineName,omitempty"`    // 线路名称，非模糊查询
	SearchChild bool   `json:"searchChild,omitempty"` // 是否查询子部门，true为是，false为否
	SearchPlace bool   `json:"searchPlace,omitempty"` // 是否查询地点列表，true为是，false为否
}

// LineModelRes 查询线路返回
type LineModelRes struct {
	ID               int             `json:"id"`                     // 自增长ID
	LineName         string          `json:"Line_Name"`              // 线路名称
	DepartmentNumber string          `json:"DepartMent_Number"`      // 部门编号
	DepartmentName   string          `json:"DepartMent_Name"`        // 部门名称
	LineExtdList     []LineExtdModel `json:"LineExtdList,omitempty"` // 线路包含地点列表 (if needed)
}

type LineExtdModel struct {
	ID            int    `json:"id"`             // 自增长ID
	LineID        int    `json:"Line_ID"`        // 线路ID
	PlaceID       int    `json:"Place_ID"`       // 地点ID
	PlaceName     string `json:"Place_Name"`     // 地点名称
	LineOrder     string `json:"Line_Order"`     // 地点在线路中的巡检顺序
	NextPlaceTime string `json:"NextPlace_Time"` // 到达下一地点时间，7000EF型号产品专用，其他型号产品无效
	StayTime      string `json:"Stay_Time"`      // 要求停留时间
	MapX          string `json:"Map_X"`          // 在jpg地图中的X坐标，百分比
	MapY          string `json:"Map_Y"`          // 在jpg地图中的y坐标，百分比
}

// LineModelAddReq 添加线路参数
type LineModelAddReq struct {
	LineName         string `json:"line_Name,omitempty"`         // 线路名称（必填）
	DepartmentNumber string `json:"Department_Number,omitempty"` // 部门编码（必填）
}

// LineModelAddRes 添加线路返回
type LineModelAddRes struct {
	Id interface{} `json:"id"` // 线路id
}

// LineModelUpdateReq 修改线路参数
type LineModelUpdateReq struct {
	Id int `json:"id"` // 线路id（必填）
	LineModelAddReq
}

// LineExtdAddReq 添加线路详情参数
type LineExtdAddReq struct {
	LineID        int    `json:"Line_ID"`        // 线路ID（必填）
	PlaceID       int    `json:"Place_ID"`       // 地点ID（必填）
	LineOrder     string `json:"Line_Order"`     // 线路中的地点顺序（必填）
	NextPlaceTime string `json:"NextPlace_Time"` // 到达下一地点时间，7000EF产品才有效，其他产品无效，默认填0即可（必填）
	StayTime      string `json:"Stay_Time"`      // 停留时间分钟数，默认填0即可（必填）
	MapX          string `json:"Map_X"`          // JPG地图坐标X比值，默认填0即可（必填）
	MapY          string `json:"Map_Y"`          // JPG地图坐标Y比值，默认填0即可（必填）
}

// LineExtdUpdateReq 修改线路详情参数
type LineExtdUpdateReq struct {
	Id int `json:"id"` // 线路id（必填）
	LineExtdAddReq
}

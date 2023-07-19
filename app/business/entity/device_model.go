package entity

// DeviceModelReq 查询巡检器列表参数
type DeviceModelReq struct {
	DeptNumber   string `json:"deptNumber,omitempty"`   // 部门编号，不填则查询所有部门
	DeviceNumber string `json:"deviceNumber,omitempty"` // 设备号码，非模糊查询
	DeviceName   string `json:"deviceName,omitempty"`   // 设备名称，非模糊查询
	SearchChild  bool   `json:"searchChild,omitempty"`  // 是否查询子部门，true为是，false为否
}

// DeviceModelRes 查询巡检器列表返回
type DeviceModelRes struct {
	ID               int    `json:"id"`                // 自增长ID
	DeviceName       string `json:"Device_Name"`       // 设备名称
	DeviceNumber     string `json:"Device_Number"`     // 设备号码
	DepartmentNumber string `json:"DepartMent_Number"` // 部门编号
	DepartmentName   string `json:"DepartMent_Name"`   // 部门名称
}

// DeviceModelAddReq 添加巡检器参数
type DeviceModelAddReq struct {
	DeviceName       string `json:"Device_Name,omitempty"`       // 设备名称（必填）
	DeviceNumber     string `json:"Device_Number,omitempty"`     // 设备唯一识别号码（必填）
	DepartmentNumber string `json:"Department_Number,omitempty"` // 部门编码（必填）
}

// DeviceModelAddRes 添加巡检器返回
type DeviceModelAddRes struct {
	Id interface{} `json:"id"` // 设备id
}

// DeviceModelUpdateReq 修改巡检器参数
type DeviceModelUpdateReq struct {
	Id int `json:"id"` // 巡检器ID（必填）
	DeviceModelAddReq
}

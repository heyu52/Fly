package contract



// Staff 通过 paas、移动网关、rio3准入网关认证的用户信息
type Staff struct {
	// 通过paas、移动网关认证的用户有效
	StaffID string `json:"staff_id,omitempty"`
	// 通过paas、移动网关认证的用户有效
	StaffName string `json:"staff_name,omitempty"`

	// TODO 以下信息需要重构
	// 通过rio3 准入网关认证的用户有效
	UID string `json:"uid,omitempty"`
	// 通过rio3 准入网关认证的用户有效
	UInfo string `json:"uinfo,omitempty"`
	// 通过rio3 准入网关认证的用户有效
	Ext interface{} `json:"ext,omitempty"`
}
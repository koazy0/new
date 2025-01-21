package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 // 管理员
	PermissionUser        Role = 2 // 普通登陆人
	PermissionVisiter     Role = 3 // 游客
	PermissionDisableUser Role = 4 // 被禁用的用户
)

func (s Role) MarshalJson() ([]byte, error) {
	return json.Marshal(s.String())

}
func (s Role) String() string {
	switch s {
	case PermissionAdmin:
		return "管理员"
	case PermissionUser:
		return "用户"
	case PermissionVisiter:
		return "游客"
	case PermissionDisableUser:
		return "被禁言"
	default:
		return "其他"
	}
}

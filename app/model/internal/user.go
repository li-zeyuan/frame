package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       int         `orm:"id,primary" json:"id"`       // 用户ID
	Passport string      `orm:"passport"   json:"passport"` // 用户账号
	Password string      `orm:"password"   json:"password"` // 用户密码
	Nickname string      `orm:"nickname"   json:"nickname"` // 用户昵称
	CreateAt *gtime.Time `orm:"create_at"  json:"createAt"` // 创建时间
	UpdateAt *gtime.Time `orm:"update_at"  json:"updateAt"` // 更新时间
}

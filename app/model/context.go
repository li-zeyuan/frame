package model

import "github.com/gogf/gf/net/ghttp"

const (
	ContextKey = "contextKey"
)

type Context struct {
	Session *ghttp.Session
	User    *ContextUser
}

type ContextUser struct {
	Id       int    //　用户id
	Passport string // 账号
	Nickname string // 用户名称
}

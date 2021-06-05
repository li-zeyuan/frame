package service

import (
	"net/http"

	"frame/app/model"
	"github.com/gogf/gf/net/ghttp"
)

// 中间件管理服务
var Middleware = middlewareService{}

type middlewareService struct{}

func (s *middlewareService) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context.Init(r, customCtx)

	if user := Session.GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:       user.Id,
			Passport: user.Passport,
			Nickname: user.Nickname,
		}
	}

	r.Middleware.Next()
}

// 登录鉴权中间件
func (s *middlewareService) Auth(r *ghttp.Request) {
	if User.IsSingedIn(r.Context()) {
		r.Middleware.Next()
	}

	r.Response.WriteStatus(http.StatusForbidden)
}

// 允许接口跨域请求
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

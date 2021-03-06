package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/li-zeyuan/frame/app/api"
	"github.com/li-zeyuan/frame/app/service"
)

func init() {
	s := g.Server()
	s.Group("/api/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,
		)
		group.ALL("/user", api.User)
		group.Group("/api/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.Auth)
			group.ALL("/user/profile", api.User.Profile)
		})
	})
}

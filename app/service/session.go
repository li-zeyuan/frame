package service

import (
	"context"

	"github.com/li-zeyuan/frame/app/model"
)

// session管理服务
var Session = sessionService{}

type sessionService struct {}

const (
	sessionKeyUser = "SessionKeyUser"
)

func (s *sessionService)SetUser(ctx context.Context, user *model.User) error   {
	return Context.Get(ctx).Session.Set(sessionKeyUser, user)
}

func (s *sessionService)GetUser(ctx context.Context) *model.User {
	customCtx := Context.Get(ctx)
	if customCtx != nil    {
		if v := customCtx.Session.GetVar(sessionKeyUser); !v.IsNil() {
			var user *model.User
			_ = v.Struct(&user)
			return user
		}
	}

	return nil
}

func (s *sessionService)RemoveUser(ctx context.Context)error  {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}

	return nil
}
package service

import (
	"context"

	"frame/app/model"
	"github.com/gogf/gf/net/ghttp"
)

// 上下文管理服务
var Context = contextService{}

type contextService struct{}

func (s *contextService) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(model.ContextKey, customCtx)
}

func (s *contextService) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}

	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}

	return nil
}

func (s *contextService) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/li-zeyuan/frame/app/dao"
	"github.com/li-zeyuan/frame/app/model"
)

var User = userService{}

type userService struct{}

// 用户注册
func (s *userService) SignUp(r *model.UserServiceSignUpReq) error {
	if r.Nickname == "" {
		r.Nickname = r.Passport
	}

	if !s.CheckPassport(r.Passport) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", r.Passport))
	}
	if !s.CheckNickName(r.Nickname) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", r.Nickname))
	}

	if _, err := dao.User.Save(r); err != nil {
		return err
	}

	return nil
}

func (s *userService) SignIn(ctx context.Context, passport, password string) error {
	user, err := dao.User.FindOne("passport=? and password=?", passport, password)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("账号密码错误")
	}

	if err := Session.SetUser(ctx, user); err != nil {
		return err
	}

	Context.SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
	})
	return nil
}

func (s *userService) SingOut(ctx context.Context) error {
	return Session.RemoveUser(ctx)
}

func (s *userService) CheckPassport(passport string) bool {
	count, err := dao.User.FindCount(dao.User.Columns.Passport, passport)
	if err != nil {
		return false
	}

	return count == 0
}

func (s *userService) CheckNickName(nickname string) bool {
	count, err := dao.User.FindCount(dao.User.Columns.Nickname)
	if err != nil {
		return false
	}

	return count == 0
}

func (s *userService) IsSingedIn(ctx context.Context) bool {
	if v := Context.Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// 获得用户信息详情
func (s *userService) GetProfile(ctx context.Context) *model.User {
	return Session.GetUser(ctx)
}

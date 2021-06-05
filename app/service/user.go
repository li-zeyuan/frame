package service

import (
	"context"
	"errors"
	"fmt"

	"frame/app/dao"
	"frame/app/model"
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

func (s *userService)IsSingedIn(ctx context.Context)  bool{
	if v := Context.Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

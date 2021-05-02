package service

import "frame/app/model"

var User = userService{}

type userService struct{}

// 用户注册
func (s *userService) SignUp(r *model.UserServiceSignUpReq) error {

	return nil
}

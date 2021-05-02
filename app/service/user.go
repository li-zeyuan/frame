package service

import "frame/app/model"

var User = userService{}

type userService struct{}

func (s *userService) SignUp(r *model.UserServiceSignUpReq) error {
	return nil
}

package dao

import (
	"frame/app/dao/internal"
)

type userDao struct {
	internal.UserDao
}

var (
	// User is globally public accessible object for table user operations.
	User = userDao{
		internal.User,
	}
)

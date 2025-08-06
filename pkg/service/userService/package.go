package userService

import "github.com/sirupsen/logrus"

func init() {
	logrus.Infof("xuanyuan::UserService initialized")
}

var (
	UserService = &userService{}
)

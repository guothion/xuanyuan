package stall

import "github.com/sirupsen/logrus"

func init() {
	logrus.Infof("xuanyuan::StallService initialized")
}

var (
	StallsService = &stallsService{}
)

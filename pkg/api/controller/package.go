package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var controllers = map[string]Controller{}

// now we defined Controller
type Controller interface {
	RegisterRouter(*gin.Engine)
	BasePath() string
}

func Init(engine *gin.Engine) {
	for k, v := range controllers {
		logrus.Debugf("reginster routes %s", k)
		v.RegisterRouter(engine)
	}
}

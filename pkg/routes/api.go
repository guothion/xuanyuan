package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var routes = map[string]Routes{}

// 定义一下 routes
type Routes interface {
	RegisterRouter(group *gin.RouterGroup)
	BasePath() string
}

func InitRoutes(engine *gin.RouterGroup) {
	for k, v := range routes {
		logrus.Debugf("reginster routes %s", k)
		v.RegisterRouter(engine)
	}
}

func SetApiGroupRoutes(router *gin.RouterGroup) {
	InitRoutes(router)
}

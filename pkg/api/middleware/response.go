package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/guothion/xuanyuan/pkg/util"
	"github.com/sirupsen/logrus"
	gormLogger "gorm.io/gorm/logger"
	"net/http"
	"runtime/debug"
	"strings"
)

func ResponseData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func RespondCreated(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, common.StatusOk)
}

func RespondUpdated(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.StatusOk)
}

func RespondForbidden(ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, common.StatusForbidden)
}

func RespondBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, common.NewBadRequestError(err.Error()))
}

func RespondFailure(ctx *gin.Context, err error) {
	if strings.Contains(err.Error(), gormLogger.ErrRecordNotFound.Error()) {
		ctx.JSON(http.StatusNotFound, common.NewNotFoundError(err.Error()))
		return
	}

	var (
		response *common.HTTPError
		ok       bool
	)
	if response, ok = err.(*common.HTTPError); !ok {
		trace := util.RandString(64)
		response = common.NewInternalServerError("Server Error. Trace ID: %s", trace)
	}
	ctx.JSON(response.Code, err)
	logrus.Errorf("[Internal Server Error] [%s]: \n%v", err, string(debug.Stack()))
}

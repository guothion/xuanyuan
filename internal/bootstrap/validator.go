package bootstrap

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/guothion/xuanyuan/internal/api/common/request"
	"github.com/guothion/xuanyuan/internal/util"
)

func InitializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 自定义验证器
		request.RegisterCustomValidators(v)
		// 注册自定义验证器
		_ = v.RegisterValidation("mobile", util.ValidateMobile)

		// 注册自定义 json tag 函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

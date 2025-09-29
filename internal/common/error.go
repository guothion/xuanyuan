package common

import (
	"fmt"
	"net/http"
	"time"

	"github.com/guothion/xuanyuan/internal/util"
)

// 自定义一个错误信息，方便统一管理
type HTTPError struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// 我们如果想将 HTTPError 作为 error 类型使用，我们就必须给 HTTPError 定义 Error 方法
func (e *HTTPError) Error() string {
	timeStr := time.Unix(e.Timestamp, 0).Format(util.SimpleTimeFormat)
	return fmt.Sprintf("%d [%d] - %s", timeStr, e.Code, e.Message)
}

func NewForbiddenError(format string, args ...interface{}) *HTTPError {
	return &HTTPError{
		Code:      http.StatusForbidden,
		Message:   fmt.Sprintf(format, args...),
		Timestamp: util.Timestamp(),
	}
}

func NewBadRequestError(format string, args ...interface{}) *HTTPError {
	return &HTTPError{
		Code:      http.StatusBadRequest,
		Message:   fmt.Sprintf(format, args...),
		Timestamp: util.Timestamp(),
	}
}

func NewNotFoundError(format string, args ...interface{}) *HTTPError {
	return &HTTPError{
		Code:      http.StatusNotFound,
		Message:   fmt.Sprintf(format, args...),
		Timestamp: util.Timestamp(),
	}
}

func NewInternalServerError(format string, args ...interface{}) *HTTPError {
	return &HTTPError{
		Code:      http.StatusInternalServerError,
		Message:   fmt.Sprintf(format, args...),
		Timestamp: util.Timestamp(),
	}
}

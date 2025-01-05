package codex

import (
	"context"
	"net/http"

	"Ai-HireSphere/common/codex/types"
)

// 成功返回的格式
type OKResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 自定义api错误返回格式
func ErrHandler(err error) (int, any) {
	code := CodeFromError(err)

	return http.StatusOK, types.Status{
		Code:    int32(code.Code()),
		Message: code.Message(),
	}
}

// 自定义api成功返回格式
func OKHandler(_ context.Context, value any) any {
	return OKResponse{
		Code:    OK.Code(),
		Message: OK.Message(),
		Data:    value,
	}
}

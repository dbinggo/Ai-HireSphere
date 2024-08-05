package xcode

import (
	"context"
	"net/http"

	"Ai-HireSphere/common/xcode/types"
)

type OKResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrHandler(err error) (int, any) {
	code := CodeFromError(err)

	return http.StatusOK, types.Status{
		Code:    int32(code.Code()),
		Message: code.Message(),
	}
}

func OKHandler(_ context.Context, value any) any {
	return OKResponse{
		Code:    OK.Code(),
		Message: OK.Message(),
		Data:    value,
	}
}

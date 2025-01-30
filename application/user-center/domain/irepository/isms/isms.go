package isms

import "context"

// todo 需要实现短信接口
type ISms interface {
	SendCaptcha(ctx context.Context, target, code string) error
}

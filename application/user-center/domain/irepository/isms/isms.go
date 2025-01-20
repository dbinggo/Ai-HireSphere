package isms

import "github.com/dbinggo/gerr"

// todo 需要实现短信接口
type ISms interface {
	Send(target string, content string) gerr.Error
}

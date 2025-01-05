package codex

import (
	"strconv"
)

// 相当于一个介于error和proto流之间的自定义中间人
type XCode interface {
	//实现Error方法，就是实现error接口
	Error() string

	Code() int

	Message() string

	//这个是留给从proto的数据（也就是rpc返回给api的protobuf数据）中提取自定义状态码出来用的
	Details() []interface{}
}

// 这里是比较常见的自定义状态码的操作，一个code对应一个message
type Code struct {
	code int
	msg  string
}

func (c Code) Error() string {
	if len(c.msg) > 0 {
		return c.msg
	}

	return strconv.Itoa(c.code)
}

func (c Code) Code() int {
	return c.code
}

func (c Code) Message() string {
	return c.Error()
}

func (c Code) Details() []interface{} {
	return nil
}

func String(s string) Code {
	if len(s) == 0 {
		return OK
	}
	code, err := strconv.Atoi(s)
	if err != nil {
		return ServerErr
	}

	return Code{code: code}
}

// 提供给外部自定义Code，比如用户模块想自定义一些状态码、购物车模块也想自定义一些状态码，这样子可以互不干扰
func New(code int, msg string) Code {
	return Code{code: code, msg: msg}
}

// 用于通用Code的添加
func add(code int, msg string) Code {
	return Code{code: code, msg: msg}
}

package utils

import (
	"Ai-HireSphere/common/codex"
	"regexp"
)

var (
	// 定义错误类型

	ErrEmailIsNotValid = codex.New(10001, "邮箱格式错误")
	ErrPhoneIsNotValid = codex.New(10002, "手机号格式错误")
)

// CheckEmail
//
//	@Description: 检查邮箱格式
//	@param email 邮箱
//	@return error 错误
func CheckEmail(email string) error {
	// 定义邮箱的正则表达式模式
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	// 编译正则表达式
	re := regexp.MustCompile(emailRegex)
	// 使用正则表达式匹配邮箱
	ok := re.MatchString(email)

	if !ok {
		return ErrEmailIsNotValid
	}
	return nil
}

// CheckPhone
//
//	@Description: 校验手机号正确性
//	@param phone
//	@return error
func CheckPhone(phone string) error {
	if len(phone) != 11 {
		return ErrPhoneIsNotValid
	}
	// 定义手机号的正则表达式模式
	phoneRegex := `^1[3-9]\d{9}$`
	// 编译正则表达式
	re := regexp.MustCompile(phoneRegex)
	// 使用正则表达式匹配手机号
	ok := re.MatchString(phone)

	if !ok {
		return ErrPhoneIsNotValid
	}
	return nil
}

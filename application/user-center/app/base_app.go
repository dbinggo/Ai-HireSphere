package app

import (
	"Ai-HireSphere/application/user-center/domain/irepository"
	"Ai-HireSphere/application/user-center/domain/irepository/isms"
	"Ai-HireSphere/application/user-center/domain/services"
	"Ai-HireSphere/common/model/enums"
	"context"
	"github.com/dbinggo/gerr"
)

// 基础功能域

type IBaseApp interface {
	// 验证码接口
	CaptchaSend(ctx context.Context, way enums.CaptchaWayType, method enums.CaptchaMethodType, key string) gerr.Error
	// 验证码校验
	CaptchaCheck(ctx context.Context, way enums.CaptchaWayType, method enums.CaptchaMethodType, key, code string) gerr.Error
}

type BaseApp struct {
	// 这里主要是依赖
	Repo irepository.IRepoBroker
	Sms  isms.ISms
}

func NewBaseApp(repo irepository.IRepoBroker, sms isms.ISms) *BaseApp {
	return &BaseApp{
		Repo: repo,
		Sms:  sms,
	}
}

func (b *BaseApp) CaptchaSend(ctx context.Context, way enums.CaptchaWayType, method enums.CaptchaMethodType, key string) gerr.Error {

	// 初期先只支持邮箱验证，之后的在扩展，只要实现了ISms接口就可以在这里进行依赖注入

	return services.NewBaseCaptcha(ctx, b.Repo, b.Sms).CaptchaSend(way, key)
}
func (b *BaseApp) CaptchaCheck(ctx context.Context, way enums.CaptchaWayType, method enums.CaptchaMethodType, key, code string) gerr.Error {

	// 初期先只支持邮箱验证，之后的在扩展，只要实现了ISms接口就可以在这里进行依赖注入

	return services.NewBaseCaptcha(ctx, b.Repo, b.Sms).CaptchaCheck(way, key, code)
}

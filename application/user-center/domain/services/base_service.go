package services

import (
	ireidsaccess "Ai-HireSphere/application/user-center/domain/irepository/ireids_access"
	"Ai-HireSphere/application/user-center/domain/irepository/isms"
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"Ai-HireSphere/common/model/enums"
	"Ai-HireSphere/common/zlog"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

// fixme 后续要优化错误处理

type IBaseService interface {
	// 验证码发送
	CaptchaSend(way enums.CaptchaWayType, key string) error
	// 验证码校验
	CaptchaCheck(way enums.CaptchaWayType, key, code string) error
}

type BaseService struct {
	ctx context.Context
	rdb ireidsaccess.IRedisAccess
	sms isms.ISms
}

func NewBaseCaptcha(ctx context.Context, rdb ireidsaccess.IRedisAccess, sms isms.ISms) *BaseService {
	return &BaseService{
		ctx: ctx,
		rdb: rdb,
		sms: sms,
	}
}

// CaptchaSend
//
//	@Description: 发送验证码
//	@receiver b
//	@param ctx
//	@param way
//	@param key
//	@return error
func (b *BaseService) CaptchaSend(way enums.CaptchaWayType, key string) error {
	// 先存再发送
	capcha := &entity.Captcha{
		Way: way,
		Key: key,
	}

	// 存
	if err := b.captchaStash(capcha); err != nil {
		return err
	}

	// 发送
	if err := b.sms.Send(key, capcha.CaptchaCode); err != nil {
		return err
	}
	return nil
}

// captchaStash
//
//	@Description:  储存验证码
//	@receiver b
//	@param captcha
//	@return error
func (b *BaseService) captchaStash(captcha *entity.Captcha) error {
	// 整合
	captcha.GenerateCaptcha()
	captcha.GenerateCaptchaCode()
	// 存到redis
	return b.rdb.Set(b.ctx, captcha.RedisKey, captcha.CaptchaCode, time.Minute*5)
}

// CaptchaCheck
//
//	@Description: 验证码检查
//	@receiver b
//	@param ctx
//	@param way
//	@param key
//	@param code
//	@return error
func (b *BaseService) CaptchaCheck(way enums.CaptchaWayType, key, code string) error {
	// 拿出来
	capcha := &entity.Captcha{
		Way: way,
		Key: key,
	}

	ret, err := b.rdb.Get(b.ctx, capcha.RedisKey)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// 验证码过期
			zlog.ErrorfCtx(b.ctx, "验证码过期或未向该手机号发送过验证码")
			return errors.New("验证码过期或未向该手机号发送过验证码")
		}
		return err
	}

	if err = b.rdb.Del(b.ctx, capcha.RedisKey); err != nil {
		return err
	}

	ret, ok := ret.(string)
	if !ok {
		// 安全断言失败
		zlog.ErrorfCtx(b.ctx, "类型断言失败")
		return errors.New("类型断言失败")
	}

	if ret != code {
		// 验证码错误
		zlog.ErrorfCtx(b.ctx, "验证码错误")
		return errors.New("验证码错误")
	}

	return nil
}

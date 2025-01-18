package entity

import "Ai-HireSphere/common/model/enums"

type Captcha struct {
	Way         enums.CaptchaWayType `json:"way"`
	Key         string               `json:"captcha_key"`
	RedisKey    string               `json:"redis_key"`
	CaptchaCode string               `json:"captcha_code"`
}

// 整合RedisKey
func (c *Captcha) GenerateCaptcha() {
	c.RedisKey = enums.RedisKeyPrefix + "Captcha:" + string(c.Way) + ":" + c.Key
}

// 生成验证码 先使用默认的123456
func (c *Captcha) GenerateCaptchaCode() {
	c.CaptchaCode = "123456"
}

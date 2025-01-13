package enums

// 全局rediskey，所有rediskey 应该都带有这个
const RedisKeyPrefix = "ai_hiresphere:"

// 验证码用途
type CaptchaWayType string

const (
	CaptchaWayTypeLogin    CaptchaWayType = "login"    // 登录
	CaptchaWayTypeRegister CaptchaWayType = "register" // 注册
	CaptchaWayTypeReset    CaptchaWayType = "reset"    // 重置
)

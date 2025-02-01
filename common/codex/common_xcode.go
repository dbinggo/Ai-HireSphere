package codex

// 通用Code
var (
	OK                 = add(200, "OK")
	NoLogin            = add(101, "NOT_LOGIN")
	RequestErr         = add(400, "INVALID_ARGUMENT")
	Unauthorized       = add(401, "UNAUTHENTICATED")
	AccessDenied       = add(403, "PERMISSION_DENIED")
	NotFound           = add(404, "NOT_FOUND")
	MethodNotAllowed   = add(405, "METHOD_NOT_ALLOWED")
	Canceled           = add(498, "CANCELED")
	ServerErr          = add(500, "系统繁忙，请稍后重试")
	ServiceUnavailable = add(503, "UNAVAILABLE")
	Deadline           = add(504, "DEADLINE_EXCEEDED")
	LimitExceed        = add(509, "RESOURCE_EXHAUSTED")
)

var (
	_                 = 10 // 用户注册登陆错误 10开头
	UserRegisterExist = add(11001, "用户已存在")

	_                         = 20 // 简历服务20开头
	ResumeUploadEmpty         = add(20001, "简历为空")
	ResumeUploadFail          = add(20002, "简历上传失败")
	ResumeUploadExist         = add(20003, "简历已存在")
	ResumeUploadMAX           = add(20004, "简历过大")
	ResumeDeleteFail          = add(20005, "简历删除失败")
	ResumeNotExist            = add(20006, "简历不存在")
	ResumeDeleteNotPermission = add(20007, "无权限删除简历")
	ResumeDeleteEmpty         = add(20008, "删除简历未找到")
)

package entity

import (
	"Ai-HireSphere/common/coze"
)

// 模型仓库

const (
	InterviewTableName = "7454197247615008820"
)

// 面试Agent
func NewInterviewAgent(api coze.CozeApi) coze.BotApi {
	return coze.NewBotApi(api.GetToken(), InterviewTableName)

}

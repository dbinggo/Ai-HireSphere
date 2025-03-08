package app

import (
	"Ai-HireSphere/application/interview/domain/irepository/idataaccess"
	"Ai-HireSphere/common/coze"
	"context"
)

type IInterviewApp interface {
	CreateInterview(ctx context.Context, userID int64, resumeUrl string, hc string, level int, num int64)
}

type InterviewApp struct {
	CozeApi coze.CozeApi
	Repo    idataaccess.IDataAccess
}

func NewInterviewApp(cozeApi *coze.CozeApi, repo idataaccess.IDataAccess) *InterviewApp {
	return &InterviewApp{
		Repo: repo,
	}
}

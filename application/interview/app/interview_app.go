package app

import (
	"Ai-HireSphere/common/coze"
)

type IInterviewApp interface {
}

type InterviewApp struct {
	CozeApi *coze.CozeApi
}

func NewInterviewApp(cozeApi *coze.CozeApi) *InterviewApp {
	return &InterviewApp{
		CozeApi: cozeApi,
	}
}

func (i *InterviewApp) NewInterview() {

}

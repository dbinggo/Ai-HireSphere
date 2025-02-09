package entity

import (
	"Ai-HireSphere/application/interview-center/repository"
	"Ai-HireSphere/common/coze"
	"fmt"
	"go.uber.org/zap"
)

// 简历逻辑体
type Entity struct {
	repo    repository.Repo
	logger  *zap.Logger
	cozeApi *coze.CozeApi
}

func (entity *Entity) UploadResume(resume repository.Interview) error {
	err := entity.repo.Create(resume)
	if err != nil {
		entity.logger.Error("upload resume error", zap.Error(err))
		return err
	}
	return nil
}

func (entity *Entity) Chat(sessionID int64, message string, isNew bool) (chan string, error) {
	//todo 查找会话表确认是否有会话,没有就创新的
	if isNew {
		sessionID = entity.cozeApi.Bot.CreateSession()
		fmt.Println("sessionID", sessionID)
	}
	ch, err := entity.cozeApi.Bot.Chat(sessionID, message)
	if err != nil {
		fmt.Printf("Chat error: %v", err)
		return nil, err
	}
	return ch, err
}

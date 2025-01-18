package entity

import (
	"Ai-HireSphere/application/interview-center/repository"
	"go.uber.org/zap"
)

// 简历逻辑体
type Entity struct {
	repo   repository.Repo
	logger *zap.Logger
}

func (entity *Entity) UploadResume(resume repository.Resume) error {
	err := entity.repo.Create(resume)
	if err != nil {
		entity.logger.Error("upload resume error", zap.Error(err))
		return err
	}
	return nil
}

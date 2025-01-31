package app

import (
	"Ai-HireSphere/application/interview/domain/irepository"
	"Ai-HireSphere/application/interview/domain/irepository/ioss"
	"Ai-HireSphere/application/interview/domain/model/entity"
	"Ai-HireSphere/application/interview/domain/service"
	"context"
	"github.com/dbinggo/gerr"
	"mime/multipart"
)

type IResumeApp interface {
	UploadResume(ctx context.Context, file multipart.File, handler *multipart.FileHeader) (string, gerr.Error)
	ListResume(ctx context.Context, userId int64, page int, pageSize int) (int64, []entity.ResumeEntity, gerr.Error)
	DeleteResume(ctx context.Context, id int64) gerr.Error
}

type ResumeApp struct {
	Oss  ioss.Ioss
	Repo irepository.IRepoBroker
}

func NewResumeApp(oss ioss.Ioss, repo irepository.IRepoBroker) *ResumeApp {
	return &ResumeApp{
		Oss:  oss,
		Repo: repo,
	}
}

func (r *ResumeApp) UploadResume(ctx context.Context, file multipart.File, handler *multipart.FileHeader) (string, gerr.Error) {
	// 调用服务
	resume, err := service.NewResumeService(ctx, r.Oss, r.Repo).UploadResume(file, handler)
	if err != nil {
		return "", err
	}
	return resume.Url, nil
}

func (r *ResumeApp) ListResume(ctx context.Context, userId int64, page int, pageSize int) (int64, []entity.ResumeEntity, gerr.Error) {
	count, resp, err := r.Repo.ListResume(ctx, userId, page, pageSize)
	if err != nil {
		return 0, nil, err
	}
	return count, resp, nil
}

func (r *ResumeApp) DeleteResume(ctx context.Context, id int64) gerr.Error {
	err := service.NewResumeService(ctx, r.Oss, r.Repo).DeleteResume(id)
	return err
}

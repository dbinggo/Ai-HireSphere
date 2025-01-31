package service

import (
	"Ai-HireSphere/application/interview/domain/irepository/idataaccess"
	"Ai-HireSphere/application/interview/domain/irepository/ioss"
	"Ai-HireSphere/application/interview/domain/model/entity"
	"Ai-HireSphere/common/utils"
	"context"
	"github.com/dbinggo/gerr"
	"mime/multipart"
	"time"
)

type IResumeService interface {
	UploadResume(file multipart.File, handler *multipart.FileHeader) (*entity.ResumeEntity, gerr.Error)
	ListResume(userId int64, page, pageSize int64) ([]entity.ResumeEntity, gerr.Error)
}

type ResumeService struct {
	oss        ioss.Ioss
	ctx        context.Context
	resumeRepo idataaccess.IResumeAccess
}

func NewResumeService(ctx context.Context, oss ioss.Ioss, repo idataaccess.IResumeAccess) *ResumeService {
	return &ResumeService{
		oss:        oss,
		ctx:        ctx,
		resumeRepo: repo,
	}
}

func (r *ResumeService) UploadResume(file multipart.File, handler *multipart.FileHeader) (*entity.ResumeEntity, gerr.Error) {
	// 新建实体上传简历
	resume := &entity.ResumeEntity{
		File:       file,
		FileName:   handler.Filename,
		Handler:    handler,
		UploadTime: time.Now(),
		UserId:     utils.GetUserId(r.ctx),
		Size:       handler.Size,
	}

	// 首先上传文件
	if err := resume.UploadResume(r.oss); err != nil {
		return nil, err
	}

	// 上传文件成功后就落库
	if err := r.resumeRepo.SaveResume(r.ctx, resume); err != nil {
		return nil, err
	}
	return resume, nil

}

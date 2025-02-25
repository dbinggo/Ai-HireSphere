package service

import (
	"Ai-HireSphere/application/interview/domain/irepository/idataaccess"
	"Ai-HireSphere/application/interview/domain/irepository/ioss"
	"Ai-HireSphere/application/interview/domain/model/entity"
	"Ai-HireSphere/common/codex"
	"Ai-HireSphere/common/utils"
	"context"
	"github.com/dbinggo/gerr"
	"mime/multipart"
	"time"
)

type IResumeService interface {
	UploadResume(file multipart.File, handler *multipart.FileHeader, folderId int64) (*entity.ResumeEntity, gerr.Error)
	DeleteResume(id int64) gerr.Error
}

type ResumeService struct {
	oss        ioss.Ioss
	ctx        context.Context
	resumeRepo idataaccess.IResumeAccess
}

func NewResumeService(ctx context.Context, oss ioss.Ioss, repo idataaccess.IResumeAccess) IResumeService {
	return &ResumeService{
		oss:        oss,
		ctx:        ctx,
		resumeRepo: repo,
	}
}

func (r *ResumeService) UploadResume(file multipart.File, handler *multipart.FileHeader, folderId int64) (*entity.ResumeEntity, gerr.Error) {
	// 新建实体上传简历
	resume := &entity.ResumeEntity{
		File:       file,
		FileName:   handler.Filename,
		Handler:    handler,
		FolderId:   folderId,
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

func (r *ResumeService) DeleteResume(id int64) gerr.Error {
	// 先查到这个简历
	resume, err := r.resumeRepo.FindResumeById(r.ctx, id)
	if err != nil {
		return err
	}
	if resume.UserId != utils.GetUserId(r.ctx) {
		return gerr.WithStack(codex.ResumeDeleteNotPermission)
	}
	// 删除oss文件
	if err = resume.DeleteResume(r.oss); err != nil {
		return err
	}
	// 删除数据库信息
	return r.resumeRepo.DeleteResume(r.ctx, id)
}

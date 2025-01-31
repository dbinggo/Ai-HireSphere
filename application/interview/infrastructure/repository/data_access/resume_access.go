package data_access

import (
	"Ai-HireSphere/application/interview/domain/model/entity"
	"Ai-HireSphere/common/codex"
	"Ai-HireSphere/common/model"
	"context"
	"errors"
	"github.com/dbinggo/gerr"
	"gorm.io/gorm"
)

// SaveResume
//
//	@Description: 保存简历
//	@receiver g
//	@param ctx
//	@param resume
//	@return gerr.Error
func (g *GormOpts) SaveResume(ctx context.Context, resume *entity.ResumeEntity) gerr.Error {
	resumeModel, err := resume.Transform().Save(ctx, g.db, resume.Transform())
	if err != nil {
		return gerr.Wraps(codex.ResumeUploadFail, err)
	}
	resume.From(resumeModel)

	return nil
}

// ListResume
//
//	@Description: 列出简历
//	@receiver g
//	@param ctx
//	@param userId
//	@param page
//	@param pageSize
//	@return []entity.ResumeEntity
//	@return gerr.Error
func (g *GormOpts) ListResume(ctx context.Context, userId int64, page, pageSize int) (int64, []entity.ResumeEntity, gerr.Error) {
	count, resumeModels, err := (&model.CommonAdapter[model.TResume]{}).List(ctx, g.db, pageSize, page-1, "user_id = ?", userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, []entity.ResumeEntity{}, nil
		}
		return 0, nil, gerr.Wraps(codex.ResumeUploadFail, err)
	}
	resumeEntitys := make([]entity.ResumeEntity, 0)
	for _, resumeModel := range resumeModels {
		resume := entity.ResumeEntity{}
		resume.From(&resumeModel)
		resumeEntitys = append(resumeEntitys, resume)
	}
	return count, resumeEntitys, nil
}

// DeleteResume
//
//	@Description: 删除指定简历
//	@receiver g
//	@param ctx
//	@param id
//	@return gerr.Error
func (g *GormOpts) DeleteResume(ctx context.Context, id int64) gerr.Error {
	err := (&model.CommonAdapter[model.TResume]{}).Delete(ctx, g.db, id)
	if err != nil {
		return gerr.Wraps(codex.ResumeDeleteFail, err)
	}
	return nil
}

// FindResumeById
//
//	@Description: 查找指定简历
//	@receiver g
//	@param ctx
//	@param id
//	@return entity.ResumeEntity
//	@return gerr.Error
func (g *GormOpts) FindResumeById(ctx context.Context, id int64) (entity.ResumeEntity, gerr.Error) {
	resumeModel, err := (&model.CommonAdapter[model.TResume]{}).GetOne(ctx, g.db, "id = ?", id)
	if err != nil {
		return entity.ResumeEntity{}, gerr.Wraps(codex.ResumeNotExist, err)
	}
	resume := new(entity.ResumeEntity)
	resume.From(resumeModel)

	return *resume, nil
}

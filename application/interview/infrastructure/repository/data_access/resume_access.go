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
func (g *GormOpts) ListResume(ctx context.Context, userId int64, page, pageSize int, folderId int64) (int64, []entity.ResumeEntity, gerr.Error) {
	count, resumeModels, err := (&model.CommonAdapter[model.TResume]{}).List(ctx, g.db, pageSize, (page-1)*pageSize, "user_id = ? and folder_id = ? and folder_id > -1", userId, folderId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, []entity.ResumeEntity{}, nil
		}
		return 0, nil, gerr.Wraps(codex.ResumeUploadFail, err)
	}
	resumeEntitys := make([]entity.ResumeEntity, 0)
	for _, resumeModel := range resumeModels {
		resume := entity.ResumeEntity{}
		resume.From(resumeModel)
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
	resume.From(*resumeModel)
	return *resume, nil
}

// FindResumeByFolderId
//
//	@Description: 根据文件夹id查找简历
//	@receiver g
//	@param ctx
//	@param folderId
//	@return []entity.ResumeEntity
//	@return gerr.Error
func (g *GormOpts) FindResumeByFolderId(ctx context.Context, folderId int64) ([]entity.ResumeEntity, gerr.Error) {
	_, resumes, err := (model.TResume{}).List(ctx, g.db, -1, -1, "folder_id = ?", folderId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gerr.Wraps(codex.ResumeFindFail, err)
	}
	var resumeEntitys []entity.ResumeEntity
	for _, resume := range resumes {
		resumeEntity := entity.ResumeEntity{}
		resumeEntity.From(resume)
		resumeEntitys = append(resumeEntitys, resumeEntity)
	}
	return resumeEntitys, nil
}

// CreateFolder
//
//	@Description: 创建简历文件夹
//	@receiver g
//	@param ctx
//	@param userId
//	@param folderName
//	@return gerr.Error
func (g *GormOpts) CreateFolder(ctx context.Context, userId int64, folderName string) (model.TFolder, gerr.Error) {

	folder := model.TFolder{
		Name:   folderName,
		UserId: userId,
	}
	folder, err := folder.Save(ctx, g.db, folder)
	if err != nil {
		return model.TFolder{}, gerr.Wraps(codex.FolderCreateFail, err)
	}
	return folder, nil
}

// ListFolder
//
//	@Description: 列出简历文件夹
//	@receiver g
//	@param ctx
//	@param userId
//	@param folderName
//	@return gerr.Error
func (g *GormOpts) ListFolder(ctx context.Context, userId int64) ([]model.TFolder, gerr.Error) {
	_, folders, err := (&model.TFolder{}).List(ctx, g.db, -1, -1, "user_id = ?", userId)
	if err != nil {
		return nil, gerr.Wraps(codex.FolderListFail, err)
	}
	return folders, nil
}

// DeleteFolder
//
//	@Description: 删除简历文件夹
//	@receiver g
//	@param ctx
//	@param userId
//	@param folderName
//	@return []model.TFolder
//	@return gerr.Error
func (g *GormOpts) DeleteFolder(ctx context.Context, id int64) gerr.Error {
	// 要删除简历文件夹就要直接删除所有简历

	list, err := g.FindResumeByFolderId(ctx, id)
	if err != nil {
		return err
	}
	// 删除
	for _, resume := range list {
		err := g.DeleteResume(ctx, resume.Id)
		if err != nil {
			return err
		}
	}

	err1 := (&model.TFolder{}).Delete(ctx, g.db, id)
	if err1 != nil {
		return gerr.Wraps(codex.FolderDeleteFail, err1)
	}
	return nil
}

// UpdateFolder
//
//	@Description: 更新简历文件夹
//	@receiver g
//	@param ctx
//	@param userId
//	@param folderName
//	@return model.TFolder
//	@return gerr.Error
func (g *GormOpts) UpdateFolder(ctx context.Context, id int64, folderName string) (model.TFolder, gerr.Error) {
	tf := new(model.TFolder)
	folder, err := tf.UpdateOne(ctx, g.db, model.TFolder{
		CommonModel: model.CommonModel{
			ID: id,
		},
		Name: folderName,
	})
	if err != nil {
		return model.TFolder{}, gerr.Wraps(codex.FolderUpdateFail, err)
	}
	return folder, nil
}

// FindFolderById
//
//	@Description: 根据id查找简历文件夹
//	@receiver g
//	@param ctx
//	@param id
//	@return model.TFolder
//	@return gerr.Error
func (g *GormOpts) FindFolderById(ctx context.Context, id int64) (model.TFolder, gerr.Error) {
	folder, err := (&model.TFolder{}).GetOne(ctx, g.db, "id = ?", id)
	if err != nil {
		return model.TFolder{}, gerr.Wraps(codex.FolderFindFail, err)
	}
	return *folder, nil
}

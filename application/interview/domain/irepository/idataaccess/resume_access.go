package idataaccess

import (
	"Ai-HireSphere/application/interview/domain/model/entity"
	"Ai-HireSphere/common/model"
	"context"
	"github.com/dbinggo/gerr"
)

type IResumeAccess interface {
	SaveResume(ctx context.Context, entity *entity.ResumeEntity) gerr.Error
	ListResume(ctx context.Context, userId int64, page, pageSize int) (int64, []entity.ResumeEntity, gerr.Error)
	DeleteResume(ctx context.Context, id int64) gerr.Error
	FindResumeById(ctx context.Context, id int64) (entity.ResumeEntity, gerr.Error)
	FindResumeByFolderId(ctx context.Context, folderId int64) ([]entity.ResumeEntity, gerr.Error)
}

type IFolderAccess interface {
	FindFolderById(ctx context.Context, id int64) (model.TFolder, gerr.Error)
	UpdateFolder(ctx context.Context, id int64, folderName string) (model.TFolder, gerr.Error)
	DeleteFolder(ctx context.Context, id int64) gerr.Error
	ListFolder(ctx context.Context, userId int64) ([]model.TFolder, gerr.Error)
	CreateFolder(ctx context.Context, userId int64, folderName string) (model.TFolder, gerr.Error)
}

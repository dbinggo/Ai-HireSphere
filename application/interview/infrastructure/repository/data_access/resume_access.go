package data_access

import (
	"Ai-HireSphere/application/interview/domain/model/entity"
	"Ai-HireSphere/common/codex"
	"context"
	"github.com/dbinggo/gerr"
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

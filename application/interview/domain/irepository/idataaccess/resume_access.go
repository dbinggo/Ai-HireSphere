package idataaccess

import (
	"Ai-HireSphere/application/interview/domain/model/entity"
	"context"
	"github.com/dbinggo/gerr"
)

type IResumeAccess interface {
	SaveResume(ctx context.Context, entity *entity.ResumeEntity) gerr.Error
}

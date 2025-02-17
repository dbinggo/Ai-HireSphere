package idataaccess

import (
	"Ai-HireSphere/application/interview/domain/model/entity"
	"context"
	"github.com/dbinggo/gerr"
)

type IChatAccess interface {
	CreateChatSession(ctx context.Context, chat *entity.ChatEntity) (id int64, error gerr.Error)
	GetOneChat(ctx context.Context, id int64) (entity.ChatEntity, gerr.Error)
	ListChatSession(ctx context.Context, userID int64, page, pageSize int) (int64, []entity.ChatEntity, gerr.Error)
	UpdateChat(ctx context.Context, chat *entity.ChatEntity) gerr.Error
}

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

const (
	DefaultChat = "未命名"
)

func (g *GormOpts) CreateChatSession(ctx context.Context, chat *entity.ChatEntity) (id int64, error gerr.Error) {
	if chat.SessionTitle == "" {
		chat.SessionTitle = DefaultChat
	}
	save, err := chat.Transform().Save(ctx, g.db, chat.Transform())
	if err != nil {
		return 0, gerr.Wraps(codex.ChatSessionCreateFail, err)
	}
	return save.ID, nil
}

func (g *GormOpts) GetOneChat(ctx context.Context, id int64) (entity.ChatEntity, gerr.Error) {
	chatModel, err := (&model.CommonAdapter[model.TChat]{}).GetOne(ctx, g.db, "id = ?", id)
	if err != nil {
		return entity.ChatEntity{}, gerr.Wraps(codex.ChatSessionNotExist, err)
	}
	chat := new(entity.ChatEntity)
	chat.From(*chatModel)

	return *chat, nil
}

func (g *GormOpts) UpdateChat(ctx context.Context, chat *entity.ChatEntity) gerr.Error {
	_, err := (&model.CommonAdapter[model.TChat]{}).UpdateOne(ctx, g.db, chat.Transform())
	if err != nil {
		return gerr.Wraps(codex.ChatSessionUpdateFail, err)
	}
	return nil
}

func (g *GormOpts) ListChatSession(ctx context.Context, userID int64, page, pageSize int) (int64, []entity.ChatEntity, gerr.Error) {
	count, chats, err := (&model.CommonAdapter[model.TChat]{}).List(ctx, g.db, pageSize, (page-1)*pageSize, "user_id = ?", userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, []entity.ChatEntity{}, nil
		}
		return 0, nil, gerr.Wraps(codex.ChatSessionNotExist, err)
	}

	chatEntitys := make([]entity.ChatEntity, 0)
	for _, chatModel := range chats {
		chat := entity.ChatEntity{}
		chat.From(chatModel)
		chatEntitys = append(chatEntitys, chat)
	}
	return count, chatEntitys, nil
}

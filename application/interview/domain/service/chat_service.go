package service

import (
	"Ai-HireSphere/application/interview/domain/irepository/idataaccess"
	"Ai-HireSphere/application/interview/domain/model/entity"
	"Ai-HireSphere/application/interview/infrastructure/repository/data_access"
	"Ai-HireSphere/common/codex"
	"Ai-HireSphere/common/coze"
	"context"
	"fmt"
	"github.com/dbinggo/gerr"
)

type IChatService interface {
	CreateSession(userID int64) (int64, gerr.Error)
	Chat(id int64, message string) (chan coze.BotStreamReply, gerr.Error)
	UpdateChatName(id int64, name string) gerr.Error
	ListChats(userID int64, pageSize, pageNum int) (int64, []entity.ChatEntity, gerr.Error)
	GetChatHistory(id int64) ([]coze.BotMessage, gerr.Error)
}

type ChatService struct {
	ctx  context.Context
	repo idataaccess.IChatAccess
	api  coze.CozeApi
}

func (c *ChatService) CreateSession(userID int64) (int64, gerr.Error) {
	session := c.api.Bot.CreateSession()
	var data entity.ChatEntity
	data.SessionID = session
	data.UserID = userID
	data.SessionTitle = data_access.DefaultChat
	return c.repo.CreateChatSession(c.ctx, &data)
}

func (c *ChatService) Chat(id int64, message string) (chan coze.BotStreamReply, gerr.Error) {
	fmt.Printf("my_session id: %d", id)
	chat, g := c.repo.GetOneChat(c.ctx, id)
	if g != nil {
		return nil, g
	}
	fmt.Printf("session id: %d", chat.SessionID)
	ch, err := c.api.Bot.Chat(chat.SessionID, message)
	if err != nil {
		return nil, gerr.Wraps(codex.ChatSessionNotExist, err)
	}
	return ch, nil
}

func (c *ChatService) UpdateChatName(id int64, name string) gerr.Error {
	return c.repo.UpdateChat(c.ctx, &entity.ChatEntity{
		ID:           id,
		SessionTitle: name,
	})
}

func (c *ChatService) ListChats(userID int64, pageSize, pageNum int) (int64, []entity.ChatEntity, gerr.Error) {
	return c.repo.ListChatSession(c.ctx, userID, pageNum, pageSize)
}

func (c *ChatService) GetChatHistory(id int64) ([]coze.BotMessage, gerr.Error) {
	chat, g := c.repo.GetOneChat(c.ctx, id)
	if g != nil {
		return nil, g
	}
	history, err := c.api.Bot.GetChatHistory(chat.SessionID)
	if err != nil {
		return nil, gerr.Wraps(codex.ChatSessionNotExist, err)
	}
	return history, nil
}

func NewChatService(ctx context.Context, repo idataaccess.IChatAccess, api coze.CozeApi) *ChatService {
	return &ChatService{
		ctx:  ctx,
		repo: repo,
		api:  api,
	}
}

var _ IChatService = &ChatService{}

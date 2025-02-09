package interview_center

import (
	"context"

	"Ai-HireSphere/application/interview-center/protocol/api/internal/svc"
	"Ai-HireSphere/application/interview-center/protocol/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.ChatRequest) (chan string, error) {
	chat, err := l.svcCtx.Entity.Chat(req.SessionID, req.Message, req.IsNew)
	if err != nil {
		l.Errorf("Chat error: %v", err)
		return nil, err
	}
	return chat, nil
}

package interview

import (
	"Ai-HireSphere/common/coze"
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatAgentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatAgentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatAgentLogic {
	return &ChatAgentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatAgentLogic) ChatAgent(req *types.ChatAgentReq) (stream chan coze.BotStreamReply, err error) {
	sessionID := req.SessionID
	if req.IsNew {
		sessionID, err = l.svcCtx.ResumeAPP.CreateSession(l.ctx, 123)
		if err != nil {
			return nil, err
		}
	}

	return l.svcCtx.ResumeAPP.Chat(l.ctx, sessionID, req.Message)
}

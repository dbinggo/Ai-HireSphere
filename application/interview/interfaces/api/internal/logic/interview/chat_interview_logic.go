package interview

import (
	"Ai-HireSphere/common/coze"
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatInterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatInterviewLogic {
	return &ChatInterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatInterviewLogic) ChatInterview(req *types.ChatInterviewReq) (stream chan coze.BotStreamReply, err error) {
	stream, g := l.svcCtx.ResumeAPP.Chat(l.ctx, req.SessionID, req.Answer)
	if g != nil {
		return nil, g
	}
	return stream, nil
}

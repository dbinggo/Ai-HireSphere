package interview

import (
	"Ai-HireSphere/application/interview/domain/model/entity"
	"Ai-HireSphere/common/coze"
	"context"
	"fmt"

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
	var message string
	if req.Hc != "" && req.QuestionNum > 0 && req.Level > 0 && req.PdfUrl != "" {
		var level string
		var ok bool
		if level, ok = entity.Level[req.Level]; !ok {
			return nil, fmt.Errorf("level error")
		}
		message = fmt.Sprintf("面试岗位:%s, 面试难度:%s, 面试题个数:%d, 简历链接:%s", req.Hc, level, req.QuestionNum, req.PdfUrl)
	}

	if req.Answer != "" {
		message = req.Answer
	}

	if message == "" {
		return nil, fmt.Errorf("message is empty")
	}

	chat, g := l.svcCtx.ResumeAPP.Chat(l.ctx, sessionID, message)
	if g != nil {
		return nil, g
	}
	return chat, nil
}

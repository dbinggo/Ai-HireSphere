package interview

import (
	"Ai-HireSphere/common/utils"
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatNewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatNewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatNewLogic {
	return &ChatNewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatNewLogic) ChatNew() (resp *types.ChatNewResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.ChatNewResp)
	resp.SessionID, err = l.svcCtx.ResumeAPP.CreateSession(l.ctx, utils.GetUserId(l.ctx))
	if err != nil {
		return nil, err
	}
	return
}

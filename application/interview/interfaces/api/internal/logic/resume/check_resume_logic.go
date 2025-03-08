package resume

import (
	"Ai-HireSphere/common/coze"
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckResumeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckResumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckResumeLogic {
	return &CheckResumeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckResumeLogic) CheckResume(req *types.CheckResumeReq) (chan coze.WorkFlowStreamResp, error) {
	resume, g := l.svcCtx.ResumeAPP.CheckResume(l.ctx, req.Condition, req.NeedNum, req.FolderID)
	if g != nil {
		l.Logger.Error(g)
		return nil, g
	}

	return resume, nil
}

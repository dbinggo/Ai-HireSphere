package resume

import (
	"Ai-HireSphere/common/coze"
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EvaluateResumeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEvaluateResumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EvaluateResumeLogic {
	return &EvaluateResumeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EvaluateResumeLogic) EvaluateResume(req *types.EvaluateResumeReq) (chan coze.WorkFlowStreamResp, error) {
	evaluate, g := l.svcCtx.ResumeAPP.EvaluateResume(l.ctx, req.ResumeUrl, req.Content, req.Jd)
	if g != nil {
		l.Logger.Error(g)
		return nil, g
	}

	return evaluate, g
}

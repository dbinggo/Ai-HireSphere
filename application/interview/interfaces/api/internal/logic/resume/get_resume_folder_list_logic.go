package resume

import (
	"Ai-HireSphere/common/utils"
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetResumeFolderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetResumeFolderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetResumeFolderListLogic {
	return &GetResumeFolderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetResumeFolderListLogic) GetResumeFolderList(req *types.GetResumeFolderListReq) (resp *types.GetResumeFolderListResp, err error) {
	total, res, err := l.svcCtx.ResumeAPP.ListResume(l.ctx, utils.GetUserId(l.ctx), req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	for _, v := range res {
		temp := types.FolderInfo{
			FolderId:   v.FolderId,
			FolderName: v.FileName,
		}
		resp.List = append(resp.List, temp)
		resp.Total = total
	}
	return resp, nil
}

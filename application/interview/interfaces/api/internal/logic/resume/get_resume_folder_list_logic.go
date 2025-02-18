package resume

import (
	"Ai-HireSphere/common/zlog"
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

func (l *GetResumeFolderListLogic) GetResumeFolderList() (resp *types.GetResumeFolderListResp, err error) {
	zlog.InfofCtx(l.ctx, "%s", l.svcCtx.ResumeAPP)
	res, err := l.svcCtx.ResumeAPP.ListFolder(l.ctx)
	if err != nil {
		return nil, err
	}
	resp = &types.GetResumeFolderListResp{
		CommonListResp: types.CommonListResp{
			Total: int64(len(res)),
		},
		List: make([]types.FolderInfo, 0),
	}
	for _, v := range res {
		resp.List = append(resp.List, types.FolderInfo{
			FolderId:   v.ID,
			FolderName: v.Name,
		})
	}
	return resp, nil
}

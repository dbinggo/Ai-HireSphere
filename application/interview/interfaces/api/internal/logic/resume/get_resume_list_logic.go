package resume

import (
	"Ai-HireSphere/common/utils"
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetResumeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetResumeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetResumeListLogic {
	return &GetResumeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetResumeListLogic) GetResumeList(req *types.GetResumeListReq) (resp *types.GetResumeListResp, err error) {
	count, info, err := l.svcCtx.ResumeAPP.ListResume(l.ctx, utils.GetUserId(l.ctx), req.Page, req.PageSize, req.FolderID)
	if err != nil {
		return nil, err
	}
	retData := make([]types.ResumeInfo, 0)
	for _, v := range info {
		temp := types.ResumeInfo{
			ResumeId:   v.Id,
			ResumeName: v.FileName,
			ResumeSize: v.Size,
			UploadTime: v.UploadTime.Format("2006-01-02 15:04:05"),
			ResumeUrl:  v.Url,
			FolderId:   v.FolderId,
		}
		retData = append(retData, temp)
	}

	return &types.GetResumeListResp{
		CommonListResp: types.CommonListResp{
			Total: count,
		},
		List: retData,
	}, err
}

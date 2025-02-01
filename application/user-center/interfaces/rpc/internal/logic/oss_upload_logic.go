package logic

import (
	"context"

	"Ai-HireSphere/application/user-center/interfaces/rpc/internal/svc"
	"Ai-HireSphere/common/call/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type OssUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOssUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OssUploadLogic {
	return &OssUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OssUploadLogic) OssUpload(in *user.OSSUploadReq) (*user.OSSUploadResp, error) {
	// todo: add your logic here and delete this line

	return &user.OSSUploadResp{}, nil
}

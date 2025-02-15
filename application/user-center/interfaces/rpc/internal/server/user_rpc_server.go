// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package server

import (
	"context"

	"Ai-HireSphere/application/user-center/interfaces/rpc/internal/logic"
	"Ai-HireSphere/application/user-center/interfaces/rpc/internal/svc"
	"Ai-HireSphere/common/call/user"
)

type UserRpcServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserRpcServer
}

func NewUserRpcServer(svcCtx *svc.ServiceContext) *UserRpcServer {
	return &UserRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *UserRpcServer) FindUserById(ctx context.Context, in *user.Id) (*user.UserInfo, error) {
	l := logic.NewFindUserByIdLogic(ctx, s.svcCtx)
	return l.FindUserById(in)
}

func (s *UserRpcServer) FindUserByPhone(ctx context.Context, in *user.Phone) (*user.UserInfo, error) {
	l := logic.NewFindUserByPhoneLogic(ctx, s.svcCtx)
	return l.FindUserByPhone(in)
}

func (s *UserRpcServer) OssUpload(ctx context.Context, in *user.OSSUploadReq) (*user.OSSUploadResp, error) {
	l := logic.NewOssUploadLogic(ctx, s.svcCtx)
	return l.OssUpload(in)
}

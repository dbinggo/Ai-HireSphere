package user

import (
	"Ai-HireSphere/common/call/user"
	"Ai-HireSphere/common/model/enums"
	"context"

	"Ai-HireSphere/application/user-center/interfaces/api/internal/svc"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserApp.RegisterUser(l.ctx, enums.UserRegisterWayType(req.Way), req.Data, req.Code)

	userInfo, err := l.svcCtx.UserRpc.FindUserByPhone(l.ctx, &user.Phone{
		Phone: "13131227873",
	})

	userInfo.Phone = "13131227873"

	return &types.RegisterResp{
		Token: res,
	}, err

}

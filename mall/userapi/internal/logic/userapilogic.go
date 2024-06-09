package logic

import (
	"context"
	"rpc_client/types/user"

	"userapi/internal/svc"
	"userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserapiLogic {
	return &UserapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserapiLogic) Register(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.SaveUser(context.Background(), &user.UserRequest{
		Name:   req.Name,
		Gender: req.Gender,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.Response)
	resp.Message = "注册成功"
	return
}

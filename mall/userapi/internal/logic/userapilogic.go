package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"rpc_client/types/user"
	"time"

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

func (l *UserapiLogic) Login(req *types.LoginRequest) (resp *types.Response, err error) {
	userId := 100
	auth := l.svcCtx.Config.Auth
	data, err := l.getToken(auth.AccessSecret, time.Now().Unix(), auth.AccessExpire, int64(userId))
	if err != nil {
		return nil, err
	}
	resp = new(types.Response)
	resp.Message = "登录成功"
	resp.Data = data
	return
}

func (l *UserapiLogic) getToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *UserapiLogic) GetUser(t *types.IdRequest) (resp *types.Response, err error) {
	userResponse, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: t.Id,
	})

	if err != nil {
		return nil, err
	}

	resp = &types.Response{
		Message: "获取成功",
		Data:    userResponse,
	}
	return
}

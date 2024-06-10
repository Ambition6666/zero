package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"rpc_client/userclient"
	"userapi/internal/config"
	"userapi/internal/handler"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     userclient.User
	MiddlerWare *handler.UserMiddleWare
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		MiddlerWare: handler.NewUserMiddleware(),
	}
}

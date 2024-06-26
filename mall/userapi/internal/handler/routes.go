// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"userapi/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	handler := NewHandler(serverCtx)

	// 全局中间件
	server.Use(serverCtx.MiddlerWare.LoginAndReg)

	server.AddRoutes(
		// 局部中间件
		rest.WithMiddlewares(
			//[]rest.Middleware{
			//	serverCtx.MiddlerWare.LoginAndReg,
			//},
			nil,
			rest.Route{
					Method:  http.MethodPost,
					Path:    "/register",
					Handler: handler.Register(),
			},
			rest.Route{
					Method:  http.MethodPost,
					Path:    "/login",
					Handler: handler.Login(),

			},
		),
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/get/:id",
				Handler: handler.GetUser(),
			},
		}, rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}

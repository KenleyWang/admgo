// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	login "github.com/admgo/admgo/services/auth/api/internal/handler/login"
	logout "github.com/admgo/admgo/services/auth/api/internal/handler/logout"
	"github.com/admgo/admgo/services/auth/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: login.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/auth/v1"),
		rest.WithTimeout(10000*time.Millisecond),
		rest.WithMaxBytes(1048576),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthenticationMiddleware, serverCtx.CSRFTokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/logout",
					Handler: logout.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/auth/v1"),
		rest.WithTimeout(10000*time.Millisecond),
		rest.WithMaxBytes(1048576),
	)
}

// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"github.com/admgo/admgo/services/auth/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: AuthHandler(serverCtx),
			},
		},
	)
}

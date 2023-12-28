package handler

import (
	"net/http"

	"github.com/admgo/admgo/services/cmdb/cmdb/internal/logic"
	"github.com/admgo/admgo/services/cmdb/cmdb/internal/svc"
	"github.com/admgo/admgo/services/cmdb/cmdb/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CmdbHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCmdbLogic(r.Context(), svcCtx)
		resp, err := l.Cmdb(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

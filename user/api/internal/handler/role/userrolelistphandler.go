package role

import (
	"net/http"

	"github.com/admgo/admgo/user/api/internal/logic/role"
	"github.com/admgo/admgo/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRoleListPHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := role.NewUserRoleListPLogic(r.Context(), svcCtx)
		resp, err := l.UserRoleListP()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

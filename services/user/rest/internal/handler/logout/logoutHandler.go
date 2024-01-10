package logout

import (
	"net/http"

	"github.com/admgo/admgo/services/user/rest/internal/logic/logout"
	"github.com/admgo/admgo/services/user/rest/internal/svc"
	"github.com/admgo/admgo/services/user/rest/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogoutReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		sessionID, err := r.Cookie("user_token")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		ok, err := svcCtx.Redis.Del("/" + svcCtx.Config.Name + "/rest/user_token/" + sessionID.Value)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		if ok == 0 {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		l := logout.NewLogoutLogic(r.Context(), svcCtx)

		resp, err := l.Logout(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

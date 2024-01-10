package middleware

import (
	"github.com/admgo/admgo/pkg/errorx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type (
	AuthenticationOptions struct {
		RedisConf redis.RedisConf
	}
	AuthenticationMiddleware struct {
		sessionRedis *redis.Redis
	}
)

func NewAuthenticationMiddleware(c redis.RedisConf) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		sessionRedis: redis.MustNewRedis(c, redis.WithPass(c.Pass)),
	}
}

func (m *AuthenticationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("user_token")
		if err != nil {
			httpx.Error(w, errorx.NoLogin)
			return
		}
		isExist, err := m.sessionRedis.Exists("/user/rest/user_token/" + cookie.Value)
		if err != nil {
			httpx.Error(w, errorx.ServerErrForRedis)
			return
		}
		if !isExist {
			httpx.Error(w, errorx.NoLogin)
			return
		}
		next(w, r)
	}
}

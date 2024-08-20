package Middleware

import (
	"context"
	"fmt"
	"go-rest-api/Constant"
	"go-rest-api/Helper/Response"
	"net/http"
	"strings"
)

type MiddlewareInterface interface {
	UserAuth() func(http.Handler) http.Handler
	BasicAuthSwagger() func(http.Handler) http.Handler
}

func (m Middleware) UserAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if len(tokenString) == 0 {
				Response.ResponseError(w, nil, Constant.StatusUnauthorizedTokenNotExists)
				return
			}

			if !strings.Contains(tokenString, "Bearer ") {
				Response.ResponseError(w, nil, Constant.StatusUnauthorizedBearerNotFound)
				return
			}

			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
			claims, err := m.Jwt.VerifyToken(tokenString)
			if err != nil {
				Response.ResponseError(w, err, Constant.StatusUnauthorizedErrorVerifying)
				return
			}

			ctx := context.WithValue(r.Context(), "claims_value", claims)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func (m Middleware) BasicAuthSwagger() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			var username string
			var password string
			username, password, Ok := r.BasicAuth()
			if !Ok {
				w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic `+username+password))
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
				return
			}

			if m.SwaggerSetting.Username == username && m.SwaggerSetting.Password == password {
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic `+username+password))
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
		})
	}
}

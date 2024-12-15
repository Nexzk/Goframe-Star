package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"star/utility"
)

func Auth(r *ghttp.Request) {
	var (
		jwtKey      = utility.JwtKey
		tokenString = r.GetHeader("Authorization")
	)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}
	r.Middleware.Next()

}

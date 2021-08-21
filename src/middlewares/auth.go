package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Cookies struct {
	Writer http.ResponseWriter
	Token string
	IsLoggedIn bool
}

func (cookies *Cookies) SetToken (token string) {
	http.SetCookie(cookies.Writer, &http.Cookie{
		Name: "authorization",
		Value: token,
		HttpOnly: true,
		Path: "/",
		Expires: time.Now().Add(time.Second * 3600),
	})
}

func setCtxValue(ctx *gin.Context, value interface{}) {
	newCtx := context.WithValue(ctx.Request.Context(), "authorization", value)
	ctx.Request= ctx.Request.WithContext(newCtx)
}

func decodeToken(tokenStr string) {
	JWT_SECRET := os.Getenv("JWT_SECRET")

	token, _ := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	log.Println(token.Claims.(jwt.MapClaims)["id"])
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookies := Cookies{Writer: ctx.Writer}
		tokenString, _ := ctx.Cookie("authorization")

		if tokenString != "" {
			cookies.Token = tokenString
		}
		
		setCtxValue(ctx, &cookies)
		ctx.Next()
	}
}
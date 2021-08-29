package middlewares

import (
	"context"
	"errors"
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
	UserId string
}

func (cookies *Cookies) SetToken (token string, duration time.Time) {
	http.SetCookie(cookies.Writer, &http.Cookie{
		Name: "authorization",
		Value: token,
		HttpOnly: true,
		Path: "/",
		Expires: duration,
	})
}

func setCtxValue(ctx *gin.Context, key string, value interface{}) {
	newCtx := context.WithValue(ctx.Request.Context(), key, value)
	ctx.Request= ctx.Request.WithContext(newCtx)
}

func getUserId(tokenStr string) (string, error) {
	JWT_SECRET := os.Getenv("JWT_SECRET")

	token, tokenError := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if tokenError != nil {
		return "", tokenError
	}

	if (token.Valid) {
		userId := token.Claims.(jwt.MapClaims)["id"].(string)
		return userId, nil
	}
	return "", errors.New("token is not valid")
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookies := Cookies{Writer: ctx.Writer, IsLoggedIn: false, Token: "", UserId: ""}
		tokenString, _ := ctx.Cookie("authorization")

		if tokenString != "" {
			cookies.Token = tokenString
			userId, userError := getUserId(tokenString)

			if userError != nil || userId != "" {
				cookies.IsLoggedIn = true
				cookies.UserId = userId
			}
		}

		setCtxValue(ctx, "authorization", &cookies)
		ctx.Next()
	}
}
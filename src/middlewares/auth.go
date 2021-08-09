package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Cookie("authorization"))
		// c.AbortWithStatus(403)
		c.Next()
	}
}
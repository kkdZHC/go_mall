package middleware

import (
	"go_mall/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = http.StatusNetworkAuthenticationRequired
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = http.StatusRequestTimeout
			}
		}
		if code != 200 {
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    "token_err",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

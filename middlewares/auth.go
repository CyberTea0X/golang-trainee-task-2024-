package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		if token != "admin_token" {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
	}
}

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		// Почти полностью повторяет AdminAuth, но не хочу усложнять логику пока что
		if token != "admin_token" && token != "user_token" {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
	}
}

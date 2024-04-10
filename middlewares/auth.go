package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenType int

const (
	AnyToken TokenType = iota
	AdminToken
)

func TokenAuth(t TokenType) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		if t == AdminToken && token != "admin_token" {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

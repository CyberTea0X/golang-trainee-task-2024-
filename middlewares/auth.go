package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func tokenAuth(validTokens []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		isValid := false
		for _, valid := range validTokens {
			if token == valid {
				isValid = true
				break
			}
		}
		if !isValid {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
	}
}

func AdminAuth() gin.HandlerFunc { return tokenAuth([]string{"admin_token"}) }
func UserAuth() gin.HandlerFunc  { return tokenAuth([]string{"admin_token", "user_token"}) }

package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
			c.Set("logged_in_user",token)
		} else {
			c.Set("is_logged_in", false)
			c.Set("logged_in_user",token)
		}
	}
}

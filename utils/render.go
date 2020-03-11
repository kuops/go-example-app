package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	loggedInUser,_ := c.Get("logged_in_user")
	data["is_logged_in"] = loggedInInterface.(bool)
	data["logged_in_user"] = loggedInUser.(string)
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}


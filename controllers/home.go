package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuops/go-example-app/config"
	"github.com/kuops/go-example-app/utils"
)

func Home(c *gin.Context){
	loggedInInterface, _ := c.Get("is_logged_in")
	environment := config.Environment
	c.Set("is_logged_in",false)
	loggedIn := loggedInInterface.(bool)
	if !loggedIn {
		c.Redirect(301,"/login")
	}
	utils.Render(c,gin.H{"home_active":"active","environment":environment},"index.html")
}

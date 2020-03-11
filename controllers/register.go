package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuops/go-example-app/models"
	"github.com/kuops/go-example-app/utils"
)

func PreRegister(c *gin.Context) {
	utils.Render(c,gin.H{},"register.html")
}

func Register(c *gin.Context) {
	var user models.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Email =c.PostForm("email")
	if user.FindUser() {
		c.String(400,"user exists.")
	} else {
		user.Create()
		c.String(200,"register successful.")
	}
}



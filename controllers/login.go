package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuops/go-example-app/models"
	"github.com/kuops/go-example-app/utils"
	"net/http"
)

func PreLogin(c *gin.Context) {
	utils.Render(c,gin.H{"title": "Successful Login"},"login.html")
}

func Login(c *gin.Context) {
	var user models.AuthUser
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	if err := models.CheckPassword(&user);err == nil {
		c.SetCookie("token", user.Email, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		c.Redirect(http.StatusMovedPermanently,"/")
	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"e_invalid":   "is-invalid",
			"e_message": "user or password error"})
	}
}

func Logout(c *gin.Context) {
	c.SetCookie("token","",-1,"","",false,false)
	utils.Render(c,gin.H{},"logout.html")
}


package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuops/go-example-app/models"
	"github.com/kuops/go-example-app/utils"
)

func GetUsers(c *gin.Context) {
	var ulist []models.User
	models.FindAllUser(&ulist)
	utils.Render(c,gin.H{"ulist":ulist,"users_active":"active"},"users-list.html")
}

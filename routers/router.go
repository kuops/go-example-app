package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuops/go-example-app/controllers"
	"github.com/kuops/go-example-app/middleware"
)

func InitRouters(r *gin.Engine) {
	r.Use(middleware.SetUserStatus())
	r.Use(gin.Logger())
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets","templates/assets")
	// Home
	r.GET("/",controllers.Home)
	// Login
	r.GET("/login", controllers.PreLogin)
	r.POST("/login", controllers.Login)
	r.GET("/logout",controllers.Logout)
	// Register
	r.GET("/register", controllers.PreRegister)
	r.POST("/register", controllers.Register)
	users := r.Group("/users")
	users.GET("list",controllers.GetUsers)
}

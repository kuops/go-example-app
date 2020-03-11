package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kuops/go-example-app/config"
	"github.com/kuops/go-example-app/database"
	"github.com/kuops/go-example-app/models"
	"github.com/kuops/go-example-app/routers"
	"github.com/kuops/go-example-app/utils"
	"github.com/spf13/pflag"
)

var (
	cfile string
)

func main() {
	pflag.Parse()
	pflag.StringVar(&cfile, "config", "config/config.yaml", "setting config file")
	config.InitConfig(cfile)
	database.DB = database.InitDatabase(config.Database)
	database.DB.AutoMigrate(&models.User{}, &models.Article{})
	utils.DatabaseStuff(database.DB)
	r := gin.New()
	routers.InitRouters(r)
	listen := fmt.Sprintf(":%s", config.Port)
	err := r.Run(listen)
	if err != nil {
		log.Panicf("server start failed, %v\n", err)
	}
}

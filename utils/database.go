package utils

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/jinzhu/gorm"
	"github.com/kuops/go-example-app/models"
	log "github.com/sirupsen/logrus"
)

var users []models.User
var articles []models.Article
var admin  = models.User{
	Username: "admin",
	Password: "admin",
	Email:    "admin@example.com",
}

func DatabaseStuff(db *gorm.DB){
	UserStuff(db)
	ArticleStuff(db)
}

func ArticleStuff(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.User{}, &models.Article{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v\n", err)
	}
	db.Debug().Find(&articles)
	if len(articles) == 0 {
		for i,v := range users {
			var article models.Article
			title := fmt.Sprintf("Test Article %v",i+1)
			content := fmt.Sprintf("Content for Test Article %v",i+1)
			article.Author = v.Username
			article.Title = title
			article.Content = content
			db.Debug().Create(&article)
		}
	}
}

func UserStuff(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.User{}, &models.Article{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v\n", err)
	}
	db.Debug().Find(&users)
	if len(users) == 0 {
		db.Debug().Create(&admin)
		for i := 0; i <= 10; i++ {
			var user models.User
			user.Username = randomdata.FirstName(2)
			user.Password = randomdata.RandStringRunes(5)
			user.Email = randomdata.Email()
			db.Debug().Create(&user)
		}
	}
}

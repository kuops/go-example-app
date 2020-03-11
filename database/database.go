package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var DB *gorm.DB

func InitDatabase(dbfile string) *gorm.DB{
	fmt.Println(dbfile)
	db,err := gorm.Open("sqlite3",dbfile)
	if err != nil {
		log.Panicf("database connect failed.")
	}
	return  db
}
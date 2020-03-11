package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/kuops/go-example-app/database"
	"strings"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

func (u *User)Verify() error {
	switch  {
	case strings.TrimSpace(u.Username) == "":
		return errors.New("user name is null.")
	case strings.TrimSpace(u.Password) == "":
		return errors.New("passowrd name is null.")
	case strings.TrimSpace(u.Email) == "":
		return errors.New("email name is null.")
	}
	return nil
}

func (u *User)FindUser() bool {
	nouser := database.DB.Debug().Where("email = ?",u.Email).First(&u).RecordNotFound()
	return  nouser
}

func FindAllUser(ulist *[]User) {
	database.DB.Find(&ulist)
}

func (u *User)Create() {
	database.DB.Debug().Create(&u)
}
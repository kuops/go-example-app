package models

import (
	"github.com/kuops/go-example-app/database"
)

type AuthUser struct {
	Email string
	Password string
}

func CheckPassword(au *AuthUser) error {
	var user User
	err := database.DB.Debug().Where("email = ? AND password = ?",au.Email,au.Password).First(&user).Error

	if err != nil {
		return err
	}
	return nil
}
package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	email string
	username string
	phone string
	password string
}

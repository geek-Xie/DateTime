package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Username string `gorm:"type:varchar(255);not null"`
	Phone    string `gorm:"type:varchar(11);not null"`
	Password string `gorm:"size:255;not null"`
}

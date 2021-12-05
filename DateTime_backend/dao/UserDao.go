package dao

import (
	"DateTime_backend/model"
	"github.com/jinzhu/gorm"
)

func GetUserByEmail(db *gorm.DB, email string) model.User {
	var user model.User
	db.Where("email = ?", email).First(&user)

	return user
}

func AddNewUserToDatabase(db *gorm.DB, user model.User)  {
	db.Create(&user)
}

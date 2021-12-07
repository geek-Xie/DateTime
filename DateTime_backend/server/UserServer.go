package server

import (
	"DateTime_backend/dao"
	"DateTime_backend/model"
	"DateTime_backend/utils"
)

func IsUserExist(email string) bool {
	user := dao.GetUserByEmail(utils.InitDB(), email)
	if user.Email == "" {
		return false
	}
	return true
}

func GetUserByEmail(email string) model.User {
	return dao.GetUserByEmail(utils.InitDB(), email)
}

func CreateNewUser(user model.User)  {
	dao.AddNewUserToDatabase(utils.InitDB(), user)
}

func CheckPassword(email string, password string) bool {
	dbPassword := dao.GetPasswordByEmail(utils.InitDB(), email)
	if dbPassword != password {
		return false
	}
	return true
}

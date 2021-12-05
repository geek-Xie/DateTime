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

func CreateNewUser(user model.User)  {
	dao.AddNewUserToDatabase(utils.InitDB(), user)
}

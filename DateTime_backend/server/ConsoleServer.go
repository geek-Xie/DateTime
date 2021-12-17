package server

import (
	"DateTime_backend/dao"
	"DateTime_backend/utils"
)

func GetEvents(userEmail string, selectDate string) {
	db := utils.InitDB()
	dao.GetEventsByEmailAndDate(db, userEmail, selectDate)
}

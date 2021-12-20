package server

import (
	"DateTime_backend/dao"
	"DateTime_backend/model"
	"DateTime_backend/utils"
)

func GetEvents(userEmail string, selectDate string) []model.Event {
	db := utils.InitDB()
	events := dao.GetEventsByEmailAndDate(db, userEmail, selectDate)
	return events
}

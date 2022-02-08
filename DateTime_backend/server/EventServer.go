package server

import (
	"DateTime_backend/dao"
	"DateTime_backend/model"
	"DateTime_backend/utils"
)

func CreateEvent(event model.Event) bool {
	db := utils.InitDB()
	pre_count := Counting()
	dao.AddEvent(db, event)
	count := Counting()
	if pre_count < count {
		return true
	} else {
		return false
	}
}

func Counting() int64 {
	db := utils.InitDB()
	count := dao.CountEvents(db)
	return count
}

func GetEventsByUseremail(useremail string) []model.Event {
	db := utils.InitDB()
	events := dao.GetEvents(db, useremail)
	return events
}

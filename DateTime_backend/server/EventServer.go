package server

import (
	"DateTime_backend/dao"
	"DateTime_backend/model"
	"DateTime_backend/utils"
)

func CreateEvent(event model.Event) {
	db := utils.InitDB()
	dao.AddEvent(db, event)
}

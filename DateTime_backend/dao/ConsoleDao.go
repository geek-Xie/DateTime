package dao

import (
	"DateTime_backend/model"
	"github.com/jinzhu/gorm"
)

func GetEventsByEmailAndDate(db *gorm.DB, userEmail string, selectDate string) []model.Event {
	var events []model.Event
	db.Where("start_date = ? AND user_email = ?", selectDate, userEmail).Find(&events)
	return events
}

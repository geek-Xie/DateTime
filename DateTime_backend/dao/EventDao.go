package dao

import (
	"DateTime_backend/model"
	"github.com/jinzhu/gorm"
)

func AddEvent(db *gorm.DB, event model.Event) {
	db.Create(&event)
}

func CountEvents(db *gorm.DB) int64 {
	var count int64
	db.Table("events").Count(&count)
	return count
}

func GetEvents(db *gorm.DB, useremail string) []model.Event {
	var events []model.Event
	db.Where("user_email = ? and end_date = ?", useremail, "").Find(&events)
	return events
}

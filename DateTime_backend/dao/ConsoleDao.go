package dao

import (
	"DateTime_backend/model"
	"github.com/jinzhu/gorm"
	"log"
)

func GetEventsByEmailAndDate(db *gorm.DB, userEmail string, selectDate string) {
	var events model.Event
	db.Where("start_date = ?", selectDate).First(&events)
	log.Println(events.UserEmail)
}

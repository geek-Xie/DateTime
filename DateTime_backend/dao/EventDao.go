package dao

import (
	"DateTime_backend/model"
	"github.com/jinzhu/gorm"
)

func AddEvent(db *gorm.DB, event model.Event) {
	db.Create(&event)
}

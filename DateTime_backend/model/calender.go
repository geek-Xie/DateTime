package model

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	UserEmail   string `gorm:"type:varchar(255);not null"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
	StartDate   string `gorm:"type:varchar(255);not null"`
	EndDate     string `gorm:"type:varchar(255);not null"`
	StartTime   string `gorm:"type:varchar(255);not null"`
	EndTime     string `gorm:"type:varchar(255);not null"`
}

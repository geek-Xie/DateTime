package controller

import (
	"DateTime_backend/model"
	"DateTime_backend/server"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

func CreateEvent(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	userEmail := requestMap["useremail"]
	title := requestMap["title"]
	description := requestMap["description"]
	date := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	startDate := strings.Split(date, " ")[0]
	startTime := strings.Split(date, " ")[1]
	event := model.Event{
		Model:       gorm.Model{},
		UserEmail:   userEmail,
		Title:       title,
		Description: description,
		StartDate:   startDate,
		EndDate:     "",
		StartTime:   startTime,
		EndTime:     "",
	}
	server.CreateEvent(event)
}

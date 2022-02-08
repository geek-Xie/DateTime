package controller

import (
	"DateTime_backend/model"
	"DateTime_backend/response"
	"DateTime_backend/server"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
	"time"
)

type EventItems struct {
	Title       string
	Description string
	StartDate   string
	StartTime   string
	L_T         string
}

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
	isAdded := server.CreateEvent(event)
	if isAdded {
		response.Response(c, http.StatusOK, 200, nil, "添加Event成功")
	} else {
		response.Response(c, http.StatusOK, 400, nil, "添加Event失败")
	}
}

func GetEvents(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	userEmail := requestMap["useremail"]
	events := server.GetEventsByUseremail(userEmail)
	//fmt.Println(events[1])
	var eventItems []EventItems
	for i := 0; i < len(events); i++ {
		eventItem := EventItems{
			Title:       events[i].Title,
			Description: events[i].Description,
			StartDate:   events[i].StartDate,
			StartTime:   events[i].StartTime,
			L_T:         events[i].Description,
		}
		eventItems = append(eventItems, eventItem)
	}
	if len(eventItems) == 0 {
		response.Response(c, http.StatusOK, 400, nil, "该用户目前暂无正在进行的Events")
	} else {
		response.Response(c, http.StatusOK, 200, gin.H{"events": eventItems}, "成功返回该用户正在进行的events")
	}

}

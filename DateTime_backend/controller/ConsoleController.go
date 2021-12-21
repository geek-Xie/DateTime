package controller

import (
	"DateTime_backend/response"
	"DateTime_backend/server"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type CtxInfo struct {
	Context   string `json:"context"`
	UserEmail string `json:"userInfo"`
}

type EventItem struct {
	Title       string
	Description string
	StartDate   string
	StartTime   string
	EndTime     string
}

func Console(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	activeYMD := requestMap["activeYMD"]
	userEmail := requestMap["userEmail"]
	//log.Println(activeYMD)
	//log.Println(userEmail)
	log.Println(time.Now())
	events := server.GetEvents(userEmail, activeYMD)
	var eventItems []EventItem
	for i := 0; i < len(events); i++ {
		eventitem := EventItem{
			Title:       events[i].Title,
			Description: events[i].Description,
			StartDate:   events[i].StartDate,
			StartTime:   events[i].StartTime,
			EndTime:     events[i].EndTime,
		}
		eventItems = append(eventItems, eventitem)
	}
	if len(events) != 0 {
		response.Response(c, http.StatusOK, 200, gin.H{"eventItems": eventItems}, "该日期有事件记录")
	} else {
		response.Response(c, http.StatusOK, 400, nil, "该日期无事件记录")
	}

}

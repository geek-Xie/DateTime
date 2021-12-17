package controller

import (
	"DateTime_backend/server"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type CtxInfo struct {
	Context   string `json:"context"`
	UserEmail string `json:"userInfo"`
}

func Console(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	activeYMD := requestMap["activeYMD"]
	userEmail := requestMap["userEmail"]
	//log.Println(activeYMD)
	//log.Println(userEmail)
	log.Println(time.Now())
	server.GetEvents(userEmail, activeYMD)

}

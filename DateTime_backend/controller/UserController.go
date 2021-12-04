package controller

import (
	"DateTime_backend/model"
	"DateTime_backend/response"
	"DateTime_backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context)  {

	c.JSON(http.StatusOK, gin.H{
		"message": "register",
	})

}

func Login(c *gin.Context)  {
	userInfo, _ := c.Get("user")
	response.Response(c, http.StatusOK, gin.H{"userInfo": userInfo}, "")
	var user model.User
	token, error := utils.ReleaseToken(user)
	// token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjAsImV4cCI6MTYzOTI0MjMzNSwiaWF0IjoxNjM4NjM3NTM1LCJpc3MiOiJEYXRlVGltZSIsInN1YiI6InVzZXIgdG9rZW4ifQ.DoJublbCIJ89jfEerYcpZ4EQul3wygQivxfpAV25I3k

	if error != nil {
		response.Response(c, http.StatusInternalServerError, nil, "token发放异常!")
		return
	}
	response.Response(c, http.StatusOK, gin.H{"token": token}, "token发放成功， 登录成功")
}
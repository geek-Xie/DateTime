package controller

import (
	"DateTime_backend/model"
	"DateTime_backend/response"
	"DateTime_backend/server"
	"DateTime_backend/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Register(c *gin.Context) {
	var requestUser = model.User{}
	c.Bind(&requestUser)
	userExist := server.IsUserExist(requestUser.Email)
	if !userExist {
		user := model.User{
			Model:    gorm.Model{},
			Email:    requestUser.Email,
			Username: requestUser.Username,
			Phone:    requestUser.Phone,
			Password: requestUser.Password,
		}
		server.CreateNewUser(user)
		token, err := utils.ReleaseToken(user)

		if err != nil {
			response.Response(c, http.StatusInternalServerError, nil, "token发放异常")
			return
		}

		response.Response(c, http.StatusCreated, gin.H{"user": user, "token": token}, "新用户创建成功")
	} else {
		response.Response(c, http.StatusOK, nil, "用户已存在")
	}

}

func Login(c *gin.Context) {
	var requestMap = make(map[string]string)
	json.NewDecoder(c.Request.Body).Decode(&requestMap)
	userExist := server.IsUserExist(requestMap["email"])
	if !userExist {
		response.Response(c, http.StatusOK, nil, "当前用户不存在，请先注册用户")
	} else {
		passwordOk := server.CheckPassword(requestMap["email"], requestMap["password"])
		if passwordOk {
			user := server.GetUserByEmail(requestMap["email"])
			token, _ := utils.ReleaseToken(user)
			response.Response(c, http.StatusOK, gin.H{"token": token}, "登陆成功，发放token")
		} else {
			response.Response(c, http.StatusOK, nil, "密码错误,请重新输入")
		}
	}

	//userInfo, _ := c.Get("user")
	//response.Response(c, http.StatusOK, gin.H{"userInfo": userInfo}, "")
	//var user model.User
	//token, error := utils.ReleaseToken(user)
	//// token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjAsImV4cCI6MTYzOTI0MjMzNSwiaWF0IjoxNjM4NjM3NTM1LCJpc3MiOiJEYXRlVGltZSIsInN1YiI6InVzZXIgdG9rZW4ifQ.DoJublbCIJ89jfEerYcpZ4EQul3wygQivxfpAV25I3k
	//
	//if error != nil {
	//	response.Response(c, http.StatusInternalServerError, nil, "token发放异常!")
	//	return
	//}
	//response.Response(c, http.StatusOK, gin.H{"token": token}, "token发放成功， 登录成功")
}

package controller

import (
	"DateTime_backend/model"
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
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
	var user model.User

	token, error := utils.ReleaseToken(user)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "token发放异常!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "token发放成功， 登录成功",
		"data": token,
	})
}
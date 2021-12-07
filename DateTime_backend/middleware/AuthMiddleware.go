package middleware

import (
	"DateTime_backend/model"
	"DateTime_backend/response"
	"DateTime_backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization的header
		tokenString := c.GetHeader("Authorization")

		// 验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(c, http.StatusUnauthorized, 400, nil, "权限不足")
			// 抛弃请求
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		// 解析token
		token, claim, err := utils.ParseToken(tokenString)

		// 解析出的token无效或者有错误
		if err != nil || !token.Valid {
			response.Response(c, http.StatusUnauthorized, 400, nil, "权限不足")
			c.Abort()
			return
		}

		// 获取token中的userId
		userId := claim.UserId
		db := utils.InitDB()
		var user model.User
		db.First(&user, userId)

		// 如果用户不存在
		//if user.ID == 0 {
		//	c.JSON(http.StatusUnauthorized, gin.H{
		//		"message": "该用户不存在",
		//	})
		//	c.Abort()
		//	return
		//}

		// 如果用户存在， 将user信息写入上下文

		c.Set("user", user)

		c.Next()

	}
}

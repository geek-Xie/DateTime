package response

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, HttpStatus int, code int, data gin.H, msg string)  {
	c.JSON(HttpStatus, gin.H{
		"code": code,
		"data": data,
		"msg": msg,
	})
}

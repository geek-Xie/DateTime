package response

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, HttpStatus int, data gin.H, msg string)  {
	c.JSON(HttpStatus, gin.H{
		"data": data,
		"msg": msg,
	})
}

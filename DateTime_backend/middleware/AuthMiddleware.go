package middleware

import (
	"DateTime_backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "权限不足",
			})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claim, err := utils.ParseToken(tokenString)

		if err != nil || token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "权限不足",
			})
			c.Abort()
			return
		}


		useId := claim.UserId
		c.JSON(http.StatusOK, gin.H{
			"message": useId,
		})
	}
}

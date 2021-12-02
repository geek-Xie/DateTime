package route

import (
	"DateTime_backend/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine  {
	r.GET("/auth/register", controller.Register)
	r.GET("/auth/login", controller.Login)

	return r
}

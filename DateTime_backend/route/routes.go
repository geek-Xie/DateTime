package route

import (
	"DateTime_backend/controller"
	"DateTime_backend/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine  {
	r.Use(middleware.CorsMiddleware())
	r.POST("/auth/register", controller.Register)
	r.POST("/auth/login", middleware.AuthMiddleware(), controller.Login)
	//r.POST("/auth/login", controller.Login)

	return r
}

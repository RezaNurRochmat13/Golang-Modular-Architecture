package user

import (
	"svc-modular-arch-go/modules/user/delivery/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(userRoutes *gin.Engine) {

	userRouteConfiguration := userRoutes.Group("/public/api/v1")
	{
		userRouteConfiguration.GET("users", handler.GetAllUsers)
	}
}

package product

import (
	"svc-modular-arch-go/modules/product/delivery/handler"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(userRoutes *gin.Engine) {

	productRouteConfiguration := userRoutes.Group("/public/api/v1")
	{
		productRouteConfiguration.GET("product", handler.GetAllProduct)
	}
}

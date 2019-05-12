package main

import (
	"fmt"

	productModulesConfiguration "svc-modular-arch-go/modules/product"
	userRouteConfiguration "svc-modular-arch-go/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello modular architecture Go :)")
	SetupRootRoutesApplication()

}

func SetupRootRoutesApplication() {
	rootRoutesApplicationConfiguration := gin.Default()

	userRouteConfiguration.UserRoutes(rootRoutesApplicationConfiguration)
	productModulesConfiguration.ProductRoutes(rootRoutesApplicationConfiguration)

	rootRoutesApplicationConfiguration.Run(":8080")
}

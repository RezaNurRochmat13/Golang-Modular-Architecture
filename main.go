package main

import (
	productModulesConfiguration "svc-modular-arch-go/modules/product"
	userModulesConfiguration "svc-modular-arch-go/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	SetupRootRoutesApplication()

}

func SetupRootRoutesApplication() {
	rootRoutesApplicationConfiguration := gin.Default()

	userModulesConfiguration.UserRoutes(rootRoutesApplicationConfiguration)
	productModulesConfiguration.ProductRoutes(rootRoutesApplicationConfiguration)

	rootRoutesApplicationConfiguration.Run(":8080")
}

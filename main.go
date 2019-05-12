package main

import (
	"fmt"

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

	rootRoutesApplicationConfiguration.Run(":8080")
}

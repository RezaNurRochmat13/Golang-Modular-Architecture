package main

import (
	"io"
	"log"
	"os"
	productModulesConfiguration "svc-modular-arch-go/modules/product"
	userModulesConfiguration "svc-modular-arch-go/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	SetupRootRoutesApplication()

}

func SetupRootRoutesApplication() {
	createLogFiles, errorCreate := os.Create("log/application.log")

	if errorCreate != nil {
		log.Fatal(errorCreate.Error())
	}

	gin.DefaultWriter = io.MultiWriter(createLogFiles, os.Stdout)

	rootRoutesApplicationConfiguration := gin.Default()

	userModulesConfiguration.UserRoutes(rootRoutesApplicationConfiguration)
	productModulesConfiguration.ProductRoutes(rootRoutesApplicationConfiguration)

	rootRoutesApplicationConfiguration.Run(":8080")
}

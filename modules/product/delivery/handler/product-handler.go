package handler

import (
	"github.com/gin-gonic/gin"
)

func GetAllProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "All data product",
	})
}

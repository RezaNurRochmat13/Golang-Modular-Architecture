package handler

import (
	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "All users data",
	})
}

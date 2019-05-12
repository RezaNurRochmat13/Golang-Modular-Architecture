package utility

import (
	"github.com/gin-gonic/gin"
)

func GlobalErrorException(errorMessage error, ctx *gin.Context) {
	ctx.JSON(500, gin.H{
		"error": errorMessage.Error,
	})
}

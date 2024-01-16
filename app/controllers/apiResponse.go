package controllers

import "github.com/gin-gonic/gin"

func ResponseWithError(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, gin.H{"status": false, "message": message})
}

func ResponseWithSuccess(ctx *gin.Context, code int, data interface{}, message string) {
	ctx.JSON(code, gin.H{"status": true, "data": data, "message": message})
}

func ResponseWithSuccessWithoutData(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{"status": true, "message": message})
}

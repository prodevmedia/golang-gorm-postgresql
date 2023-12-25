package controllers

import "github.com/gin-gonic/gin"

func ResponseWithError(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, gin.H{"status": false, "message": message})
}

func ResponseWithSuccess(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, gin.H{"status": true, "data": data})
}

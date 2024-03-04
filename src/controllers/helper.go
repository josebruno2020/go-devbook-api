package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func SendError(httpError HttpError, ctx *gin.Context) {
	ctx.JSON(httpError.StatusCode, gin.H{
		"error": httpError.Error,
	})
}

func SendValidationErrors(httpError HttpError, ctx *gin.Context) {
	ctx.JSON(httpError.StatusCode, gin.H{
		"errors": strings.Split(httpError.Error, "\n"),
	})
}

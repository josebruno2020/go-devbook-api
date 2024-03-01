package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c UserController) List(ctx *gin.Context) {
	users := []map[string]string{
		{
			"name": "legal",
		},
		{
			"name": "legal2",
		},
	}

	ctx.JSON(http.StatusOK, users)
}

func (c UserController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"top": "legal",
	})
}

func (c UserController) FindById(ctx *gin.Context) {
	log.Println(ctx.Params)
	ctx.JSON(http.StatusOK, gin.H{
		"ID recebido": ctx.Param("id"),
	})
}

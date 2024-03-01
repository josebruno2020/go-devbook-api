package router

import (
	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-devbook-api/src/router/routes"
)

func Setup() *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r)

	return r
}

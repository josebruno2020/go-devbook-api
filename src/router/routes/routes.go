package routes

import "github.com/gin-gonic/gin"

type Route struct {
	Uri            string
	Method         string
	Handle         func(*gin.Context)
	IsAuthRequired bool
}

func SetupRoutes(r *gin.Engine) {
	group := r.Group("/users")
	for _, userRoute := range UsersRoutes {
		group.Handle(userRoute.Method, userRoute.Uri, userRoute.Handle)
	}
}

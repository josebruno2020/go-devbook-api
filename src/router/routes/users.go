package routes

import (
	"net/http"

	"github.com/josebruno2020/go-devbook-api/src/controllers"
)

var userController = controllers.UserController{}

var UsersRoutes = []Route{
	{
		Uri:    "/",
		Method: http.MethodGet,
		Handle: userController.List,
	},
	{
		Uri:    "/",
		Method: http.MethodPost,
		Handle: userController.Create,
	},
	{
		Uri:    "/:id",
		Method: http.MethodGet,
		Handle: userController.FindById,
	},
	{
		Uri:    "/:id",
		Method: http.MethodPut,
		Handle: userController.Update,
	},
	{
		Uri:    "/:id",
		Method: http.MethodDelete,
		Handle: userController.Delete,
	},
}

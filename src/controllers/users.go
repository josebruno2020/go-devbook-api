package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-devbook-api/src/models"
	"github.com/josebruno2020/go-devbook-api/src/repositories"
)

type UserController struct{}

func (c UserController) List(ctx *gin.Context) {
	userRepository := repositories.NewUserRepository()
	users, err := userRepository.List()
	if err != nil {
		httpError := HttpError{http.StatusBadRequest, err.Error()}
		SendError(httpError, ctx)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c UserController) Create(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		httpError := HttpError{http.StatusBadRequest, err.Error()}
		SendValidationErrors(httpError, ctx)
		return
	}
	userRepository := repositories.NewUserRepository()
	id, err := userRepository.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	user.ID = id
	ctx.JSON(http.StatusCreated, user)
}

func (c UserController) Update(ctx *gin.Context) {
	userRepository := repositories.NewUserRepository()
	idParam := ctx.Param("id")
	idInt, _ := strconv.Atoi(idParam)
	user, err := userRepository.GetById(idInt)
	if err != nil {
		SendError(HttpError{http.StatusNotFound, "User not found"}, ctx)
		return
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		SendValidationErrors(HttpError{http.StatusBadRequest, err.Error()}, ctx)
		return
	}

	log.Println(user)
	if err := userRepository.UpdateById(user); err != nil {
		SendError(HttpError{http.StatusNotFound, err.Error()}, ctx)
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func (c UserController) FindById(ctx *gin.Context) {
	userRepository := repositories.NewUserRepository()
	idParam := ctx.Param("id")
	idInt, _ := strconv.Atoi(idParam)
	user, err := userRepository.GetById(idInt)
	if err != nil {
		httpError := HttpError{http.StatusNotFound, err.Error()}
		SendError(httpError, ctx)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c UserController) Delete(ctx *gin.Context) {
	userRepository := repositories.NewUserRepository()
	idParam := ctx.Param("id")
	idInt, _ := strconv.Atoi(idParam)
	if err := userRepository.DeleteById(idInt); err != nil {
		httpError := HttpError{http.StatusNotFound, err.Error()}
		SendError(httpError, ctx)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

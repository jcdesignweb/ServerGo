package controllers

import (
	"github.com/labstack/echo"
	"ServerGo/app/models"
	"net/http"
	"strconv"
)

type Users struct {}

var userModel = models.User{}

func (u Users) Routes(Server *echo.Echo) {
	Server.GET("/users/:id", func(context echo.Context) error {
		var response models.UserResponse

		pageId, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			return context.JSON(http.StatusBadRequest, "Check id param")
		} else {
			response, err = userModel.GetUsers(pageId)
		}

		return context.JSON(http.StatusOK, response)
	})
}
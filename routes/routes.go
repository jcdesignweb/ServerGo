package routes

import (
	"github.com/labstack/echo"
)
func StartRoutes() {
	e := echo.New()
	e.GET("/", index)
	e.GET("/users", get)
	e.GET("/users/:id", getID)
	e.POST("/users", create)
	e.DELETE("/users/:id", del)
	e.PUT("/users/:id", upd)

	e.Start(":8080")
}


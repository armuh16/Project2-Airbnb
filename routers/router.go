package routers

import (
	"alta/airbnb/constants"
	"alta/airbnb/controllers"
	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// ------------------------------------------------------------------
	// LOGIN & REGISTER USER
	// ------------------------------------------------------------------
	e.POST("/register", controllers.RegisterUsersController)
	e.POST("/login", controllers.LoginUsersController) // jwt login

	r := e.Group("")
	r.Use(echoMid.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/:id", controllers.GetUsersController)      //jwt
	r.DELETE("/users/:id", controllers.DeleteUserController) // jwt delete
	r.PUT("/users/:id", controllers.UpdateUserController)    // jwt put
	return e
}

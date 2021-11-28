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
	// ------------------------------------------------------------------
	// USER ROUTER
	// ------------------------------------------------------------------
	r.GET("/users/:id", controllers.GetUsersController)      //jwt
	r.DELETE("/users/:id", controllers.DeleteUserController) // jwt delete
	r.PUT("/users/:id", controllers.UpdateUserController)    // jwt put
	// ------------------------------------------------------------------
	// HOMESTAY ROUTER
	// ------------------------------------------------------------------
	e.GET("/homestays", controllers.GetHomeStayController)
	e.GET("/homestays/filter/:type", controllers.GetHomeStayFilterController)
	e.GET("/homestays/:id", controllers.GetHomeStayDetailController)
	r.GET("/homestays/my", controllers.GetMyHomestayController)
	r.POST("/homestays", controllers.CreateHomestayController)
	r.PUT("/homestays/:id", controllers.UpdateHomeStayController)
	r.DELETE("/homestays/:id", controllers.DeleteHomeStayController)
	// ------------------------------------------------------------------
	// HOMESTAY BOOKING
	// ------------------------------------------------------------------
	r.POST("/reservations", controllers.CreateBookingController)
	r.GET("/reservations", controllers.GetBookingController)
	r.POST("/reservations/check", controllers.CheckReserveController)

	return e
}

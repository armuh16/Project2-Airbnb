package routers

import (
	"alta/airbnb/constants"
	"alta/airbnb/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete},
	}))
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
	e.GET("/homestays/type/:type", controllers.GetHomeStayFilterTypeController)
	e.GET("/homestays/feature/:type", controllers.GetHomeStayFilterFeatureController)
	e.GET("/homestays/location/:request", controllers.GetHomeStayFilterLocationController)
	e.GET("/homestays/:id", controllers.GetHomeStayDetailController)
	r.GET("/homestays/my", controllers.GetMyHomestayController)
	r.POST("/homestays", controllers.CreateHomestayController)
	r.PUT("/homestays/:id", controllers.UpdateHomeStayController)
	r.DELETE("/homestays/:id", controllers.DeleteHomeStayController)
	// ------------------------------------------------------------------
	// HOMESTAY BOOKING
	// ------------------------------------------------------------------
	r.POST("/reservations", controllers.CreateBookingController)
	r.GET("/reservations", controllers.GetAllBookingHistoriesControllers)
	r.GET("/reservations/:id", controllers.GetBookingControllers)
	r.DELETE("/reservations/:id", controllers.CancelBookingController)
	e.POST("/reservations/check", controllers.ReservationCheckController)
	// ------------------------------------------------------------------
	// FACILITY
	// ------------------------------------------------------------------
	r.POST("/feature", controllers.InsertFeatureController)
	r.GET("/feature/:id", controllers.GetFeatureController)

	return e
}

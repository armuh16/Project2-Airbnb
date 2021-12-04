package controllers

import (
	"alta/airbnb/lib/database"
	responses "alta/airbnb/lib/response"
	"alta/airbnb/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReservationCheckController(c echo.Context) error {
	reqDate := models.BodyCheckIn{}
	c.Bind(&reqDate)
	book, longstay, totalprice := database.CheckAvailability(reqDate)
	if book > 0 {
		return c.JSON(http.StatusOK, responses.StatusSuccess("Not Available"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessAvail("Available, please book now", longstay, totalprice))
}

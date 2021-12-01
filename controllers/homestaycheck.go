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
	respon := database.CheckAvailability(reqDate)
	if respon > 0 {
		return c.JSON(http.StatusOK, responses.StatusSuccess("Not Available"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccess("Available, please book now"))
}

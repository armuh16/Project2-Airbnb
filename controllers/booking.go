package controllers

import (
	"alta/airbnb/lib/database"
	responses "alta/airbnb/lib/response"
	"alta/airbnb/middlewares"
	"alta/airbnb/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// Controller untuk memasukkan barang baru ke Booking
func CreateBookingController(c echo.Context) error {
	inputReq := models.BodyCheckIn{}
	c.Bind(&inputReq)
	format := "2006-01-02"
	checkIn, _ := time.Parse(format, inputReq.CheckIn)
	checkOut, _ := time.Parse(format, inputReq.CheckOut)
	input := models.Booking{
		Homestay_ID: inputReq.Homestay_ID,
		CheckIn:     checkIn,
		CheckOut:    checkOut,
	}
	logged := middlewares.ExtractTokenUserId(c)
	input.User_ID = int(logged)
	book, err := database.CreateBooking(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
	}
	database.AddLongstay(checkIn, checkOut, book.ID)
	database.AddHargaToReservation(input.Homestay_ID, book.ID)
	database.InsertDateToCalendar(book.Homestay_ID, book.ID)
	return c.JSON(http.StatusOK, responses.StatusSuccess("success"))
}

func GetBookingControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FalseParamResponse())
	}
	userId, _ := database.GetReservationOwner(id)
	if err != nil || userId == 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
	}
	logged := middlewares.ExtractTokenUserId(c)
	if int(logged) != userId {
		return c.JSON(http.StatusBadRequest, responses.StatusUnauthorized())
	}
	booking, _ := database.GetReservation(id)
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success", booking))
}

func CancelBookingController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FalseParamResponse())
	}
	userId, _ := database.GetReservationOwner(id)
	if err != nil || userId == 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
	}
	logged := middlewares.ExtractTokenUserId(c)
	if int(logged) != userId {
		return c.JSON(http.StatusBadRequest, responses.StatusUnauthorized())
	}
	database.CancelReservation(id)
	return c.JSON(http.StatusOK, responses.StatusSuccess("success"))
}

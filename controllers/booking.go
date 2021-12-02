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
	respon := database.CheckAvailability(inputReq)
	if respon > 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("The request date is already booked"))
	}
	now := time.Now()
	zona, _ := now.Zone()
	format := "2006-01-02 15:04:05 MST"
	timeIn := " 14:00:00 " + zona
	timeOut := " 12:00:00 " + zona
	checkIn, _ := time.Parse(format, inputReq.CheckIn+timeIn)
	if inputReq.CheckIn == inputReq.CheckOut {
		checkIn = time.Now()
	}
	checkOut, _ := time.Parse(format, inputReq.CheckOut+timeOut)
	input := models.Booking{
		Homestay_ID: inputReq.Homestay_ID,
		CheckIn:     checkIn,
		CheckOut:    checkOut,
	}
	logged := middlewares.ExtractTokenUserId(c)
	input.User_ID = int(logged)
	book, err := database.CreateBooking(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedBook())
	}
	database.AddLongstay(checkIn, checkOut, book.ID)
	database.AddHargaToReservation(input.Homestay_ID, book.ID)
	database.InsertDateToCalendar(book.Homestay_ID, book.ID)
	return c.JSON(http.StatusOK, responses.SuccessBook())
}

func GetBookingControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FalseParamResponse())
	}
	userId, _ := database.GetReservationOwner(id)
	if err != nil || userId == 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("Failed to Get Data"))
	}
	logged := middlewares.ExtractTokenUserId(c)
	if int(logged) != userId {
		return c.JSON(http.StatusBadRequest, responses.WrongIdBook())
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
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("Failed to Get Data"))
	}
	logged := middlewares.ExtractTokenUserId(c)
	if int(logged) != userId {
		return c.JSON(http.StatusBadRequest, responses.WrongIdBook())
	}
	database.CancelReservation(id)
	return c.JSON(http.StatusOK, responses.SuccessCancelBook())
}

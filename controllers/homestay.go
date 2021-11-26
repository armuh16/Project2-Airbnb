package controllers

import (
	"alta/airbnb/lib/database"
	responses "alta/airbnb/lib/response"
	"alta/airbnb/middlewares"
	"alta/airbnb/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateHomestayController(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)
	homestay := models.Homestay{}
	if err := c.Bind(&homestay); err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("status bad request"))
	}
	respon, err := database.InsertHomestay(homestay, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	if respon == nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("data is already exist"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccess("success create new homestay"))
}

func GetHomeStayController(c echo.Context) error {
	homestays, err := database.GetAllHomeStay()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success get homestay", homestays))
}

func GetMyHomestayController(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)
	homestays, err := database.GetMyHometay(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success get homestay", homestays))
}

func GetHomeStayFilterController(c echo.Context) error {
	request := c.Param("type")
	homestays, err := database.GetHomeStayByType(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success get homestay", homestays))
}

func GetHomeStayDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("invalid method"))
	}
	homestay, err := database.GetHomeStayDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	if homestay == nil {
		return c.JSON(http.StatusNotFound, responses.StatusFailed("data not found"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success get homestay", homestay))
}

func UpdateHomeStayController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("invalid method"))
	}
	homeRequest := models.HomeStayRespon{}
	user_id := middlewares.ExtractTokenUserId(c)
	if err := c.Bind(&homeRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("bad request"))
	}
	respon, err := database.EditHomestay(&homeRequest, id, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	if respon == nil {
		return c.JSON(http.StatusNotFound, responses.StatusFailed("data not found"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccess("success edit homestay"))
}

func DeleteHomeStayController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("invalid method"))
	}
	user_id := middlewares.ExtractTokenUserId(c)
	respon, err := database.DeleteHomestay(id, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	if respon == nil {
		return c.JSON(http.StatusNotFound, responses.StatusFailed("data not found"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccess("success delete homestay"))
}

func CreateHomestayControllerTest() echo.HandlerFunc {
	return CreateHomestayController
}

func GetMyHomestayControllerTest() echo.HandlerFunc {
	return GetMyHomestayController
}

func UpdateHomeStayControllerTest() echo.HandlerFunc {
	return UpdateHomeStayController
}
func DeleteHomeStayControllerTest() echo.HandlerFunc {
	return DeleteHomeStayController
}

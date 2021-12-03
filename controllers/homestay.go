package controllers

import (
	"alta/airbnb/lib/database"
	responses "alta/airbnb/lib/response"
	"alta/airbnb/middlewares"
	"alta/airbnb/models"
	"alta/airbnb/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateHomestayController(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)
	newHomestay := models.PostHomestayRequest{}
	if err := c.Bind(&newHomestay); err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("status bad request"))
	}
	addresses, lat, lng, e := util.GetGeocodeLocations(newHomestay.Address)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusFailed("cannot generate the address"))
	}
	homestay := models.Homestay{
		Name:        newHomestay.Name,
		Type:        newHomestay.Type,
		Description: newHomestay.Description,
		Price:       newHomestay.Price,
		Address:     newHomestay.Address,
		Latitude:    lat,
		Longitude:   lng,
	}
	respon, err := database.InsertHomestay(homestay, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	if respon == nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("data is already exist"))
	} else {
		if _, err := database.InsertFasilities(newHomestay.Facility, respon.ID); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
		}
	}
	if _, err := database.InsertAddress(respon.ID, addresses); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
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

func GetHomeStayFilterTypeController(c echo.Context) error {
	request := c.Param("type")
	homestays, err := database.GetHomeStayByType(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success get homestay", homestays))
}

func GetHomeStayFilterFeatureController(c echo.Context) error {
	request := c.Param("type")
	homestays, err := database.GetHomeStayByFacility(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success get homestay", homestays))
}

func GetHomeStayFilterLocationController(c echo.Context) error {
	request := c.Param("request")
	homestays, err := database.GetHomeStayByLocation(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success get homestay", homestays))
}

func GetHomeStayDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadGateway, responses.StatusFailed("invalid method"))
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
		return c.JSON(http.StatusBadGateway, responses.StatusFailed("invalid method"))
	}
	homeRequest := models.HomeStayRespon{}
	user_id := middlewares.ExtractTokenUserId(c)
	if err := c.Bind(&homeRequest); err != nil {
		return c.JSON(http.StatusBadGateway, responses.StatusFailed("bad request"))
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
		return c.JSON(http.StatusBadGateway, responses.StatusFailed("invalid method"))
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

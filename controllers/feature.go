package controllers

import (
	"alta/airbnb/lib/database"
	responses "alta/airbnb/lib/response"
	"alta/airbnb/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Controller untuk memasukkan barang baru ke feature
func InsertFeatureController(c echo.Context) error {
	// Mendapatkan data feature baru dari client
	input := models.Feature{}
	c.Bind(&input)
	duplicate, _ := database.GetFeatureByName(input.Feature_name)
	if duplicate > 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("feature was inputed, try input another feature"))
	}
	// Menyimpan data barang baru menggunakan fungsi InsertFeature
	_, e := database.InsertFeature(input)
	if e != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed to input feature"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccess("success to input feature"))
}

func GetFeatureController(c echo.Context) error {
	idFeature, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, responses.FalseParamResponse())
	}
	feature, err := database.GetFeature(idFeature)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success get feature", feature))
}

package controllers

import (
	"alta/airbnb/lib/database"
	responses "alta/airbnb/lib/response"
	"alta/airbnb/models"
	"net/http"

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

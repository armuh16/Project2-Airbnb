package controllers

import (
	"alta/airbnb/lib/database"
	responses "alta/airbnb/lib/response"
	"alta/airbnb/models"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PhotoControllers(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}

		des, err := os.Create(file.Filename)
		if err != nil {
			return err
		}

		if _, err := io.Copy(des, src); err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, responses.StatusSuccess)
}

func InsertPhotoController(c echo.Context) error {
	var photo models.Photo
	c.Bind(photo)
	_, err := database.InsertPhoto(&photo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed)
	}
	return c.JSON(http.StatusOK, responses.StatusSuccess)
}

func GetAllPhotoController(c echo.Context) error {
	_, err := database.GetAllPhoto()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed)
	}
	return c.JSON(http.StatusOK, responses.StatusSuccess)
}

func DeletePhotoController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FalseParamResponse())
	}
	_, err = database.DeletePhoto(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed)
	}
	return c.JSON(http.StatusOK, responses.StatusSuccess)
}

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

var user models.Users

//register user
func RegisterUsersController(c echo.Context) error {
	c.Bind(&user)
	duplicate, _ := database.GetUserByEmail(user.Email)
	if duplicate > 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("Email was used, try another email"))
	}

	Password, _ := database.GeneratehashPassword(user.Password)
	user.Password = Password
	user.Role = "customer"
	_, err := database.RegisterUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("bad request"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccess("success create new user"))
}

//login users
func LoginUsersController(c echo.Context) error {
	user := models.UserLogin{}
	c.Bind(&user)
	users, err := database.LoginUsers(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("bad request"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success login", users.Token))
}

//get user by id
func GetUsersController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("bad request"))
	}

	loginuser := middlewares.ExtractTokenUserId(c)
	if loginuser != id {
		return c.JSON(http.StatusUnauthorized, responses.StatusUnauthorized())
	}

	datauser, e := database.GetUser(id)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusFailedInternal)
	}

	respon := models.GetUserResponse{
		Name:        datauser.Name,
		Email:       datauser.Email,
		PhoneNumber: datauser.PhoneNumber,
		Gender:      datauser.Gender,
	}

	return c.JSON(http.StatusOK, responses.StatusSuccessData("success get user", respon))
}

//delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("bad request"))
	}
	_, error := database.DeleteUser(id)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusFailed("internal service error"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccess("success delete user"))
}

//update user by id
func UpdateUserController(c echo.Context) error {
	c.Bind(&user)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("bad request"))
	}
	_, e := database.UpdateUser(id, user)
	if e != nil {
		return c.JSON(http.StatusUnauthorized, responses.StatusUnauthorized())
	}
	return c.JSON(http.StatusOK, responses.StatusSuccess("success update user"))
}

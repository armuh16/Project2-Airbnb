package controllers

import (
	"alta/airbnb/config"
	"alta/airbnb/constants"
	"alta/airbnb/lib/database"
	"alta/airbnb/middlewares"
	"alta/airbnb/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

type HomestayResponSuccess struct {
	Status  string
	Message string
	Data    []models.Homestay
}

type ResponseFailed struct {
	Status  string
	Message string
}

//without slice
type SingleHomestayResponseSuccess struct {
	Status  string
	Message string
	Data    models.HomeStayRespon
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type PostHomeStayErr struct {
	ID   int
	Name int
}

var logininfo = LoginUserRequest{
	Email:    "root@gmail.com",
	Password: "root",
}

var (
	mock_data_homestay = models.Homestay{
		Name:        "villa mutiara",
		Type:        "villa",
		Description: "ini desc",
		Price:       1000,
		Latitude:    6.02,
		Longitude:   10.88,
		User_ID:     1,
	}
)

var (
	mock_data_user = models.Users{
		Name:        "root",
		Email:       "root@gmail.com",
		Password:    "root",
		PhoneNumber: "081222333444",
		Gender:      "male",
	}
)

func InsertMockDataHomestayToDB() error {
	var err error
	if err = config.DB.Save(&mock_data_homestay).Error; err != nil {
		return err
	}
	return nil
}

var xpass string

func InsertMockDataUserToDB() error {
	xpass, _ = database.GeneratehashPassword(mock_data_user.Password)
	mock_data_user.Password = xpass
	var err error
	if err = config.DB.Save(&mock_data_user).Error; err != nil {
		return err
	}
	return nil
}

func TestCreateHomestaySuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	body, err := json.Marshal(mock_data_homestay)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	var userDetail models.Users
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/homestays", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/homestays")
	middleware.JWT([]byte(constants.SECRET_JWT))(CreateHomestayControllerTest())(context)
	var homestay SingleHomestayResponseSuccess
	bodyRes := res.Body.String()
	json.Unmarshal([]byte(bodyRes), &homestay)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "success create new homestay", homestay.Message)
	assert.Equal(t, "success", homestay.Status)
}

func TestCreateProductFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	var userDetail models.Users
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	t.Run("TestPOST_Duplicate", func(t *testing.T) {
		InsertMockDataHomestayToDB()
		newbody, err := json.Marshal(mock_data_homestay)
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/homestays", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/homestays")
		middleware.JWT([]byte(constants.SECRET_JWT))(CreateHomestayControllerTest())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "data is already exist", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestPOST_ErrorBind", func(t *testing.T) {
		newbody, err := json.Marshal(PostHomeStayErr{})
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/homestays", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/homestays")
		middleware.JWT([]byte(constants.SECRET_JWT))(CreateHomestayControllerTest())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "status bad request", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestPOST_ErrorDB", func(t *testing.T) {
		newbody, err := json.Marshal(mock_data_homestay)
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/homestays", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/homestays")
		config.DB.Migrator().DropTable(&models.Homestay{})
		middleware.JWT([]byte(constants.SECRET_JWT))(CreateHomestayControllerTest())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "internal server error", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
}

func TestGetHomestaySuccess(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataHomestayToDB()
	req := httptest.NewRequest(http.MethodGet, "/homestays", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/homestays")
	if assert.NoError(t, GetHomeStayController(context)) {
		body := rec.Body.String()
		var homestay HomestayResponSuccess
		err := json.Unmarshal([]byte(body), &homestay)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get homestay", homestay.Message)
		assert.Equal(t, "success", homestay.Status)
		assert.Equal(t, 1, len(homestay.Data))
		assert.Equal(t, "villa mutiara", homestay.Data[0].Name)
	}
}

func TestGetHomestayFailed(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataHomestayToDB()
	req := httptest.NewRequest(http.MethodGet, "/homestays", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/homestays")
	config.DB.Migrator().DropTable(&models.Homestay{})
	if assert.NoError(t, GetHomeStayController(context)) {
		body := rec.Body.String()
		var homestay HomestayResponSuccess
		err := json.Unmarshal([]byte(body), &homestay)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "internal server error", homestay.Message)
		assert.Equal(t, "failed", homestay.Status)
	}
}

func TestGetMyHomestaySuccess(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	body, err := json.Marshal(mock_data_homestay)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	var userDetail models.Users
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	InsertMockDataHomestayToDB()
	req := httptest.NewRequest(http.MethodPost, "/homestays/my", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/homestays/my")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetMyHomestayControllerTest())(context)
	var homestay HomestayResponSuccess
	bodyRes := res.Body.String()
	json.Unmarshal([]byte(bodyRes), &homestay)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "success get homestay", homestay.Message)
	assert.Equal(t, "success", homestay.Status)
	assert.Equal(t, 1, len(homestay.Data))
	assert.Equal(t, "villa mutiara", homestay.Data[0].Name)
}

func TestGetMyHomestayFailed(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	body, err := json.Marshal(mock_data_homestay)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	var userDetail models.Users
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	InsertMockDataHomestayToDB()
	config.DB.AutoMigrate()
	req := httptest.NewRequest(http.MethodPost, "/homestays/my", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/homestays/my")
	config.DB.Migrator().DropTable(&models.Homestay{})
	middleware.JWT([]byte(constants.SECRET_JWT))(GetMyHomestayControllerTest())(context)
	var homestay HomestayResponSuccess
	bodyRes := res.Body.String()
	json.Unmarshal([]byte(bodyRes), &homestay)
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Equal(t, "internal server error", homestay.Message)
	assert.Equal(t, "failed", homestay.Status)
}

func TestGetHomestayFilterSuccess(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataHomestayToDB()
	req := httptest.NewRequest(http.MethodGet, "/homestays/filter/:type", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/homestays/filter/:type")
	context.SetParamNames("type")
	context.SetParamValues("villa")
	if assert.NoError(t, GetHomeStayFilterController(context)) {
		body := rec.Body.String()
		var homestay HomestayResponSuccess
		err := json.Unmarshal([]byte(body), &homestay)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get homestay", homestay.Message)
		assert.Equal(t, "success", homestay.Status)
		assert.Equal(t, 1, len(homestay.Data))
		assert.Equal(t, "villa", homestay.Data[0].Type)
	}
}

func TestGetHomestayFilterFailed(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataHomestayToDB()
	req := httptest.NewRequest(http.MethodGet, "/homestays/filter/:type", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/homestays/filter/:type")
	context.SetParamNames("type")
	context.SetParamValues("villa")
	config.DB.Migrator().DropTable(&models.Homestay{})
	if assert.NoError(t, GetHomeStayFilterController(context)) {
		body := rec.Body.String()
		var homestay HomestayResponSuccess
		err := json.Unmarshal([]byte(body), &homestay)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "internal server error", homestay.Message)
		assert.Equal(t, "failed", homestay.Status)
	}
}

func TestGetHomestayDetailSuccess(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataHomestayToDB()
	req := httptest.NewRequest(http.MethodGet, "/homestays/:id", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/homestays/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	if assert.NoError(t, GetHomeStayDetailController(context)) {
		body := rec.Body.String()
		var homestay HomestayResponSuccess
		err := json.Unmarshal([]byte(body), &homestay)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get homestay", homestay.Message)
		assert.Equal(t, "success", homestay.Status)
	}
}
func TestGetHomestayDetailFailed(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataHomestayToDB()
	t.Run("TestGETHomestayDetail_InvalidMethod", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/homestays/:id", nil)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath("/homestays/:id")
		context.SetParamNames("id")
		context.SetParamValues("#")
		if assert.NoError(t, GetHomeStayDetailController(context)) {
			body := rec.Body.String()
			var homestay HomestayResponSuccess
			err := json.Unmarshal([]byte(body), &homestay)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "invalid method", homestay.Message)
			assert.Equal(t, "failed", homestay.Status)
		}
	})
	t.Run("TestGETHomestayDetail_NotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/homestays/:id", nil)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath("/homestays/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")
		if assert.NoError(t, GetHomeStayDetailController(context)) {
			body := rec.Body.String()
			var homestay HomestayResponSuccess
			err := json.Unmarshal([]byte(body), &homestay)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "data not found", homestay.Message)
			assert.Equal(t, "failed", homestay.Status)
		}
	})
	t.Run("TestGETHomestayDetail_ErrorDB", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/homestays/:id", nil)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath("/homestays/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		config.DB.Migrator().DropTable(&models.Homestay{})
		if assert.NoError(t, GetHomeStayDetailController(context)) {
			body := rec.Body.String()
			var homestay HomestayResponSuccess
			err := json.Unmarshal([]byte(body), &homestay)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "internal server error", homestay.Message)
			assert.Equal(t, "failed", homestay.Status)
		}
	})
}

func TestUpdateHomestaySuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataHomestayToDB()
	var newdata = models.HomeStayRespon{
		Name:        "villa mutiara edited",
		Type:        "villa",
		Description: "ini desc",
		Price:       1000,
		Latitude:    6.02,
		Longitude:   10.88,
	}
	newbody, err := json.Marshal(newdata)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	var userDetail models.Users
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/homestays/:id", bytes.NewBuffer(newbody))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/homestays/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateHomeStayControllerTest())(context)
	var homestay SingleHomestayResponseSuccess
	body := res.Body.String()
	json.Unmarshal([]byte(body), &homestay)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "success edit homestay", homestay.Message)
	assert.Equal(t, "success", homestay.Status)
}
func TestUpdateHomestayFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataHomestayToDB()
	var newdata = models.HomeStayRespon{
		Name:        "villa mutiara edited",
		Type:        "villa",
		Description: "ini desc",
		Price:       1000,
		Latitude:    6.02,
		Longitude:   10.88,
	}
	newbody, err := json.Marshal(newdata)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	var userDetail models.Users
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	t.Run("TestEdiHomestayDetail_InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/homestays/:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/homestays/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateHomeStayControllerTest())(context)
		var homestay SingleHomestayResponseSuccess
		body := res.Body.String()
		json.Unmarshal([]byte(body), &homestay)
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "data not found", homestay.Message)
		assert.Equal(t, "failed", homestay.Status)
	})
	t.Run("TestEdiHomestayDetail_InvalidMethod", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/homestays/:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/homestays/:id")
		context.SetParamNames("id")
		context.SetParamValues("#")
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateHomeStayControllerTest())(context)
		var homestay SingleHomestayResponseSuccess
		body := res.Body.String()
		json.Unmarshal([]byte(body), &homestay)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "invalid method", homestay.Message)
		assert.Equal(t, "failed", homestay.Status)
	})
	t.Run("TestEdiHomestayDetail_ErrorBind", func(t *testing.T) {
		newbody, err := json.Marshal(PostHomeStayErr{})
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPut, "/homestays/:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/homestays/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateHomeStayControllerTest())(context)
		var homestay SingleHomestayResponseSuccess
		body := res.Body.String()
		json.Unmarshal([]byte(body), &homestay)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "bad request", homestay.Message)
		assert.Equal(t, "failed", homestay.Status)
	})
	t.Run("TestEdiHomestayDetail_ErrorDB", func(t *testing.T) {
		newbody, err := json.Marshal(newdata)
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPut, "/homestays/:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/homestays/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		config.DB.Migrator().DropTable(&models.Homestay{})
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateHomeStayControllerTest())(context)
		var homestay SingleHomestayResponseSuccess
		body := res.Body.String()
		json.Unmarshal([]byte(body), &homestay)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "internal server error", homestay.Message)
		assert.Equal(t, "failed", homestay.Status)
	})
}
func TestDeleteHomestaySuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataHomestayToDB()
	var userDetail models.Users
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/homestays/:id", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/homestays/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteHomeStayControllerTest())(context)
	var homestay SingleHomestayResponseSuccess
	body := res.Body.String()
	json.Unmarshal([]byte(body), &homestay)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "success delete homestay", homestay.Message)
	assert.Equal(t, "success", homestay.Status)
}
func TestDeleteHomestayFailed(t *testing.T) {
	e := InitEchoTestAPI()
	t.Run("TestDeleteHomestay_InvalidID", func(t *testing.T) {
		InsertMockDataUserToDB()
		InsertMockDataHomestayToDB()
		var userDetail models.Users
		tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
		if tx.Error != nil {
			t.Error(tx.Error)
		}
		token, err := middlewares.CreateToken(int(userDetail.ID))
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPut, "/homestays/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/homestays/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")
		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteHomeStayControllerTest())(context)
		var homestay SingleHomestayResponseSuccess
		body := res.Body.String()
		json.Unmarshal([]byte(body), &homestay)
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "data not found", homestay.Message)
		assert.Equal(t, "failed", homestay.Status)
	})
	t.Run("TestDeleteHomestay_InvalidMethod", func(t *testing.T) {
		InsertMockDataUserToDB()
		InsertMockDataHomestayToDB()
		var userDetail models.Users
		tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
		if tx.Error != nil {
			t.Error(tx.Error)
		}
		token, err := middlewares.CreateToken(int(userDetail.ID))
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPut, "/homestays/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/homestays/:id")
		context.SetParamNames("id")
		context.SetParamValues("#")
		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteHomeStayControllerTest())(context)
		var homestay SingleHomestayResponseSuccess
		body := res.Body.String()
		json.Unmarshal([]byte(body), &homestay)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "invalid method", homestay.Message)
		assert.Equal(t, "failed", homestay.Status)
	})
	t.Run("TestDeleteHomestay_ErrorDB", func(t *testing.T) {
		InsertMockDataUserToDB()
		InsertMockDataHomestayToDB()
		var userDetail models.Users
		tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
		if tx.Error != nil {
			t.Error(tx.Error)
		}
		token, err := middlewares.CreateToken(int(userDetail.ID))
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPut, "/homestays/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/homestays/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		config.DB.Migrator().DropTable(&models.Homestay{})
		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteHomeStayControllerTest())(context)
		var homestay SingleHomestayResponseSuccess
		body := res.Body.String()
		json.Unmarshal([]byte(body), &homestay)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "internal server error", homestay.Message)
		assert.Equal(t, "failed", homestay.Status)
	})

}

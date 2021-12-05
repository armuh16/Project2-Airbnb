package controllers

import (
	"alta/airbnb/lib/database"
	responses "alta/airbnb/lib/response"
	"alta/airbnb/middlewares"
	"alta/airbnb/models"
	"alta/airbnb/util"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

func CreateHomestayController(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)
	newHomestay := models.PostHomestayRequest{}
	if err := c.Bind(&newHomestay); err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("status bad request"))
	}

	bucket := "alta_airbnb"
	var err error

	ctx := appengine.NewContext(c.Request())
	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("credential.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"error":   true,
		})
	}

	f, uploaded_file, err := c.Request().FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"error":   true,
		})
	}
	defer f.Close()

	ext := strings.Split(uploaded_file.Filename, ".")
	extension := ext[len(ext)-1]
	photoname := string(uploaded_file.Filename)
	t := time.Now()
	formatted := fmt.Sprintf("%d%02d%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	homestay_name := strings.ReplaceAll(newHomestay.Name, " ", "+")
	uploaded_file.Filename = fmt.Sprintf("%v-%v.%v", homestay_name, formatted, extension)
	sw := storageClient.Bucket(bucket).Object(uploaded_file.Filename).NewWriter(ctx)
	if _, err := io.Copy(sw, f); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"error":   true,
		})
	}

	if err := sw.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"error":   true,
		})
	}

	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	addresses, lat, lng, e := util.GetGeocodeLocations(newHomestay.Address)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusFailed("cannot generate the address"))
	}
	homestay := models.Homestay{
		Name:        newHomestay.Name,
		Type:        newHomestay.Type,
		Description: newHomestay.Description,
		Guests:      newHomestay.Guests,
		Beds:        newHomestay.Beds,
		Bedrooms:    newHomestay.Bedrooms,
		Bathrooms:   newHomestay.Bathrooms,
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
		photo := models.Photo{
			Homestay_ID: respon.ID,
			Photo_Name:  photoname,
			Url:         fmt.Sprintf("%v", u),
		}
		if _, err := database.InsertPhoto(&photo); err != nil {
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
	homeRequest := models.PostHomestayRequest{}
	user_id := middlewares.ExtractTokenUserId(c)
	if err := c.Bind(&homeRequest); err != nil {
		return c.JSON(http.StatusBadGateway, responses.StatusFailed("bad request"))
	}
	respon, addresses, err := database.EditHomestay(&homeRequest, id, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	if respon == nil {
		return c.JSON(http.StatusNotFound, responses.StatusFailed("data not found"))
	} else {
		if _, err := database.EditFacilities(homeRequest.Facility, respon.ID); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
		}
	}
	if _, err := database.EditAddress(respon.ID, addresses); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
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

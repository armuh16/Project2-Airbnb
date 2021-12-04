package controllers

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

var uploader *s3manager.Uploader

func NewUploader() *s3manager.Uploader {
	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials("KeyID", "SecretKey", ""),
	}

	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)
	return uploader
}

func upload(test []byte) {
	log.Println("uploading")
	// file, err := ioutil.ReadFile(test)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	upInput := &s3manager.UploadInput{
		Bucket:      aws.String("airbnbupload"),         // bucket's name
		Key:         aws.String("myfiles/project2.jpg"), // files destination location
		Body:        bytes.NewReader(test),              // content of the file
		ContentType: aws.String("image/jpg"),            // content type
	}
	res, err := uploader.UploadWithContext(context.Background(), upInput)
	log.Printf("res %+v\n", res)
	log.Printf("err %+v\n", err)
}

func UploadController(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	filebyte, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	uploader = NewUploader()
	upload(filebyte)
	return nil
}

/*
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
*/

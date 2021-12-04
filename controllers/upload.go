package controllers

import (
	responses "alta/airbnb/lib/response"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string
var MyBucket string
var filepath string

//GetEnvWithKey : get env value
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

func ConnectAws() *session.Session {
	AccessKeyID = GetEnvWithKey("AWS_ACCESS_KEY_ID")
	SecretAccessKey = GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	MyRegion = GetEnvWithKey("AWS_REGION")

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})

	if err != nil {
		panic(err)
	}

	return sess
}

func UploadController(c echo.Context) error {
	sess := c.Get("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)

	MyBucket = GetEnvWithKey("BUCKET_NAME")
	file, header, _ := c.Request().FormFile("file")
	filename := header.Filename

	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusFailedInternal("Failed to upload file", up))
	}
	filepath = "https://" + MyBucket + "." + "s3-" + MyRegion + ".amazonaws.com/" + filename
	return c.JSON(http.StatusOK, responses.StatusSuccessData("filepath", filepath))
}

/*
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

	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}

	defer src.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}


	filebyte, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	uploader = NewUploader()
	upload(filebyte)
	return nil

	return c.JSON(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", file.Filename))

}


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

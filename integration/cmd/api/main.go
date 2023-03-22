package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var awsConfig = &aws.Config{
	Region:                        aws.String("us-east-1"),
	Endpoint:                      aws.String("http://localstack:4566"),
	S3ForcePathStyle:              aws.Bool(true),
	CredentialsChainVerboseErrors: aws.Bool(true),
}

type UploadRequest struct {
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(bodyDumpHandler))

	e.POST("/post", writeToS3)

	e.Logger.Fatal(e.Start(":8080"))
}

func writeToS3(c echo.Context) error {
	sess := session.Must(session.NewSession(awsConfig))
	var requestBody UploadRequest
	err := c.Bind(&requestBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	r := strings.NewReader(requestBody.Content)
	rs := io.ReadSeeker(r)

	svc := s3.New(sess)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("test"),
		Key:    aws.String(requestBody.Name),
		Body:   rs,
	})
	if err != nil {
		panic(err)
	}

	return c.NoContent(http.StatusAccepted)
}

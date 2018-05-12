package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"path"
	"time"

	"github.com/labstack/echo"

	"github.com/iwaaya/roomo/db"
	"github.com/iwaaya/roomo/obs"
)

const OBJECTKEY_FORMAT = "2006-01-02-15-04-05"

type Handler struct {
	db  *db.RDB
	obs *obs.OBS
}

func (h *Handler) AddImage(c echo.Context) error {
	fmt.Println("func AddImage")

	key := time.Now().Format(OBJECTKEY_FORMAT)

	// get an image from request body
	req := c.Request()
	body := new(bytes.Buffer)
	io.Copy(body, req.Body)
	buf := bytes.NewReader(body.Bytes())
	defer req.Body.Close()

	// upload an image to s3
	if err := h.obs.PutObject(key, buf); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	location := path.Join(h.obs.BaseURL, key)
	if err := h.db.CreateImage(location); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) GetImageList(c echo.Context) error {
	/*	sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-2")},
		)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, nil)
		}

		svc := s3.New(sess)

		result, err := svc.ListObjects(&s3.ListObjectsInput{
			Bucket: aws.String("roomo-test"),
		})
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, nil)
		}

		fmt.Println(result)*/
	return c.JSON(http.StatusOK, nil)

}

package main

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo"

	"github.com/iwaaya/roomo/db"
	"github.com/iwaaya/roomo/obs"
)

func TestAddImage(t *testing.T) {
	file, err := os.Open("./data/test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", file)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cf, err := NewConfig("./config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	d, err := db.New(cf.DB)
	if err != nil {
		t.Fatal(err)
	}

	o, err := obs.New(cf.OBS)
	if err != nil {
		t.Fatal(err)
	}

	h := &Handler{db: d, obs: o}

	h.AddImage(c)
}

func TestGetImageList(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cf, err := NewConfig("./config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	d, err := db.New(cf.DB)
	if err != nil {
		t.Fatal(err)
	}

	o, err := obs.New(cf.OBS)
	if err != nil {
		t.Fatal(err)
	}

	h := &Handler{db: d, obs: o}

	h.GetImageList(c)
	t.Log(rec.Body.String())
}

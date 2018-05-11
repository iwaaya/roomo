package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo"

	"github.com/iwaaya/roomo/db"
	"github.com/iwaaya/roomo/obs"
)

func main() {
	var confFile string
	flag.StringVar(&confFile, "config", "", "config file path")
	flag.Parse()

	cf, err := NewConfig(confFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	d, err := db.New(cf.DB)
	if err != nil {
		fmt.Println(err)
		return
	}

	o, err := obs.New(cf.OBS)
	if err != nil {
		fmt.Println(err)
		return
	}

	e := echo.New()
	h := &Handler{db: d, obs: o}

	e.POST("/collections", h.AddImage)
	e.GET("/collections", h.GetImageList)

	e.Logger.Fatal(e.Start(":8080"))
}

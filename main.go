package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/labstack/echo"

	"github.com/iwaaya/roomo/db"
	"github.com/iwaaya/roomo/obs"
)

func main() {
	if err := Run(os.Args[1:]...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run(args ...string) error {
	var confFile string
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.StringVar(&confFile, "config", "", "config file path")
	fs.Parse(args)

	cf, err := NewConfig(confFile)
	if err != nil {
		return err
	}

	d, err := db.New(cf.DB)
	if err != nil {
		return err
	}

	o, err := obs.New(cf.OBS)
	if err != nil {
		return err
	}

	e := echo.New()
	h := &Handler{db: d, obs: o}

	e.POST("/collections", h.AddImage)
	e.GET("/collections", h.GetImageList)

	return e.Start(":8080")
}

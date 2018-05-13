package main

import (
	"fmt"
	"os"
	"text/template"
)

type Config struct {
	DB_USER       string
	DB_PASSWORD   string
	DB_HOST       string
	OBS_ACCESSKEY string
	OBS_SECRETKEY string
	OBS_REGION    string
	OBS_URL       string
	OBS_BUCKET    string
}

func main() {
	t, err := template.ParseFiles("/etc/roomo/config.yaml.tpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	c := Config{
		DB_USER:       os.Getenv("DB_USER"),
		DB_PASSWORD:   os.Getenv("DB_PASSWORD"),
		DB_HOST:       os.Getenv("DB_HOST"),
		OBS_ACCESSKEY: os.Getenv("OBS_ACCESSKEY"),
		OBS_SECRETKEY: os.Getenv("OBS_SECRETKEY"),
		OBS_REGION:    os.Getenv("OBS_REGION"),
		OBS_URL:       os.Getenv("OBS_URL"),
		OBS_BUCKET:    os.Getenv("OBS_BUCKET"),
	}

	file, err := os.Create("/etc/roomo/config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	t.Execute(file, c)
}

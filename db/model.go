package db

import (
	"time"
)

type Collection struct {
	Id       int        `json:id`
	Date     *time.Time `json:date`
	Location string     `json:location`
}

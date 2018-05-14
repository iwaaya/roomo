package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	User string `yaml:user`
	Pass string `yaml:password`
	Host string `yaml:host`
}

type RDB struct {
	db *gorm.DB
}

func New(c Config) (*RDB, error) {
	dst := fmt.Sprintf("%s:%s@tcp(%s:3306)/roomo?charset=utf8&parseTime=True&loc=Local", c.User, c.Pass, c.Host)
	db, err := gorm.Open("mysql", dst)
	if err != nil {
		return nil, err
	}

	r := &RDB{db}
	return r, nil
}

func (r *RDB) CreateImage(location string) error {
	return r.db.Create(&Collection{Location: location}).Error
}

func (r *RDB) GetImageList() ([]Collection, error) {
	collections := []Collection{}
	err := r.db.Find(&collections).Error
	return collections, err
}

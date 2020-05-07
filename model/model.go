package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strings"
)

type Data struct {
	Path string `json:"path"`
	Data string `json:"data"`
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./data.db")
	if err == nil {
		DB = db
		return DB, err
	}
	return nil, err
}

func Get(path string) ([]*Data, error) {
	var data []*Data
	var err error
	path = strings.ToLower(path)
	err = DB.Where("path", path).Find(&data).Error
	return data, err
}

func (data *Data) Post() error {
	var err error
	err = DB.Create(data).Error
	return err
}

func (data *Data) Delete() error {
	var err error
	err = DB.Delete(data).Error
	return err
}

func (data *Data) Put() error {
	var err error
	err = DB.Update(data).Error
	return err
}

package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type Data struct {
	Path string `json:"path" gorm:"primary_key;type:string;not null"`
	Data string `json:"data" gorm:"type:string"`
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./data.db")
	if err == nil {
		DB = db
		db.AutoMigrate(&Data{})
		return DB, err
	} else {
		log.Fatal(err)
	}
	return nil, err
}

func Query(path string) (*Data, error) {
	data := new(Data)
	err := DB.Where("path = ?", path).First(&data).Error
	return data, err
}

func Create(data *Data) error {
	err := DB.Create(data).Error
	return err
}

func Remove(path string) error {
	err := DB.Where("path = ?", path).Delete(Data{}).Error
	return err
}

func Update(data *Data) error {
	err := DB.Where(Data{Path: data.Path}).Assign(Data{Data: data.Data}).FirstOrCreate(&data).Error
	return err
}

package db

import (
	"../models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB
var err error

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	db, _ = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/sys?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	err = db.AutoMigrate(models.User{}).Error
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.AutoMigrate(&models.Discussion{}).Error
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetDB() *gorm.DB {
	if db == nil {
		Init()
	}
	return db
}

func CloseDB() {
	db.Close()
}

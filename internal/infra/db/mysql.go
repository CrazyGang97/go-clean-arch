package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() {
	var err error
	db, err = gorm.Open(
		mysql.Open("connect your db"))
	if err != nil {
		panic(err)
	}
	SetDefault(db)
}

func DB() *gorm.DB {
	return db
}

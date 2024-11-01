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
		mysql.Open("root:123456@(9.134.234.41:3306)/clean_arch?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return db
}

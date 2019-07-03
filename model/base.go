package model

import (
    "github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() *gorm.DB{
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/go_test")
    defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	return db
}
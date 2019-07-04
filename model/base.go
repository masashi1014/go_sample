package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Model int = 22222

func DBConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/go_test")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Aaa() {
	fmt.Println("aaa")
}

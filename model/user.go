package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
}

func GetAllUsers() []User {
	dbInit()
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/go_test?parseTime=true")
	if err != nil {
        panic("データベース開けず！（dbDelete)")
    }
	user := []User{}
	db.Find(&user)

	return user
}

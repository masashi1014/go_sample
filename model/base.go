package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Model int = 22222

func dbInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/go_test")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getAll(db *gorm.DB) {
	name := "name"
	password := "pass"
	email := "mail"

	rec := User{
		Name:     name,
		Password: password,
		Email:    email,
	}

	db.NewRecord(&rec)
	db.Create(&rec)
}

func getRecord(db *gorm.DB) {
	name := "name"
	password := "pass"
	email := "mail"

	rec := User{
		Name:     name,
		Password: password,
		Email:    email,
	}

	db.NewRecord(&rec)
	db.Create(&rec)
}

func Insert(db *gorm.DB) {
	name := "name"
	password := "pass"
	email := "mail"

	rec := User{
		Name:     name,
		Password: password,
		Email:    email,
	}

	db.NewRecord(&rec)
	db.Create(&rec)
}

func Update(db *gorm.DB) {
	name := "name"
	password := "pass"
	email := "mail"

	rec := User{
		Name:     name,
		Password: password,
		Email:    email,
	}

	db.NewRecord(&rec)
	db.Create(&rec)
}

func Delete(db *gorm.DB) {
	name := "name"
	password := "pass"
	email := "mail"

	rec := User{
		Name:     name,
		Password: password,
		Email:    email,
	}

	db.NewRecord(&rec)
	db.Create(&rec)
}

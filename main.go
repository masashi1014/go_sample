package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "layout.html", gin.H{})
	})

	router.Run()
}

func DBConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/go_test")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	return db
}

//レコード挿入
func insert(db *gorm.DB) {
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

//レコード検索
func selecta(db *gorm.DB) {
	user := []User{}
	db.Select("name").Where("name = ?", "uehara").Find(&user)

	fmt.Println(user)
}

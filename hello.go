package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gin-gonic/gin"
)

type User struct {
    gorm.Model
    Name string `gorm:"size:255"`
    Password string `gorm:"size:255"`
    Email string `gorm:"size:255"`
}

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("./*.html")

    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(200, "layout.html", gin.H{})
    })

    router.Run()
}

//レコード挿入
func insert(db *gorm.DB) {
    name := "name"
    password := "pass"
    email := "mail"

    rec := User{
        Name: name,
        Password: password,
        Email: email,
    }

    db.NewRecord(&rec);
    db.Create(&rec);
}

//レコード検索
func selecta(db *gorm.DB) {
    user := []User{}
    db.Select("name").Where("name = ?", "uehara").Find(&user)

    fmt.Println(user)
}

package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masashi1014/go_sample/model"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*.html")

	router.GET("/", func(ctx *gin.Context) {
		users := model.GetAllUsers()
		ctx.HTML(200, "users.html", gin.H{
			"users": users,
		})
	})

	router.Run()
}

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masashi1014/go_sample/model"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*.html")

	//SELECE ALL(ユーザー一覧)
	router.GET("/", func(ctx *gin.Context) {
		users := model.GetAllUsers()
		ctx.HTML(200, "index.html", gin.H{
			"users": users,
		})
	})

	//INSERT(ユーザー新規作成)
	router.POST("/new", func(ctx *gin.Context) {

		//バリデーション
		var form model.User
		if err := ctx.Bind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest", "error": err})
			return
		}

		name := ctx.PostForm("name")
		password := ctx.PostForm("password")
		email := ctx.PostForm("email")
		new_user := model.User{Name: name, Password: password, Email: email}
		model.CreateUser(new_user)
		message := "作成しました"
		users := model.GetAllUsers()
		ctx.HTML(200, "index.html", gin.H{
			"users":   users,
			"message": message,
		})
	})

	//SELECT WHERE(ユーザー詳細画面)
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		user := model.GetUserDetail(id)
		ctx.HTML(200, "detail.html", gin.H{
			"user": user,
			"id":   id,
		})
	})

	//UPDATE(ユーザー更新)
	router.POST("/update/:id", func(ctx *gin.Context) {

		//バリデーション
		var form model.User
		if err := ctx.Bind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest", "error": err})
			return
		}

		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		name := ctx.PostForm("name")
		password := ctx.PostForm("password")
		email := ctx.PostForm("email")
		update_user := model.User{Name: name, Password: password, Email: email}
		model.UpdateUser(id, update_user)
		message := "更新しました"
		ctx.HTML(200, "detail.html", gin.H{
			"user":    update_user,
			"message": message,
		})
	})

	//DELETE確認画面表示
	router.GET("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		user := model.GetUserDetail(id)
		ctx.HTML(200, "delete.html", gin.H{
			"user": user,
			"id":   id,
		})
	})

	//DELETE(ユーザー削除)
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		model.DeleteUser(id)
		message := "削除しました"
		users := model.GetAllUsers()
		ctx.HTML(200, "index.html", gin.H{
			"users":   users,
			"message": message,
		})
	})

	router.Run()
}

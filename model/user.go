package model

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/masashi1014/go_sample/config"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Password string `form:"password" binding:"required" gorm:"size:255"` //パスワード入力必須
	Email    string `gorm:"size:255"`
}

//データベース接続
func dbInit() *gorm.DB {
	//conf.tomlからDBの接続情報を取得
	conf, err := config.NewConfig()
	if err != nil {
		panic(err.Error())
	}

	//DB接続
	db, err := gorm.Open("mysql", conf.DB.User+":"+conf.DB.Password+"@tcp("+conf.DB.Host+":"+strconv.Itoa(conf.DB.Port)+")/"+conf.DB.Db+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	// マイグレーション(テーブルがなかったら自動生成)
	db.AutoMigrate(&User{})
	return db
}

//ユーザー全権取得
func GetAllUsers() []User {
	db := dbInit()
	defer db.Close()

	user := []User{}
	db.Find(&user)

	return user
}

//ユーザー1件取得
func GetUserDetail(id int) User {
	db := dbInit()
	defer db.Close()

	user := User{}
	db.First(&user, id)

	return user
}

//ユーザー新規作成
func CreateUser(user User) {
	db := dbInit()
	defer db.Close()

	rec := User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}

	db.NewRecord(&rec)
	db.Create(&rec)
}

//ユーザー更新
func UpdateUser(id int, update_user User) {
	db := dbInit()
	defer db.Close()

	user := User{}
	db.First(&user, id)
	user.Name = update_user.Name
	user.Password = update_user.Password
	user.Email = update_user.Email

	db.Save(&user)
}

//ユーザー削除
func DeleteUser(id int) {
	db := dbInit()
	defer db.Close()

	user := User{}
	db.First(&user, id)
	db.Delete(&user)
}

package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Password string `form:"password" binding:"required" gorm:"size:255"` //パスワード入力必須
	Email    string `gorm:"size:255"`
}

//データベース接続
func dbInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/go_test")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	
	// マイグレーション(テーブルがなかったら自動生成)
    db.AutoMigrate(&User{})
	return db
}

//ユーザー全権取得
func GetAllUsers() []User {
	dbInit()
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/go_test?parseTime=true")
	defer db.Close()
	if err != nil {
		panic("DB open failed")
	}
	user := []User{}
	db.Find(&user)

	return user
}

//ユーザー1件取得
func GetUserDetail(id int) User {
	dbInit()
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/go_test?parseTime=true")
	defer db.Close()
	if err != nil {
		panic("DB open failed")
	}
	user := User{}
	db.First(&user, id)

	return user
}

//ユーザー新規作成
func CreateUser(user User) {
	dbInit()
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/go_test?parseTime=true")
	defer db.Close()
	if err != nil {
		panic("DB open failed")
	}
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
	dbInit()
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/go_test?parseTime=true")
	defer db.Close()
	if err != nil {
		panic("DB open failed")
	}

	user := User{}
	db.First(&user, id)
	user.Name = update_user.Name
	user.Password = update_user.Password
	user.Email = update_user.Email

	db.Save(&user)
}

//ユーザー削除
func DeleteUser(id int) {
	dbInit()
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/go_test?parseTime=true")
	defer db.Close()
	if err != nil {
		panic("DB open failed")
	}

	user := User{}
	db.First(&user, id)
	db.Delete(&user)
}

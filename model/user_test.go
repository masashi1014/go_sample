package model

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDbInit(t *testing.T) {
	db := dbInit()
	if db == nil {
		t.Fatal("failed test &db is unexpected type")
	}
}

func TestGetAllUsers(t *testing.T) {
	users := GetAllUsers()
	if len(users) == 0 {
		t.Error("no reccord")
	}
}

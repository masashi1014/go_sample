package model

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title   string `gorm:"size:255"`
	Content string `gorm:"size:255"`
	Flag    bool   `gorm:"size:1"`
}

package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string `gorm:"unique"`
	Body   string `gorm:"type:text"`
	Author string `gorm:"type:varchar(100)"`
}

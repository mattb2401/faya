package models

import (
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
)

type UserLocation struct {
	gorm.DB
	UserId int `gorm:"column:user_id"`
	Lat string `gorm:"column:lat"`
	Log string `gorm:"column:log"`
}
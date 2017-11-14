package models

import (
	_"github.com/jinzhu/gorm/dialects/mysql"
)

type CashHolder struct {
	UserId int `gorm:"column:userId"`
}
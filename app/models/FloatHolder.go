package models 

import (
	_"github.com/jinzhu/gorm/dialects/mysql"
)

type FloatHolder struct {
	UserId int `gorm:"colum:user_id"`
}
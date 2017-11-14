package models

import (
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
)

type Request struct {
	gorm.Model
	RequestUserId int `gorm:"column:request_user_id"`
	RequestType string `gorm:"column:request_type"`
	RequestAmount string `gorm:"column:request_amount"`
	RequestStatus string `gorm:"column:request_status"`
}

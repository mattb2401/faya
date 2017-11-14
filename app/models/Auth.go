package models

import (
    "time"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time `json:"-"`
    UpdatedAt time.Time `json:"-"`
    DeletedAt *time.Time `json:"-"`
  }

  type Auth struct {
	  gorm.Model
	  UserID uint `gorm:"column:user_id"`
	  AuthToken string
	  ExpireAt time.Time
  }

  type Auths []Auth
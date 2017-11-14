package models 

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

type User struct { 
    gorm.Model
    Fname string
    Lname string
    Email string `gorm:"not null;unique"`
    ProviderId string `grom:"column:provider_id"`
    Provider string
    Password string `json:"-"`
}

type Users []User
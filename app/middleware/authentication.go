package middleware

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"../models"
)

type ErrorRetn struct {
	Status string `json:"status"`    
	Code int `json:"code"`
	Message string `json:"message"`
}

func Authenticate(db *gorm.DB, w http.ResponseWriter, r *http.Request) (bool) {
	accessToken := r.Header.Get("access_token")
	if accessToken != "" {
		auth := models.Auth{}
		if err := db.Where("auth_token = ? AND expires_at > now()", accessToken).First(&auth).Error; err != nil {
			return false
		}else{
			return true
		}
	}else {
		return false
	}
}
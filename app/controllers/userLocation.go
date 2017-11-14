package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/jinzhu/gorm"
    "../models"
)

func Locate(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	type RequestParams struct {
		UserId int `json:"user_id"`
		Lat string `json:"lat"`
		Log string `json:"log"`
	}
	var p RequestParams
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		returnVals := ReturnVals{"FAILED", 203, "Missing or invalid json parameters"}
		js,_ := json.Marshal(returnVals)
		w.Header().Set("Content-Type" , "application/json")
		w.Write(js)
		return
	}
	user_loc := models.UserLocation{}
	if err := db.Where("user_id = ?", p.UserId).First(&user_loc).Error; err != nil {
		user_location := models.UserLocation{UserId: p.UserId, Lat: p.Lat, Log: p.Log}
		if err = db.Create(&user_location).Error; err != nil {
			returnVals := ReturnVals{"FAILED", 209, "Something has gone wrong. Location not saved."}
			js,_ := json.Marshal(returnVals)
			w.Header().Set("Content-Type" , "application/json")
			w.Write(js)
			return
		}else{
			returnVals := ReturnVals{"OK", 200, "Location saved."}
			js,_ := json.Marshal(returnVals)
			w.Header().Set("Content-Type" , "application/json")
			w.Write(js)
			return
		}
	}else {
		user_loc.Lat = p.Lat
		user_loc.Log = p.Log
		if err := db.Save(&user_loc).Error; err != nil {
			returnVals := ReturnVals{"FAILED", 209, "Location not saved. Something went wrong"}
			js,_ := json.Marshal(returnVals)
			w.Header().Set("Content-Type" , "application/json")
			w.Write(js)
			return
		}else{
			returnVals := ReturnVals{"OK", 200, "Location updated."}
			js,_ := json.Marshal(returnVals)
			w.Header().Set("Content-Type" , "application/json")
			w.Write(js)
			return
		}
	}
	
}
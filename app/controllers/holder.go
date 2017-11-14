package controllers 

import (
    "encoding/json"
    "net/http"
    "github.com/jinzhu/gorm"
    "../models"
)

func AddCashHolder(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	type Params struct {
		UserId int `json:user_id`
	}
	var p Params
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		returnVals := ReturnVals{"FAILED", 204, "Invalid parameters detected."}
		js,_ := json.Marshal(returnVals)
		w.Header().Set("Content-Type" , "application/json")
		w.Write(js)
		return
	}
	cash_holder := models.CashHolder{UserId: p.UserId}
	if err := db.Create(&cash_holder).Error; err != nil {
		returnVals := ReturnVals{"FAILED", 209, "Database error has occured."}
		js,_ := json.Marshal(returnVals)
		w.Header().Set("Content-Type" , "application/json")
		w.Write(js)
		return
	}else{
		returnVals := ReturnVals{"OK", 200, "You been added to the list of cash holders."}
		js,_ := json.Marshal(returnVals)
		w.Header().Set("Content-Type" , "application/json")
		w.Write(js)
		return
	}
}

func AddFloatHolder(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	type Params struct {
		UserId int `json:user_id`
	}
	var p Params
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		returnVals := ReturnVals{"FAILED", 204, "Invalid parameters detected."}
		js,_ := json.Marshal(returnVals)
		w.Header().Set("Content-Type" , "application/json")
		w.Write(js)
		return
	}
	float_holder := models.FloatHolder{UserId: p.UserId}
	if err := db.Create(&float_holder).Error; err != nil {
		returnVals := ReturnVals{"FAILED", 209, "Database error has occured."}
		js,_ := json.Marshal(returnVals)
		w.Header().Set("Content-Type" , "application/json")
		w.Write(js)
		return
	}else{
		returnVals := ReturnVals{"OK", 200, "You been added to the list of float holders."}
		js,_ := json.Marshal(returnVals)
		w.Header().Set("Content-Type" , "application/json")
		w.Write(js)
		return
	}
}






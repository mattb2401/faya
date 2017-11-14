package controllers 

import (
	"log"
    "encoding/json"
    "net/http"
    "github.com/jinzhu/gorm"
    "../models"
)

func Make(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	type RequestParams struct {
		RequestUserId int `json:"request_user_id"`
		RequestType string `json:"request_type"`
		RequestAmount string `json:"request_amount"`
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
	request := models.Request{RequestUserId: p.RequestUserId, RequestType: p.RequestType, RequestAmount: p.RequestAmount, RequestStatus: "Pending"}
	if err := db.Create(&request).Error; err != nil {
		log.Print(err)
		returnVals := ReturnVals{"FAILED", 203, "Something went wrong. We are looking into it."}
		js,_ := json.Marshal(returnVals)
		w.Header().Set("Content-Type" , "application/json")
		w.Write(js)
		return
	}
	returnVals := ReturnVals{"OK", 200, "Request is being processed."}
	js,_ := json.Marshal(returnVals)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}


func getRequests(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	type RequestParams struct {
		UserId int `json:"user_id"`
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
}


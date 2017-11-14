package controllers

import (
	"net/http/httptest"
	"bytes"
	"testing"
	"github.com/jinzhu/gorm"
	"net/http"
	"../../config"
	"fmt"
	"io/ioutil"
	"encoding/json"	
)


func TestTransaction(t *testing.T){
	config := config.GetConfig()
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=true",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
    if err != nil {
        panic(err)
	}
	bt := []byte(`{"user_id": 1, "lat": 32.44444, "log" : "0.4311111"}`)
	req, err := http.NewRequest("POST", "http://localhost:4000/v1/location/set", bytes.NewBuffer(bt))
	if err != nil {
		t.Fatalf("Error occured %v", err)
	}
	rec := httptest.NewRecorder()
	Locate(db, rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected error to be %v status but got %v", http.StatusOK,  res.StatusCode)
	}
	rBody, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()	
	if err != nil {
		t.Errorf("Could not read response. %v", err)
	}
	var rb map[string]interface{}
	err = json.Unmarshal(rBody, &rb)
	if err != nil {
		t.Errorf("Could not read body into json. %v", err)
	}
	if rb.Status != "OK"{

	}
}

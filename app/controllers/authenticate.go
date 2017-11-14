package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
    "../models"
    "time"
    "math/rand"
)



type  ReturnVals struct {
        Status string `json:"status"`    
        Code int `json:"code"`
        Message string `json:"message"`
}


type  ReturnInterface struct {
    Status string `json:"status"`    
    Code int `json:"code"`
    Result interface{} `json:"result"`
}

type SuccessRetn struct {
		Status string `json:"status"`
		Code int `json:"code"`
		User interface{} `json:"user"`
		AuthToken interface{} `json:"authToken"`
}

func Authenticate(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	type Params struct {
		Provider string `json:"provider", omitempty`
        Fname string `json:"fname", omitempty`
        Lname string `json:"lname", omitempty`
        Email string `json:"email", omitempty`
        Msisdn string `json:"msisdn", omitempty`
        ProviderId string `json:"provider_id, omitempty"`
        Password string `json:"password, omitempty"`
	}
	var p Params
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		returns := ReturnVals{"FAILED", 400, "Missing parameters or invalid json in request"}
		js, _ := json.Marshal(returns)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
	var returnvals ReturnVals
	users := models.User{}
    if p.Provider == "fb" || p.Provider == "gm" {
        if err := db.Where("provider_id = ? and provider = ?", p.ProviderId, p.Provider).First(&users).Error; err != nil {
            user := models.User{Fname: p.Fname, Lname: p.Lname, Provider: p.Provider, Email: p.Email, ProviderId: p.ProviderId}
            db.Create(&user)
            returns := ReturnInterface{"OK", 200, &user}
            js, _ := json.Marshal(returns)
            w.Header().Set("Content-Type", "application/json")
            w.Write(js)
            return
        }else {
            returns := ReturnInterface{"OK", 200, users}
            js, _ := json.Marshal(returns)
            w.Header().Set("Content-Type", "application/json")
            w.Write(js)
            return			
        }
    }else if p.Provider == "email" {
        if err := db.Where("email = ?", p.Email).First(&users).Error; err != nil {
            returns := ReturnVals{"FAILED", 203, "Wrong password or email."}
            js, _ := json.Marshal(returns)
            w.Header().Set("Content-Type", "application/json")
            w.Write(js)
            return
        }else{
            pass_err := CheckPasswordHash(p.Password, users.Password)
            if  pass_err == nil {
                returns := ReturnInterface{"OK", 200, &users}
                js, _ := json.Marshal(returns)
                w.Header().Set("Content-Type", "application/json")
                w.Write(js)
                return
            }else {
                returns := ReturnVals{"FAILED", 203, "Wrong password or email."}
                js, _ := json.Marshal(returns)
                w.Header().Set("Content-Type", "application/json")
                w.Write(js)
                return
            }
        }                
    
    }else {
        returnvals = ReturnVals{"FAILED", 400, "Invalid provider provided."}
        js , _ := json.Marshal(returnvals)
        w.Header().Set("Content-Type" , "application/json")
        w.Write(js)
        return
    }
}

// SignUp members that are using emails 

func SignUp(db *gorm.DB, w http.ResponseWriter, r *http.Request){
    type signUpParams struct {
        Fname string `json:"fname"`
        Lname string `json:"lname"`
        Email string `json:"email"`
        Password string `json:"password"`
    }
    decoder := json.NewDecoder(r.Body)
    var sp signUpParams
    err := decoder.Decode(&sp)
    if err != nil {
        returnvals := ReturnVals{"FAILED", 203, "Invalid Json request parameters"}
        js,_ := json.Marshal(returnvals)
        w.Header().Set("Content-Type" , "application/json")
        w.Write(js)
        return 
    }
    hashPass, err := HashPassword(sp.Password)
    if err != nil{
        returnvals := ReturnVals{"FAILED", 203, "Something has gone completely wrong. We are looking into it"}
        js , _ := json.Marshal(returnvals)
        w.Header().Set("Content-Type" , "application/json")
        w.Write(js)
        return 
    }
    allUser := models.User{}
    if err := db.Where("email = ?", sp.Email).First(&allUser).RecordNotFound; err != nil {
        user := models.User{Fname: sp.Fname, Lname: sp.Lname, Email: sp.Email, Password: hashPass, Provider: "email"}
        if err := db.Create(&user).Error; err != nil {
            returnvals := ReturnVals{"FAILED", 209, "Error occured while processing create."}
            js , _ := json.Marshal(returnvals)
            w.Header().Set("Content-Type" , "application/json")              
            w.Write(js)
            return
        }else{
            returnvals := ReturnVals{"OK", 200, "User has been created."}
            js , _ := json.Marshal(returnvals)
            w.Header().Set("Content-Type" , "application/json")              
            w.Write(js)
            return
        }
    }else{
        returnvals := ReturnVals{"FAILED", 203, "User already exists. Please try again."}
        js , _ := json.Marshal(returnvals)
        w.Header().Set("Content-Type" , "application/json")
        w.Write(js)
        return
    }
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil{
        return err
    }else{
        return nil
    }
	
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

const charset = "abcdefghijklmnopqrstuvwxyz" +  
"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(  
rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {  
b := make([]byte, length)
for i := range b {
  b[i] = charset[seededRand.Intn(len(charset))]
}
return string(b)
}

func generateToken(length int) string {  
return StringWithCharset(length, charset)
}






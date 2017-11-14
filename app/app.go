package app

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
	"../config"
	"./controllers"
    "log"
    "fmt"

)

type App struct {
    Route *mux.Router
    DB *gorm.DB
}

type ErrorRetn struct {
	Status string `json:"status"`    
	Code int `json:"code"`
	Message string `json:"message"`
}

func (a *App) Initialize(config *config.Config) {
    dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=true",config.DB.Username ,config.DB.Password ,config.DB.Name ,config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
    if err != nil {
        panic(err)
    }
    a.DB = db
    a.Route = mux.NewRouter()
    a.setRouters()
}

func (a *App) setRouters(){
    // Routing for authentication
	a.Post("/api/v1/authenticate", a.authenticate)
	a.Post("/api/v1/signup", a.signUp)
	a.Post("/api/v1/requests/make", a.makeRequest)
	a.Post("/api/v1/location/set", a.locate)

}

func (a *App) Get(path string, f func(w  http.ResponseWriter, r *http.Request)){
    a.Route.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)){
    a.Route.HandleFunc(path, f).Methods("POST")
}

func (a *App) signUp(w http.ResponseWriter, r *http.Request){
	controllers.SignUp(a.DB, w, r)
}

func (a *App) authenticate(w http.ResponseWriter, r *http.Request){
	controllers.Authenticate(a.DB, w, r)
}

func (a *App) makeRequest(w http.ResponseWriter, r *http.Request){
	// if middleware.Authenticate(a.DB, w, r) == true {
        controllers.Make(a.DB, w, r)
    // }else{
    //     returnvals := ErrorRetn{"FAILED", 401, "Bad authentication"}
    //     js, _ := json.Marshal(returnvals)
    //     w.Header().Set("Content-Type", "application/json")
    //     w.WriteHeader(http.StatusUnauthorized)
    //     w.Write(js)
    // }
}

func (a *App) locate(w http.ResponseWriter, r *http.Request){
	// if middleware.Authenticate(a.DB, w, r) == true {
		controllers.Locate(a.DB, w, r)
	// }else{
    //     returnvals := ErrorRetn{"FAILED", 401, "Bad authentication"}
    //     js, _ := json.Marshal(returnvals)
    //     w.Header().Set("Content-Type", "application/json")
    //     w.WriteHeader(http.StatusUnauthorized)
    //     w.Write(js)
    // }
}
// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Route))
}
package server

import (
	"net/http"

	"github.com/DanielMEH/apiclient/src/db"
	"github.com/DanielMEH/apiclient/src/models"
	"github.com/DanielMEH/apiclient/src/utils"
	"github.com/gorilla/mux"
)

func InitialServer() {
	db.DbConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router := mux.NewRouter()
	// cors
	router.Headers("Access-Control-Allow-Origin", "*")
	router.Headers("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	
	router.Headers("Content-Type", "application/json")
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("404"))
	})
	http.Handle("/", utils.HomeRouters())
	http.ListenAndServe(":3005", nil)

}

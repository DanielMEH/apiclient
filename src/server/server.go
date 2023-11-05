package server

import (
	"net/http"

	"github.com/DanielMEH/apiclient/src/utils"
	"github.com/gorilla/mux"
)

func InitialServer() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router := mux.NewRouter()
	router.Headers("Content-Type", "application/json")
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("404"))
	})
	http.Handle("/", utils.HomeRouters())
	http.ListenAndServe(":8080", nil)

}

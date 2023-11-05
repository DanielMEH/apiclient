package utils

import (
	"github.com/DanielMEH/apiclient/src/routers"
	"github.com/gorilla/mux"
)

func HomeRouters() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", routers.HomeHandle)

	return router

}

package main

import (
	"net/http"

	"github.com/penthaapatel/goodfoodsAPI/service"

	"github.com/gorilla/mux"
)

func main() {

	collection := service.ConnectToDB()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api", service.CreateDataHandler(collection)).Methods("POST")
	router.HandleFunc("/api", service.ViewAllDataHandler(collection)).Methods("GET")
	router.HandleFunc("/api/find/{id}", service.ViewDataByIDHandler(collection)).Methods("GET")
	router.HandleFunc("/api/delete/{id}", service.DeleteDataByIDHandler(collection)).Methods("DELETE")

	if err := http.ListenAndServe(":5000", router); err != nil {
		panic(err)
	}

}

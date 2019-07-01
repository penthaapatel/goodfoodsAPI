package main

import (
	"fmt"
	"log"
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
	router.HandleFunc("/api/update/{id}", service.UpdateDataByIDHandler(collection)).Methods("PUT")

	fmt.Println("Starting server on port :5000")
	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}

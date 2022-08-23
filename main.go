package main

import (
	"drone/db"
	"drone/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server is running!")
	DB, err := db.Init()
	if err != nil {
		log.Println("cant connect to database")
		return
	}
	h := server.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/api/drone", h.RegisterDrone).Methods(http.MethodPost)
	router.HandleFunc("/api/drone", h.GetDrones).Methods(http.MethodGet)
	router.HandleFunc("/api/drone/{id}/load-medication", h.LoadMedication).Methods(http.MethodPost)
	router.HandleFunc("/api/drone/{id}/check-medication", h.CheckingLoadedMedication).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)

}

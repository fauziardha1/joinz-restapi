package main

import (
	"joinz-api/pkg/api"
	"joinz-api/pkg/db"
	"log"
	"net/http"
)

func main() {
	_, err := db.SetupDB()
	if err != nil {
		log.Println("terjadi kesalahan pada setup db")
	}

	router := api.SetupApi()

	log.Println("We're otw running")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}

	log.Println("We're done running on 8080")

}

package main

import (
	"joinz-api/pkg/repository/db"
	"joinz-api/pkg/server"
	"log"
	"net/http"
)

func main() {
	pgDB, err := db.SetupDB()
	if err != nil {
		log.Println("terjadi kesalahan pada setup db")
	}

	router := server.InitApp(pgDB)

	log.Println("We're otw running")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}

	log.Println("We're done running on 8080")

}

package main

import (
	"fmt"
	"joinz-api/pkg/repository/db"
	"joinz-api/pkg/server"
	"log"
	"net/http"
	"os"
)

func main() {
	pgDB, err := db.SetupDB()
	if err != nil {
		log.Println("terjadi kesalahan pada setup db")
	}

	router := server.InitApp(pgDB)

	log.Println("We're up and running")
	port := os.Getenv("PORT")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}

	log.Println("We're done running on 8080")

}

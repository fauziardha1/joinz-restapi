package main

import (
	"joinz-api/pkg/api"
	"log"
	"net/http"
)

func main() {
	router := api.SetupApi()

	log.Println("We're otw running")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}

	log.Println("We're done running on 8080")

}

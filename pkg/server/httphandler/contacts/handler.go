package contacts

import "net/http"

// func GetContactsByUserID returns a list of contacts by user id
func GetContactsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetContactsByUserID"))
}

// func CreateNewContact is a function that creates a new contact
func CreateNewContact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateNewContact"))
}

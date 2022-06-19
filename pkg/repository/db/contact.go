package db

import (
	"log"

	"github.com/go-pg/pg/v10"
)

// func CreateContact is a function that creates a new contact
// it takes a request body of type CreateContactRequest
// it returns a success when the contact is created, otherwise it returns an error
func CreateContact(db *pg.DB, req *Contact) (*Contact, error) {

	// insert new contact to db
	_, err := db.Model(req).Insert()
	if err != nil {
		log.Printf("error inserting contact: %s\n", err)
		return nil, err
	}

	// get the new contact from db
	// Select contact by primary key.
	contact := &Contact{ID: req.ID}
	err = db.Model(contact).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return contact, nil
}

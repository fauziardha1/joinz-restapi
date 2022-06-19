package contacts

import (
	"encoding/json"
	"joinz-api/pkg/entity"
	"joinz-api/pkg/repository/db"
	"log"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
)

// func GetContactsByUserID returns a list of contacts by user id
func GetContactsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetContactsByUserID"))
}

// func CreateNewContact is a function that creates a new contact
func CreateNewContact(w http.ResponseWriter, r *http.Request) {
	// parse the request body into a contact
	req := &entity.CreateContactRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &entity.CreateContactResponse{
			Success: false,
			Error:   err.Error(),
			Contact: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error encoding response: %s\n", err)
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &entity.CreateContactResponse{
			Success: false,
			Error:   err.Error(),
			Contact: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error get database from context: %s\n", err)
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// prepare data for insert
	data := &db.Contact{}
	data.UserID = req.UserID
	data.Name = req.Name
	data.Email = req.Email
	if req.Phone != "" {
		data.Phone = req.Phone
	} else {
		data.Phone = ""
	}

	if req.Relation != "" {
		data.Relationship = req.Relation
	} else {
		data.Relationship = ""
	}

	if req.ProfilePicture != "" {
		data.ProfilePicture = req.ProfilePicture
	} else {
		data.ProfilePicture = ""
	}

	if req.AlternativeName != "" {
		data.AlternativeName = req.AlternativeName
	} else {
		data.AlternativeName = ""
	}

	if req.Message != "" {
		data.Messages = req.Message
	} else {
		data.Messages = ""
	}

	if req.Priority != 0 {
		data.Priorities = req.Priority
	} else {
		data.Priorities = 0
	}

	data.CreatedAt = time.Now().Format("20060102150405")
	data.UpdatedAt = time.Now().Format("20060102150405")

	// insert new contact to db
	contact, err := db.CreateContact(pgdb, data)
	if err != nil {
		res := &entity.CreateContactResponse{
			Success: false,
			Error:   err.Error(),
			Contact: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error encoding response: %s\n", err)
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return response
	res := &entity.CreateContactResponse{
		Success: true,
		Error:   "",
		Contact: contact,
	}
	_ = json.NewEncoder(w).Encode(res)
	return
}

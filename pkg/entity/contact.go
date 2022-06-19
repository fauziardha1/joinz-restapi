package entity

import "joinz-api/pkg/repository/db"

// CreateContact request is a struct that represents a request to create a new contact
type CreateContactRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email,omitempty"`
	Phone           string `json:"phone"`
	UserID          int64  `json:"user_id"`
	Relation        string `json:"relation,omitempty"`
	Priority        int    `json:"priority,omitempty"`
	ProfilePicture  string `json:"profile_picture,omitempty"`
	Message         string `json:"message,omitempty"`
	AlternativeName string `json:"alternative_name,omitempty"`
}

// CreateContactResponse is a struct that represents a response to create a new contact
type CreateContactResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Contact *db.Contact
}

// GetContactByUserIDRequest is a struct that represents a request to get a contact by user id
type GetContactByUserIDRequest struct {
	UserID int64 `json:"user_id"`
}

// GetContactByUserIDResponse is a struct that represents a response to get a contacts by user id
type GetContactByUserIDResponse struct {
	Success  bool   `json:"success"`
	Error    string `json:"error,omitempty"`
	Contacts []*db.Contact
}

package model

import "time"

// Rents Table Structure
type Rents struct {
	ID     int `json:"id" gorm:"primary_key"`
	BookID int
	UserID int

	ExpectedReturnDate time.Time `json:"expected_return_date"`
	BackCheck          int       `json:"backcheck"`

	Model
	Books Books
	Users Users
}

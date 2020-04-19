package rent

import "time"

// CreateParams is define when to create create
type CreateParams struct {
	BookID int `json:"book_id"`
	UserID int `json:"user_id"`
}

// UpdateParams is define when to create create
type UpdateParams struct {
	BookID             int       `json:"book_id"`
	UserID             int       `json:"user_id"`
	ExpectedReturnDate time.Time `json:"expected_return_date"`
	BackCheck          int       `json:"backcheck"`
}

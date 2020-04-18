package model

import "time"

// Model is Datetime settings
type Model struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

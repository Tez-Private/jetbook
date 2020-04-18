package model

// Users Table Structure
type Users struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Model
}

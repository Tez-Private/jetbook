package user

// CreateParams is define when to create create
type CreateParams struct {
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email"`
	Password string `json:"Password" gorm:"not null"`
}

// UpdateParams is define when to create create
type UpdateParams struct {
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email"`
	Password string `json:"Password" gorm:"not null"`
}

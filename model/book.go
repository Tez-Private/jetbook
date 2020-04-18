package model

// Books Table Structure
type Books struct {
	ID int `json:"id" gorm:"primary_key"`

	Title string `json:"title" gorm:"not null"`
	Isbn  string `json:"isbn"`

	Rent int `json:"rent"`
	Model
}

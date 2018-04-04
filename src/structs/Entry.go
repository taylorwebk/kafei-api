package structs

import (
	"github.com/jinzhu/gorm"
)

// Entry Structure for the model Link
type Entry struct {
	gorm.Model
	UserID  uint32  `gorm:"index" json:"user_id"`
	Literal string  `json:"literal"`
	Amount  float64 `json:"amount"`
}

// EntryResponse Entry struct for json response
type EntryResponse struct {
	ID      uint    `json:"id"`
	Date    string  `json:"date"`
	Time    string  `json:"time"`
	Literal string  `json:"literal"`
	Amount  float64 `json:"amount"`
}

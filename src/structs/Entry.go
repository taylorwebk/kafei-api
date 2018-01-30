package structs

import "github.com/jinzhu/gorm"

// Entry Structure for the model Link
type Entry struct {
	gorm.Model
	UserID  uint32  `gorm:"index" json:"user_id"`
	Literal string  `json:"literal"`
	Amount  float64 `json:"amount"`
}

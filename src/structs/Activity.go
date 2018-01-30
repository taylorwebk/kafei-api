package structs

import "github.com/jinzhu/gorm"

// Activity Structure for the model Link
type Activity struct {
	gorm.Model
	UserID  uint32   `gorm:"index" json:"user_id"`
	Literal string   `json:"literal"`
	Status  string   `gorm:"type:enum('STARTED', 'ENDED'); default:'STARTED'" json:"status"`
	Intervs []Interv `json:"intervals"`
}

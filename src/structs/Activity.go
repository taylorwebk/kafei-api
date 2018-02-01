package structs

import (
	"github.com/jinzhu/gorm"
	"github.com/taylorwebk/kafei-api/src/database"
)

// Activity Structure for the model Link
type Activity struct {
	gorm.Model
	UserID  uint32   `gorm:"index" json:"user_id"`
	Literal string   `json:"literal"`
	Status  string   `gorm:"type:enum('STARTED', 'ENDED'); default:'STARTED'" json:"status"`
	Intervs []Interv `json:"intervals"`
}

// LoadIntervals loads intervals
func (a *Activity) LoadIntervals() {
	db := database.SQL
	var intervals []Interv
	db.Where("activity_id = ?", a.ID).Find(&intervals)
	a.Intervs = intervals
}

// AddInterval ads an interval
func (a *Activity) AddInterval(item Interv) {
	a.Intervs = append(a.Intervs, item)
}

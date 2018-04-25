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

// ActivityResponse response structure for activity
type ActivityResponse struct {
	ID      uint    `json:"id"`
	Literal string  `json:"literal"`
	Date    string  `json:"date"`
	Time    string  `json:"time"`
	Hours   float64 `json:"hours"`
	Mins    float64 `json:"minutes"`
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

package structs

import "time"
import "github.com/jinzhu/gorm"

// Interv Structure for the model Link
type Interv struct {
	gorm.Model
	ActivityID uint32    `gorm:"index" json:"activity_id"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
}

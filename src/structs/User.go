package structs

import (
	"github.com/jinzhu/gorm"
	"github.com/taylorwebk/kafei-api/src/database"
)

// User Structure for the model Link
type User struct {
	gorm.Model
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Entrys    []Entry    `json:"entries"`
	Activitys []Activity `json:"activitys"`
}

// UserToken data for the token autentication
type UserToken struct {
	ID       uint
	Username string
}

// LoadActivities load activities with intervals
func (u *User) LoadActivities() {
	var activities []Activity
	database.SQL.Where("user_id = ?", u.ID).Find(&activities)
	for _, activity := range activities {
		activity.LoadIntervals()
	}
	u.Activitys = activities
}

// LoadEntries Load entries for user u
func (u *User) LoadEntries() {
	var entries []Entry
	database.SQL.Where("user_id = ?", u.ID).Find(&entries)
	u.Entrys = entries
}

// AddActivity ads an activity
func (u *User) AddActivity(item Activity) {
	u.Activitys = append(u.Activitys, item)
}

// WithActivities Returns a basic json for User struct
func (u *User) WithActivities() interface{} {
	return &struct {
		Name       string     `json:"name"`
		Email      string     `json:"email"`
		Username   string     `json:"username"`
		Activities []Activity `json:"activities"`
	}{
		Name:       u.Name,
		Email:      u.Email,
		Username:   u.Username,
		Activities: u.Activitys,
	}
}

// AddEntry ads an Entry
func (u *User) AddEntry(item Entry) {
	u.Entrys = append(u.Entrys, item)
}

// WithEntries Returns a basic json for User struct
func (u *User) WithEntries() interface{} {
	return &struct {
		Name     string  `json:"name"`
		Email    string  `json:"email"`
		Username string  `json:"username"`
		Entrys   []Entry `json:"entries"`
	}{
		Name:     u.Name,
		Email:    u.Email,
		Username: u.Username,
		Entrys:   u.Entrys,
	}
}

// NewStruct Returns a new Struct
func (u *User) NewStruct() interface{} {
	return &struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}{
		Name:     u.Name,
		Email:    u.Email,
		Username: u.Username,
	}
}

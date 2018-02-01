package main

import (
	"encoding/json"

	"github.com/taylorwebk/kafei-api/src/config"
	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/structs"
)

func main() {
	config.Load("../../config/config.json", conf)
	database.Connect(conf.Database)
	database.SQL.SingularTable(true)
	DB := database.SQL
	DB.CreateTable(&structs.User{}, &structs.Entry{}, &structs.Activity{}, &structs.Interv{})
}

var conf = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database database.Info `json:"Database"`
	KeyJWT   string        `json:"KeyJWT"`
	Static   string        `json:"Static"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

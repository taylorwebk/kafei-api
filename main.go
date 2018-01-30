package main

import (
	"os"

	"github.com/taylorwebk/kafei-api/src/config"
	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/routes"
)

func main() {
	config.Load("src"+string(os.PathSeparator)+"config"+string(os.PathSeparator)+"config.json", config.Conf)
	database.Connect(config.Conf.Database)
	database.SQL.SingularTable(true)
	defer database.SQL.Close()
	routes.RunAndServe()
}

package database

import (
	"fmt"
	"log"
	"os"

	//_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //MysqlDriver
)

var (
	// SQL wrapper
	SQL *gorm.DB
	// Database info
	databases Info
	//SERVER: Descomentar 19,20,21
	connectionName = mustGetenv("CLOUDSQL_CONNECTION_NAME")
	user           = mustGetenv("CLOUDSQL_USER")
	password       = os.Getenv("CLOUDSQL_PASSWORD")
)

// Type is the type of database from a Type* constant
type Type string

const (
	// TypeMySQL is MySQL
	TypeMySQL Type = "MySQL"
)

// Info contains the database configurations
type Info struct {
	// Database type
	Type Type
	// MySQL info if used
	MySQL MySQLInfo
}

// MySQLInfo is the details for the database connection
type MySQLInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

// Server Descomentar50-56
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}

// DSN returns the Data Source Name
// SERVER: Descomentar 60,61 y comentar 63-71. LOCAL: !SERVER
func DSN(ci MySQLInfo) string {
	constr := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/kafeidb%s", user, password, connectionName, ci.Parameter)
	return constr
	// Example: root:@tcp(localhost:3306)/test
	/* return ci.Username +
	":" +
	ci.Password +
	"@tcp(" +
	ci.Hostname +
	":" +
	fmt.Sprintf("%d", ci.Port) +
	")/" +
	ci.Name + ci.Parameter */
}

// Connect to the database
func Connect(d Info) {
	var err error

	// Store the config
	databases = d

	switch d.Type {
	case TypeMySQL:
		// Connect to MySQL
		if SQL, err = gorm.Open("mysql", DSN(d.MySQL)); err != nil {
			fmt.Println("ERROR: ", err.Error())
		}
	default:
		log.Println("No registered database in config")
	}
}

// ReadConfig returns the database information
func ReadConfig() Info {
	return databases
}

package main

import (
	"os"

	"github.com/taylorwebk/kafei-api/src/config"
	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/routes"
	"google.golang.org/appengine"
)

func main() {
	config.Load("src"+string(os.PathSeparator)+"config"+string(os.PathSeparator)+"config.json", config.Conf)
	database.Connect(config.Conf.Database)
	database.SQL.SingularTable(true)
	defer database.SQL.Close()
	routes.RunAndServe()
	appengine.Main()
}

/* package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/appengine"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var (
		connectionName = mustGetenv("CLOUDSQL_CONNECTION_NAME")
		user           = mustGetenv("CLOUDSQL_USER")
		password       = os.Getenv("CLOUDSQL_PASSWORD") // NOTE: password may be empty
	)

	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/", user, password, connectionName))
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	http.HandleFunc("/", handler)
	appengine.Main()
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not query db: %v", err), 500)
		return
	}
	defer rows.Close()

	buf := bytes.NewBufferString("Databases:\n")
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			http.Error(w, fmt.Sprintf("Could not scan result: %v", err), 500)
			return
		}
		fmt.Fprintf(buf, "- %s\n", dbName)
	}
	w.Write(buf.Bytes())
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
} */

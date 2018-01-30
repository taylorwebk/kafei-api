package routes

import (
	"github.com/gorilla/mux"
	"github.com/taylorwebk/kafei-api/src/actions"
	"github.com/taylorwebk/kafei-api/src/middlewares"
	"github.com/urfave/negroni"
)

// RunAndServe funcion que Inicia la App
func RunAndServe() {
	r := mux.NewRouter().StrictSlash(false)
	pr := mux.NewRouter().PathPrefix("/kafeiapi/user").Subrouter().StrictSlash(false)
	pr.HandleFunc("/entry", actions.NewEntry).Methods("POST")
	pr.HandleFunc("/entries", actions.AllEntries).Methods("GET")

	r.HandleFunc("/kafeiapi/register", actions.RegisterUser).Methods("POST")
	r.HandleFunc("/kafeiapi/login", actions.Login).Methods("POST")
	r.PathPrefix("/kafeiapi/user").Handler(negroni.New(
		negroni.HandlerFunc(middlewares.Login),
		negroni.Wrap(pr),
	))
	n := negroni.Classic()
	//n.Use(negroni.HandlerFunc(middlewares.Login))
	n.UseHandler(r)
	n.Run(":8080")
}

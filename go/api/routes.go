package api

import (
	"github.com/gorilla/mux"
)

// Router =: properties of all the routes available for the server to respond.
type Router struct {
	Root *mux.Router
}

func (a *API) setupRoutes() {
	a.Router.Root = a.MainRouter
	a.InitRoutes()
}

// InitRoutes := intializing all the endpoints
func (a *API) InitRoutes() {

	a.Router.Root.Handle("/", a.requestHandler(a.landing)).Methods("GET")
	a.Router.Root.Handle("/swagger-testing", a.requestHandler(a.swaggerTest)).Methods("POST")
}

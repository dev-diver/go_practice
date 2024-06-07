package router

import (
	"github.com/gorilla/mux"
)

// NewRouter initializes the main router and adds subrouters
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	addMemberRoutes(r)
	return r
}

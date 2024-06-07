package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type profile struct {
	Name string
	Age  int
}

var profiles = map[string]*profile{
	"devdiver": {"devdiver", 23},
	"go":       {"go", 10},
}

func addMemberRoutes(r *mux.Router) {
	memberRouter := r.PathPrefix("/member").Subrouter()
	memberRouter.HandleFunc("/{name}", getProfile).Methods("GET")
	memberRouter.HandleFunc("/{name}", updateProfile).Methods("POST")
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	profile := profiles[name]

	fmt.Fprintf(w, "member %s 's profile in %d\n", profile.Name, profile.Age)
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	profile := profiles[name]

	var data struct {
		Age int `json:"age"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	age := data.Age

	profile.Age = age
	fmt.Fprintf(w, "%s change age to %d", name, age)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	type profile struct {
		Name string
		Age  int
	}

	profiles := make(map[string]*profile)
	profiles["devdiver"] = &profile{"devdiver", 23}
	profiles["go"] = &profile{"go", 10}

	r := mux.NewRouter()

	memberRouter := r.PathPrefix("/member").Subrouter()
	memberRouter.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		profile := profiles[name]

		fmt.Fprintf(w, "member %s 's profile in %d\n", profile.Name, profile.Age)
	}).Methods("GET")

	memberRouter.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
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
	})

	fs := http.FileServer(http.Dir("file/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	// .Handle이 http는 /로 끝나면 이하 모든 url을 핸들링하지만,  mux는 정확히 일치하는 것만 핸들링
	// 따라서 mux는 PathPrefix를 써야 한다.

	http.ListenAndServe(":80", r)
}

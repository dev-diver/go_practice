package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/member/{name}/age/{age}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := vars["age"]

		fmt.Fprintf(w, "member %s 's profile in %s\n", name, age)
	})

	fs := http.FileServer(http.Dir("file/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	// .Handle이 http는 /로 끝나면 이하 모든 url을 핸들링하지만,  mux는 정확히 일치하는 것만 핸들링
	// 따라서 mux는 PathPrefix를 써야 한다.

	http.ListenAndServe(":80", r)
}

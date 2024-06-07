package main

import (
	"net/http"
	"practice/web/router"
)

func main() {

	r := router.NewRouter()

	fs := http.FileServer(http.Dir("file/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	// .Handle이 http는 /로 끝나면 이하 모든 url을 핸들링하지만,  mux는 정확히 일치하는 것만 핸들링
	// 따라서 mux는 PathPrefix를 써야 한다.

	http.ListenAndServe(":80", r)
}

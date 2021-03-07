package main

import (
	"github.com/balamuru/goweb/pkg/utils"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi, the time is now: " + utils.NowTime())
		
		
	}).Methods("GET")

	r.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars["name"])
		fmt.Fprintln(w, "Hi ", vars["name"] , ". the time is now: " + utils.NowTime() )
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}

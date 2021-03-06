package main

import (
	"github.com/balamuru/goweb/pkg/utils"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi " + time.Now().String())
		NowTime()
	}).Methods("GET")

	r.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars["name"])
		fmt.Fprintln(w, "Hi " + time.Now().String())
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}

package main

import (
	"fmt"
	"net/http"
	// "time"

	"github.com/gorilla/mux"
	"github.com/balamuru/goweb/pkg/utils"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		// fmt.Println()
		// fmt.Fprintln(w, "Hi " + time.Now().String()time.Now().String())
		fmt.Fprintln(w, "Under construction")
		//do something ...
		
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}


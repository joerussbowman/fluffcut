package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the fluffcut api")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login will happen here")
}

func main() {
	r := mux.NewRouter()
        r.StrictSlash(true)
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", LoginHandler)
	http.Handle("/", r)
        http.ListenAndServe(":8080", nil)
}

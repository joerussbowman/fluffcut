package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Root of the api, at some point this might be a good place
// to document the api for developers.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the fluffcut api")
}

// This is the base for the login. Apps should be able to direct
// applications to this end point for login. This documentation
// will be flushed out a lot more.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login will happen here")
}

// The landing point for oauth2 callbacks from providers. Got this
// from starting to set up for Google login, realized I like this idea
// of a callback point for each type of authentication technology. Should
// make for easier to read code.
func Oauth2CallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "landing point for oauth2callbacks")
}

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", LoginHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func contactPageHandler(w http.ResponseWriter, r *http.Request) {

	// setting headers
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Contact bitch, Fuck this shit server you at contact page niggaa.</h1>")
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	// setting headers
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Fuck this shit server you at homepage niggaa.</h1>")
}

func notFoundPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404 bitch, Fuck this shit server you at page not found niggaa.</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePageHandler)
	r.HandleFunc("/contact", contactPageHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundPageHandler)
	http.ListenAndServe(":3000", r)
}

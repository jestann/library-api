package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// main view and serve static assets
	// http.HandleFunc("/", homeHandler)
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	err := http.ListenAndServe(":"+os.Getenv(`PORT`), r)
	/* if err != nil {
	    log.Fatal("ListenAndServe: ", err)
	} */

	r.Handle("/status", NotImplemented).Methods("GET")
	r.Handle("/products", NotImplemented).Methods("GET")
	r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")

	// Here we are implementing the NotImplemented handler. Whenever an API endpoint is hit
	// we will simply return the message "Not Implemented"
	var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Not Implemented"))
	})
}

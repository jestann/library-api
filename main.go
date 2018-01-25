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
}

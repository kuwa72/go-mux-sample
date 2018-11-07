package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// MultipartHandler is handler
func MultipartHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "name: %s\n, %#v, %#v", r.PostFormValue("name"), r.FormValue("name"), r)
}

func main() {
	fmt.Println("Server Start")
	r := mux.NewRouter()
	r.HandleFunc("/", MultipartHandler).Methods("POST")

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

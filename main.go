package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/courses", getCourses).Methods("GET")

	srv := &http.Server{
		Handler:      http.TimeoutHandler(router, 3*time.Second, "Timeout"),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	time.Sleep(6 * time.Second)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

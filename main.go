package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	fmt.Println(apiKey)
	if port == "" {
		port = "3000"
	}
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	}).Methods("GET")

	r.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Fprintln(w, "Hello World", id)
	}).Methods("GET")

	log.Printf("Server started on port :%v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
	mux.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		v := r.PathValue("id")
		fmt.Fprintln(w, "Hello World ", v)
	})
	log.Println("server started on 127.0.0.1:3000")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", mux))
}

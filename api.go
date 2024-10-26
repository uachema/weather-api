package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	router     *mux.Router
}

func NewAPIServer(listenAddr string) *APIServer {
	server := &APIServer{
		listenAddr: listenAddr,
		router:     mux.NewRouter(),
	}
	server.setupRoutes()
	return server
}

func (s *APIServer) setupRoutes() {
	s.router.HandleFunc("/", s.handleRoot).Methods("GET")
	s.router.HandleFunc("/{id}", s.handleRootID).Methods("GET")
}

func (s *APIServer) Run() error {
	log.Printf("Server started on port %s", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, s.router)
}

func (s *APIServer) handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func (s *APIServer) handleRootID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Hello World %s\n", id)
}

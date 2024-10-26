package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	listenAddr := fmt.Sprint(":", port)
	server := NewAPIServer(listenAddr)
	if err := server.Run(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

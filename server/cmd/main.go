package main

import (
	"fmt"
	"log"
	"net/http"
	"server/internal/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", api.HealthHandler)
	mux.HandleFunc("/handleReq", api.HandleReq)

	port := "8080"
	fmt.Printf("Server starting in port: %s", port)

	log.Fatal(http.ListenAndServe(":"+port, mux))

}

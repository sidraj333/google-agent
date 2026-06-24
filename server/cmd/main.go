package main

import (
	"fmt"
	"log"
	"net/http"
	"server/internal/api"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", api.HealthHandler)
	mux.HandleFunc("/handleReq", api.HandleReq)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server starting in port: %s\n", port)

	log.Fatal(http.ListenAndServe(":" + port, mux))

}

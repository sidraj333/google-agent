package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/internal/agent"
	"server/internal/auth"
	"server/internal/health"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.HealthHandler)
	mux.HandleFunc("/handleReq", agent.HandleReq)
	mux.HandleFunc("/auth", auth.HandleAuth)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server starting in port: %s\n", port)

	log.Fatal(http.ListenAndServe(":" + port, mux))

}

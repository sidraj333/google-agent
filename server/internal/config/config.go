package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port               string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	GCPProjectID       string
	JWTSecret          string
}

func LoadConfig() (Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	if googleClientID == "" {
		return Config{}, fmt.Errorf("GOOGLE_CLIENT_ID is required")
	}

	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if googleClientSecret == "" {
		return Config{}, fmt.Errorf("GOOGLE_CLIENT_SECRET is required")
	}

	googleRedirectURL := os.Getenv("GOOGLE_REDIRECT_URL")
	if googleRedirectURL == "" {
		return Config{}, fmt.Errorf("GOOGLE_REDIRECT_URL is required")
	}

	gcpProjectID := os.Getenv("GCP_PROJECT_ID")
	if gcpProjectID == "" {
		return Config{}, fmt.Errorf("GCP_PROJECT_ID is required")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return Config{}, fmt.Errorf("JWT_SECRET is required")
	}

	return Config{
		Port:               port,
		GoogleClientID:     googleClientID,
		GoogleClientSecret: googleClientSecret,
		GoogleRedirectURL:  googleRedirectURL,
		GCPProjectID:       gcpProjectID,
		JWTSecret:          jwtSecret,
	}, nil
}

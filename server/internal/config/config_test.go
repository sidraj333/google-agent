package config

import "testing"

func TestLoadConfigSuccess(t *testing.T) {
	t.Setenv("PORT", "9000")
	t.Setenv("GOOGLE_CLIENT_ID", "client-id")
	t.Setenv("GOOGLE_CLIENT_SECRET", "client-secret")
	t.Setenv("GOOGLE_REDIRECT_URL", "http://localhost/callback")
	t.Setenv("GCP_PROJECT_ID", "my-project")
	t.Setenv("JWT_SECRET", "secret-key")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.Port != "9000" {
		t.Fatalf("expected port %q, got %q", "9000", cfg.Port)
	}

	if cfg.GoogleClientID != "client-id" {
		t.Fatalf("expected client id %q, got %q", "client-id", cfg.GoogleClientID)
	}

	if cfg.GoogleClientSecret != "client-secret" {
		t.Fatalf("expected client secret %q, got %q", "client-secret", cfg.GoogleClientSecret)
	}

	if cfg.GoogleRedirectURL != "http://localhost/callback" {
		t.Fatalf("expected redirect url %q, got %q", "http://localhost/callback", cfg.GoogleRedirectURL)
	}

	if cfg.GCPProjectID != "my-project" {
		t.Fatalf("expected gcp project id %q, got %q", "my-project", cfg.GCPProjectID)
	}

	if cfg.JWTSecret != "secret-key" {
		t.Fatalf("expected jwt secret %q, got %q", "secret-key", cfg.JWTSecret)
	}
}

func TestLoadConfigDefaultPort(t *testing.T) {
	t.Setenv("GOOGLE_CLIENT_ID", "client-id")
	t.Setenv("GOOGLE_CLIENT_SECRET", "client-secret")
	t.Setenv("GOOGLE_REDIRECT_URL", "http://localhost/callback")
	t.Setenv("GCP_PROJECT_ID", "my-project")
	t.Setenv("JWT_SECRET", "secret-key")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.Port != "8080" {
		t.Fatalf("expected default port %q, got %q", "8080", cfg.Port)
	}
}

func TestLoadConfigMissingClientID(t *testing.T) {
	t.Setenv("GOOGLE_CLIENT_SECRET", "client-secret")
	t.Setenv("GOOGLE_REDIRECT_URL", "http://localhost/callback")
	t.Setenv("GCP_PROJECT_ID", "my-project")
	t.Setenv("JWT_SECRET", "secret-key")

	_, err := LoadConfig()
	if err == nil {
		t.Fatal("expected error when GOOGLE_CLIENT_ID is missing")
	}
}

func TestLoadConfigMissingGoogleRedirectURL(t *testing.T) {
	t.Setenv("GOOGLE_CLIENT_ID", "client-id")
	t.Setenv("GOOGLE_CLIENT_SECRET", "client-secret")
	t.Setenv("GCP_PROJECT_ID", "my-project")
	t.Setenv("JWT_SECRET", "secret-key")

	_, err := LoadConfig()
	if err == nil {
		t.Fatal("expected error when GOOGLE_REDIRECT_URL is missing")
	}
}

func TestLoadConfigMissingGCPProjectID(t *testing.T) {
	t.Setenv("GOOGLE_CLIENT_ID", "client-id")
	t.Setenv("GOOGLE_CLIENT_SECRET", "client-secret")
	t.Setenv("GOOGLE_REDIRECT_URL", "http://localhost/callback")
	t.Setenv("JWT_SECRET", "secret-key")

	_, err := LoadConfig()
	if err == nil {
		t.Fatal("expected error when GCP_PROJECT_ID is missing")
	}
}

func TestLoadConfigMissingJWTSecret(t *testing.T) {
	t.Setenv("GOOGLE_CLIENT_ID", "client-id")
	t.Setenv("GOOGLE_CLIENT_SECRET", "client-secret")
	t.Setenv("GOOGLE_REDIRECT_URL", "http://localhost/callback")
	t.Setenv("GCP_PROJECT_ID", "my-project")

	_, err := LoadConfig()
	if err == nil {
		t.Fatal("expected error when JWT_SECRET is missing")
	}
}

func TestLoadConfigMissingMultipleRequiredVars(t *testing.T) {
	t.Setenv("GOOGLE_CLIENT_SECRET", "client-secret")

	_, err := LoadConfig()
	if err == nil {
		t.Fatal("expected error when multiple required env vars are missing")
	}
}

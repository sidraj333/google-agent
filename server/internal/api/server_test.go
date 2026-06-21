package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

func TestServerHealthEndpoint(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/handleReq", HandleReq)
	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	resp, err := http.Get(testServer.URL + "/health")
	if err != nil {
		t.Fatalf("failiure calling health endpoint: %v", err)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", resp.Header.Get("Content-Type"))
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status: %v but got status: %v", http.StatusOK, resp.StatusCode)
	}
	var respBodyJSON map[string]string
	err = json.NewDecoder(resp.Body).Decode(&respBodyJSON)
	if err != nil {
		t.Fatalf("failiure converting response body to json: %v", err)
	}
	if respBodyJSON["status"] != "ok" {
		t.Fatalf("expected status ok but got: %v", respBodyJSON["status"])
	}





}

func TestServerHandleReqEndpoint(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/handleReq", HandleReq)
	testServer := httptest.NewServer(mux)
	defer testServer.Close()
	resp, err := http.Get(testServer.URL + "/handleReq")
	if err != nil {
		t.Fatalf("failiure calling health endpoint: %v", err)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", resp.Header.Get("Content-Type"))
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status: %v but got status: %v", http.StatusOK, resp.StatusCode)
	}
	var respBodyJSON map[string]string
	err = json.NewDecoder(resp.Body).Decode(&respBodyJSON)
	if err != nil {
		t.Fatalf("failiure converting response body to json: %v", err)
	}
	if respBodyJSON["status"] != "ok" {
		t.Fatalf("expected status ok but got: %v", respBodyJSON["status"])
	}
}

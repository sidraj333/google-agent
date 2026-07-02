package agent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleReq(t *testing.T) {
	req := httptest.NewRequest("GET", "/handleReq", nil)
	w := httptest.NewRecorder()
	HandleReq(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status: %v but got status: %v", http.StatusOK, resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", resp.Header.Get("Content-Type"))
	}

	var responseBodyJSON map[string]string
	err := json.NewDecoder(resp.Body).Decode(&responseBodyJSON)

	if err != nil {
		t.Fatalf("error retrieving response body json: %v", err)
	}
	if responseBodyJSON["status"] != "ok" {
		t.Fatalf("expected status ok but got: %v", responseBodyJSON["status"])
	}
}

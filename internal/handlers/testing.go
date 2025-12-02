package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// creates a mock request for testing, returns data as map
func mockRequest(t *testing.T, method, path string, expectedStatus int, reqBody any, handler func(http.ResponseWriter, *http.Request)) map[string]any {
	var body []byte
	if reqBody != nil {
		body, _ = json.Marshal(reqBody)
	}

	req := httptest.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler(rr, req)

	// Check status code
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, expectedStatus)
	}

	var res SuccessResponse
	if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	data, ok := res.Data.(map[string]any)
	if !ok {
		t.Fatalf("Expected Data to be a map, got %T", res.Data)
	}

	return data
}

func mockRequestWithID(t *testing.T, method, path string, expectedStatus int, reqBody any, handler func(http.ResponseWriter, *http.Request, int), id int) map[string]any {
	var body []byte
	if reqBody != nil {
		body, _ = json.Marshal(reqBody)
	}

	req := httptest.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler(rr, req, id)

	// Check status code
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, expectedStatus)
	}

	var res SuccessResponse
	if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	data, ok := res.Data.(map[string]any)
	if !ok {
		t.Fatalf("Expected Data to be a map, got %T", res.Data)
	}

	return data
}

func assertField(t *testing.T, name string, expected any, data map[string]any) {
	expStr := fmt.Sprintf("%v", expected)
	dataStr := fmt.Sprintf("%v", data[name])
	if expStr != dataStr {
		t.Errorf("Expected %s %s, got %s", name, expStr, dataStr)
	}
}

package test

import (
	"bytes"
	"fmt"
	"kanggo/absenService/configs"
	v1 "kanggo/absenService/server/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostAttendance(t *testing.T) {
	// Set env
	configs.SetEnv()

	// Injecting Server method into a tests server
	h := v1.Host("8093")
	// call the server
	_, router, _ := h.Server()
	testServer := httptest.NewServer(router)
	// Shut down the server and block until all requests have gone through
	defer testServer.Close()

	// set body
	var body = []byte(`{
        "project_id": 5678,
		"worker_id": 1234,
		"image": "http://asd",
		"status": 2
    }`)

	// Make a request to our server with the {base url}/api/v1/attendance
	resp, err := http.Post(
		fmt.Sprintf("%s/api/v1/attendance", testServer.URL),
		"application/json" , bytes.NewBuffer(body))

	// if error
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// if response status not OK
	if resp.StatusCode != 200 && resp.StatusCode != 201{
		t.Fatalf("Expected status code 201/ 200, got %v", resp.StatusCode)
	}

	// set response header
	val, ok := resp.Header["Content-Type"]

	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json" {
		t.Fatalf( "Expected \"application/json\", got %s", val[0])
	}
}

package test

import (
	"fmt"
	"kanggo/absenService/configs"
	v1 "kanggo/absenService/server/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAttendance(t *testing.T) {
	// Set env
	configs.SetEnv()

	// Injecting Server method into a tests server
	h := v1.Host("8093")
	// call the server
	_, router, _ := h.Server()
	testServer := httptest.NewServer(router)
	// Shut down the server and block until all requests have gone through
	defer testServer.Close()

	resp, err := http.Get(
		fmt.Sprintf("%s/api/v1/attendance?project_id=2587&worker_id=269&date=2020-09-03",
			testServer.URL))

	// if error
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// if response status not OK
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
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

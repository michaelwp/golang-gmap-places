package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kanggo/absenService/configs"
	"kanggo/absenService/models"
	v1 "kanggo/absenService/server/v1"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetAttendance(t *testing.T) {
	// Set env
	configs.SetEnv()

	// define the model
	var respModel models.ResultPlaces

	// Injecting Server method into a tests server
	h := v1.Host("8093")
	// call the server
	_, router, _ := h.Server()
	testServer := httptest.NewServer(router)
	// Shut down the server and block until all requests have gone through
	defer testServer.Close()

	resp, err := http.Get(fmt.Sprintf("%s/api/v1/map?place=kanggo", testServer.URL))
	defer resp.Body.Close()

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

	// get string of the body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error read body response %v", err)
	}
	bodyString := string(bodyBytes)

	// convert to json
	err = json.Unmarshal([]byte(bodyString), &respModel)
	if err != nil {
		t.Fatalf("Error unmarshal body %v", err)
	}

	// Assert the response model
	respModelType := reflect.TypeOf(models.ResultPlaces{})
	dataType := reflect.TypeOf(respModel)
	if respModelType != dataType {
		t.Fatalf("Expected data type \"models.ResultPlaces\", got %v", dataType)
	}

	// Assert the response content
	if respModel.Code != 1 {
		t.Fatalf("Expected response code 1, got %v", respModel.Code)
	}
	if respModel.Message != "Places" {
		t.Fatalf("Expected response messages places, got %v", respModel.Message)
	}
	if respModel.Data == nil {
		t.Fatalf("Expected response data not nil, got %v", respModel.Data)
	}

	arrayPlaces := reflect.TypeOf([]models.Places{})
	dataType = reflect.TypeOf(respModel.Data)
	if arrayPlaces != dataType {
		t.Fatalf("Expected data type \"[]models.Places\", got %v", arrayPlaces)
	}
}

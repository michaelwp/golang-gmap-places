package v1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	v1 "kanggo/absenService/controllers/v1"
	"kanggo/absenService/errHandler"
	"net/http"
)

func RouterProfile(r *mux.Router) {
	// set attendance prefix
	maps := r.PathPrefix("/map").Subrouter()

	// attendance
	maps.HandleFunc("", v1.GetPlaces).Methods("GET")

	//PING
	maps.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]string{
			"Code":    "1",
			"Message": "PONG",
		}
		err := json.NewEncoder(w).Encode(response)
		errHandler.ErrHandler("Error json response: ", err)
	}).Methods("GET")
}

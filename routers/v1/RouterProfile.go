package v1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	v1 "kanggo/absenService/controllers/v1"
	"kanggo/absenService/errHandler"
	"kanggo/absenService/models"
	"net/http"
)

func RouterProfile(r *mux.Router) {
	// set attendance prefix
	attendance := r.PathPrefix("/attendance").Subrouter()

	// attendance
	attendance.HandleFunc("", v1.GetAttendance).Methods("GET")
	attendance.HandleFunc("", v1.PostAttendance).Methods("POST")

	// code
	attendance.HandleFunc("/code", v1.GenerateAttendanceCode).Methods("OPTIONS")
	attendance.HandleFunc("/code", v1.PostAttendanceCode).Methods("POST")

	//PING
	attendance.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := models.Result{
			Code:    1,
			Message: "PONG",
		}
		err := json.NewEncoder(w).Encode(response)
		errHandler.ErrHandler("Error json response: ", err)
	}).Methods("GET")
}

package v1

import (
	"encoding/json"
	"kanggo/absenService/errHandler"
	"kanggo/absenService/models"
	"net/http"
)

func GetPlaces(w http.ResponseWriter, r *http.Request)  {
	// set json response
	w.Header().Set("Content-type", "application/json")

	// get query params
	place := r.FormValue("place")

	if place == "" {
		errHandler.ErrorResponse(w, 0, http.StatusBadRequest, "Place required")
		return
	}

	w.WriteHeader(http.StatusOK)
	response := models.ResultPlaces{
		Code:    1,
		Message: place,
	}
	err := json.NewEncoder(w).Encode(response)
	errHandler.ErrHandler("Error json response: ", err)
}

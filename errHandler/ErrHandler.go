package errHandler

import (
	"encoding/json"
	"kanggo/absenService/models"
	"log"
	"net/http"
)

func ErrHandler(s string,e error) {
	if e != nil {
		log.Printf("%s : %s", s, e)
	}
}

func ErrorResponse(w http.ResponseWriter, code int, httpCode int, message string){
	w.WriteHeader(httpCode)
	response := models.Result{
		Code:    code,
		Message: message,
	}
	err := json.NewEncoder(w).Encode(response)
	ErrHandler("Error json response: ", err)
}
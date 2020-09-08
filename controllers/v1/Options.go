package v1

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"kanggo/absenService/errHandler"
	"kanggo/absenService/helpers"
	"kanggo/absenService/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GenerateAttendanceCode(w http.ResponseWriter, r *http.Request) {
	var res []models.AttendanceCode

	// set header response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	// set url query request
	projectId := r.FormValue("project_id")
	//check if project_id exist
	if projectId == "" {
		errHandler.ErrorResponse(w, 0, http.StatusBadRequest, "project_id is required")
		return
	}
	// convert project id to integer
	pId, _ := strconv.Atoi(projectId)

	//generate code
	code := helpers.GenerateCode()

	// check if code already registered
	c := make(chan []models.AttendanceCode)
	go checkProject(pId, c)
	res = <-c

	if len(res) > 0 {
		// update current attendance code
		res = UpdateNewCode(code, pId)
	} else {
		// add new attendance code
		res = AddNewCode(code, pId)
	}

	// set api response
	w.WriteHeader(http.StatusOK)
	response := models.ResultSingleAttendanceCode{
		Code:    1,
		Message: "Attendance Code",
		Data: res[0],
	}

	// convert to json
	err := json.NewEncoder(w).Encode(response)
	errHandler.ErrHandler("Error json response: ", err)
}

func checkProject(projectId int, c chan []models.AttendanceCode) {
	var results []models.AttendanceCode

	// set filter
	filter := bson.M{
		"projectId": projectId,
		"isDeleted": false,
	}

	// get data from database
	cur, err := DbCon.Collection("code_attendance").Find(context.Background(), filter)
	errHandler.ErrHandler("Error get data Attendance: ", err)
	// looping data
	for cur.Next(context.Background()) {
		var s models.AttendanceCode

		// create a value into which the single document can be decoded
		err := cur.Decode(&s)
		errHandler.ErrHandler("Error decode data to attendance code model: ", err)

		// convert time to local time asia/jakarta - wib
		s.CreatedDate = helpers.DateTimeFormat(s.CreatedDate)
		s.UpdatedDate = helpers.DateTimeFormat(s.UpdatedDate)

		// append to array result
		results = append(results, s)
	}

	//set attendance code data to channel
	c <- results
}

/*
	ADD NEW ATTENDANCE CODE
*/
func AddNewCode(code string, projectId int) []models.AttendanceCode{
	// set input data
	inputData := bson.D{
		{"projectId", projectId},
		{"code", code},
		{"isDeleted", false},
		{"createdDate", time.Now()},
		{"updatedDate", time.Now()},
	}
	// input data
	res, err := DbCon.Collection("code_attendance").InsertOne(context.Background(), inputData)
	errHandler.ErrHandler("Error add new data: ", err)
	log.Printf("inserted document with ID %v\n\n", res.InsertedID)

	// get the data had been inserted
	c := make(chan []models.AttendanceCode)
	go checkProject(projectId, c)
	result := <-c

	return result
}

/*
	UPDATE ATTENDANCE CODE
*/
func UpdateNewCode(code string, projectId int) []models.AttendanceCode{
	// set input data
	updateData := bson.M{
		"$set": bson.M{
			"code":        code,
			"updatedDate": time.Now(),
		},
	}

	// filter data
	filter := bson.M{
		"projectId": projectId,
	}

	// input data
	res, err := DbCon.Collection("code_attendance").UpdateOne(context.Background(), filter, updateData)
	errHandler.ErrHandler("Error update new code: ", err)
	log.Printf("Upserted document with ID %v\n\n", res.UpsertedID)

	// get the data had been updated
	c := make(chan []models.AttendanceCode)
	go checkProject(projectId, c)
	result := <-c

	return result

}

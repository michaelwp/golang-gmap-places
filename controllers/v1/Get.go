package v1

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"kanggo/absenService/errHandler"
	"kanggo/absenService/helpers"
	"kanggo/absenService/models"
	"net/http"
	"strconv"
	"time"
)

/**
GET Attendance
URL: 	{base_url}/api/v1/attendance?project_id={project_id}&worker_id={worker_id}
QUERY: 	project_id -> string,
		worker_id -> string,
		date -> string
*/
func GetAttendance(w http.ResponseWriter, r *http.Request) {
	var d time.Time
	var err error

	// return api response
	w.Header().Set("Content-type", "application/json")

	// set url query request
	projectId := r.FormValue("project_id")
	workerId := r.FormValue("worker_id")
	date := r.FormValue("date")

	// create channel for data result
	c := make(chan []models.Attendance)

	// convert project & worker id to int32
	pId, _ := strconv.Atoi(projectId)
	wId, _ := strconv.Atoi(workerId)

	// convert to date
	if date != "" {
		d, err = time.Parse("2006-01-02", date)
		if err != nil {
			errHandler.ErrorResponse(w, 0, http.StatusBadRequest, "Error: " + err.Error())
			return
		}
	}

	// find attendance
	go FindAttendance(pId, wId, d, c)
	res := <-c

	// set api response
	w.WriteHeader(http.StatusOK)
	response := models.Result{
		Code:    1,
		Message: "Worker Attendance",
		Data:    res,
	}

	// convert to json
	err = json.NewEncoder(w).Encode(response)
	errHandler.ErrHandler("Error json response: ", err)
}

func FindAttendance(projectId int, workerId int, d time.Time, c chan []models.Attendance) {
	var results []models.Attendance

	// get filter model
	filter := setFilter(projectId, workerId, d)

	// get data from database
	cur, err := DbCon.Collection("worker_attendance").Find(context.Background(), filter)
	errHandler.ErrHandler("Error get data Attendance: ", err)

	// looping data
	for cur.Next(context.Background()) {
		var s models.Attendance

		// create a value into which the single document can be decoded
		err := cur.Decode(&s)
		errHandler.ErrHandler("Error decode data to attendance model: ", err)

		// get attendance code
		c := make(chan models.AttendanceCode)
		go FindAttendanceCode(int(s.ProjectId), c)
		resAttCode := <-c

		s.AttendanceCode = resAttCode

		// convert time to local time asia/jakarta - wib
		s.CreatedDate = helpers.DateTimeFormat(s.CreatedDate)
		s.UpdatedDate = helpers.DateTimeFormat(s.UpdatedDate)

		// append to array result
		results = append(results, s)
	}

	// set result to channel
	c <- results
}

func setFilter(projectId int, workerId int, d time.Time) bson.M {
	// date format
	dt, _ := helpers.DateFormat(d, "ISO")

	//set default filter
	filter := bson.M{
		"projectId": projectId,
		"workerId":  workerId,
		"date":      dt,
		"isDeleted": false,
	}

	// if all parameters value are empty
	if projectId == 0 && workerId == 0 && d.IsZero() {
		filter = bson.M{
			"isDeleted": false,
		}
	}
	// if projectId is not empty and the rest are empty
	if projectId != 0 && workerId == 0 && d.IsZero() {
		filter = bson.M{
			"projectId": projectId,
			"isDeleted": false,
		}
	}
	// if workerId is not empty and the rest are empty
	if projectId == 0 && workerId != 0 && d.IsZero() {
		filter = bson.M{
			"workerId":  workerId,
			"isDeleted": false,
		}
	}
	// if date is not empty and the rest are empty
	if projectId == 0 && workerId == 0 && !d.IsZero() {
		filter = bson.M{
			"date":      dt,
			"isDeleted": false,
		}
	}
	// if projectId and workerId is not empty and the rest are empty
	if projectId != 0 && workerId != 0 && d.IsZero() {
		filter = bson.M{
			"projectId": projectId,
			"workerId":  workerId,
			"isDeleted": false,
		}
	}
	// if projectId and date is not empty and the rest are empty
	if projectId != 0 && workerId == 0 && !d.IsZero() {
		filter = bson.M{
			"projectId": projectId,
			"date":      dt,
			"isDeleted": false,
		}
	}
	// if workerId and date is not empty and the rest are empty
	if projectId == 0 && workerId != 0 && !d.IsZero() {
		filter = bson.M{
			"workerId":  workerId,
			"date":      dt,
			"isDeleted": false,
		}
	}

	return filter
}

func FindAttendanceCode(
	projectId int,
	c chan models.AttendanceCode) {
	var results models.AttendanceCode

	// get filter model
	filter := bson.M{
		"projectId": projectId,
	}

	// get data from database
	err := DbCon.Collection("code_attendance").FindOne(context.Background(), filter).Decode(&results)
	errHandler.ErrHandler("Error get Attendance Code: ", err)

	// convert time to local time asia/jakarta - wib
	results.CreatedDate = helpers.DateTimeFormat(results.CreatedDate)
	results.UpdatedDate = helpers.DateTimeFormat(results.UpdatedDate)

	// set result to channel
	c <- results
}

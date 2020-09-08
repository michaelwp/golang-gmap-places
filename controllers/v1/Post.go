package v1

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"kanggo/absenService/db/v1"
	"kanggo/absenService/errHandler"
	"kanggo/absenService/helpers"
	"kanggo/absenService/models"
	"log"
	"net/http"
	"time"
)

var DbCon, _ = db.DbCon("kanggo_attendance")

/**
POST ATTENDANCE
URL: {base_url}/api/v1/attendance
BODY: {
	"project_id": 1234,
	"worker_id": 321,
	"image": "http://imageurl",
	"status": 2
}
*/
func PostAttendance(w http.ResponseWriter, r *http.Request) {
	var body models.Attendance
	var res []models.Attendance

	// return api response
	w.Header().Set("Content-type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errHandler.ErrorResponse(w, 0, http.StatusBadRequest, "Body error")
		return
	}

	// check if data already registered
	c := make(chan []models.Attendance)
	go FindAttendance(int(body.ProjectId), int(body.WorkerId), time.Now(), c)
	res = <-c

	// if it has been registered
	if len(res) > 0 {
		if res[0].Status > 0 {
			// if data already registered then stop here
			errHandler.ErrorResponse(w, 1, http.StatusOK, "Data already registered")
			return
		} else {
			// update status
			go UpdateAttendanceStatus(int(body.WorkerId), int(body.ProjectId), int(body.Status))
		}
	} else {
		// save data
		attChan := make(chan []models.Attendance)
		go addAttendance(body, attChan)
		res = <-attChan
	}

	// return api response
	w.WriteHeader(http.StatusCreated)
	response := models.ResultSingle{
		Code:    1,
		Message: "Data saved",
		Data:    res[0],
	}
	err = json.NewEncoder(w).Encode(response)
	errHandler.ErrHandler("Error json response: ", err)
}

/*
	ADD NEW ATTENDANCE DATA
*/
func addAttendance(body models.Attendance, r chan []models.Attendance) {
	// date format
	d, t := helpers.DateFormat(time.Now(), "ISO")

	// set input data
	inputData := bson.D{
		{"projectId", body.ProjectId},
		{"workerId", body.WorkerId},
		{"image", body.Image},
		{"status", body.Status},
		{"lat", body.Lat},
		{"lon", body.Lon},
		{"date", d},
		{"time", t},
		{"isDeleted", false},
		{"createdDate", time.Now()},
		{"updatedDate", time.Now()},
	}
	// input data
	res, err := DbCon.Collection("worker_attendance").InsertOne(context.Background(), inputData)
	errHandler.ErrHandler("Error add new data: ", err)
	log.Printf("inserted document with ID %v\n\n", res.InsertedID)

	c := make(chan []models.Attendance)
	go FindAttendance(int(body.ProjectId), int(body.WorkerId), time.Now(), c)
	result := <-c

	r <- result
}

/*
	UPDATE ATTENDANCE STATUS
*/
func UpdateAttendanceStatus(workerId int, projectId int, status int) ([]models.Attendance, error) {
	var res []models.Attendance

	// set input data
	updateData := bson.M{
		"$set": bson.M{
			"status":      status,
			"updatedDate": time.Now(),
		},
	}

	// filter data
	filter := bson.M{
		"projectId": projectId,
		"workerId":  workerId,
	}

	// input data
	_, err := DbCon.Collection("worker_attendance").UpdateOne(context.Background(), filter, updateData)
	log.Printf("Upserted document with projectID %d and workerID %d", projectId, workerId)

	// find attendance
	c := make(chan []models.Attendance)
	go FindAttendance(projectId, workerId, time.Now(), c)
	res = <-c

	return res, err
}

/*
	POST ATTENDANCE BY CODE
*/
func PostAttendanceCode(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Code      string  `json:"code"`
		WorkerId  int32   `json:"worker_id"`
		ProjectId int32   `json:"project_id"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
	}

	// default response
	response := models.ResultSingle{
		Code:    1,
		Message: "Worker successfully absent",
	}

	//var res []models.Attendance

	// return api response
	w.Header().Set("Content-type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errHandler.ErrorResponse(w, 0, http.StatusBadRequest, "Body error")
		return
	}

	// match attendance code
	_, err = MatchAttendanceCode(body.Code, body.ProjectId)
	if err != nil {
		errHandler.ErrorResponse(w, 1, http.StatusOK, "Code not found")
		return
	}

	// find attendance
	c := make(chan []models.Attendance)
	go FindAttendance(int(body.ProjectId), int(body.WorkerId), time.Now(), c)
	res := <-c

	if len(res) == 0 {
		// if attendance not found create new data attendance
		attChan := make(chan []models.Attendance)
		go addAttendance(models.Attendance{
			ProjectId: body.ProjectId,
			WorkerId:  body.WorkerId,
			Image:     "",
			Status:    1,
			Lat:       body.Lat,
			Lon:       body.Lon,
		}, attChan)
		res = <-attChan
	} else {
		if res[0].Status == 1 || res[0].Status == 2{
			// if worker already absent
			response.Message = "Worker already absent"
		} else {
			// update attendance data
			res, err = UpdateAttendanceStatus(int(body.WorkerId), int(body.ProjectId), 1)
			errHandler.ErrHandler("Error update status: ", err)
		}
	}

	w.WriteHeader(http.StatusOK)
	response.Data = res[0]

	err = json.NewEncoder(w).Encode(response)
	errHandler.ErrHandler("Error json response: ", err)
}

func MatchAttendanceCode(code string, projectId int32) (models.AttendanceCode, error) {
	var results models.AttendanceCode

	// get filter model
	filter := bson.M{
		"code": code,
		"projectId": projectId,
	}

	// get data from database
	err := DbCon.Collection("code_attendance").FindOne(context.Background(), filter).Decode(&results)

	// convert time to local time asia/jakarta - wib
	results.CreatedDate = helpers.DateTimeFormat(results.CreatedDate)
	results.UpdatedDate = helpers.DateTimeFormat(results.UpdatedDate)

	return results, err
}

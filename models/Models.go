package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Attendance struct {
	Id             primitive.ObjectID `bson:"_id" json:"id"`
	ProjectId      int32              `json:"project_id"`
	WorkerId       int32              `json:"worker_id"`
	Image          string             `json:"image"`
	Status         int32              `json:"status"`
	Lat            float64            `json:"lat"`
	Lon            float64            `json:"lon"`
	Date           string             `json:"date"`
	Time           string             `json:"time"`
	AttendanceCode AttendanceCode     `json:"attendance_code"`
	IsDeleted      bool               `json:"is_deleted"`
	CreatedDate    time.Time          `json:"created_date"`
	UpdatedDate    time.Time          `json:"updated_date"`
}

type AttendanceCode struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	ProjectId   int32              `json:"project_id"`
	Code        string             `json:"code"`
	IsDeleted   bool               `json:"is_deleted"`
	CreatedDate time.Time          `json:"created_date"`
	UpdatedDate time.Time          `json:"updated_date"`
}

type ResultSingle struct {
	Code    int32      `json:"code"`
	Message string     `json:"message"`
	Data    Attendance `json:"data"`
}

type Result struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []Attendance `json:"data"`
}

type ResultSingleAttendanceCode struct {
	Code    int32          `json:"code"`
	Message string         `json:"message"`
	Data    AttendanceCode `json:"data"`
}

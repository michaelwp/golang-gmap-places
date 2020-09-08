package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Keywords struct {
	Id      primitive.ObjectID `bson:"_id" json:"id"`
	Keyword string             `json:"keyword"`
	Places  []Places           `json:"places"`
}

type Places struct {
	Id    primitive.ObjectID `bson:"_id" json:"id"`
	Place string             `json:"place"`
	Lat   float64            `json:"lat"`
	Lon   float64            `json:"lon"`
}

type ResultKeywords struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Keywords `json:"data"`
}

type ResultPlaces struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

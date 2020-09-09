package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"googlemaps.github.io/maps"
	"kanggo/absenService/db/v1"
	"kanggo/absenService/errHandler"
	"kanggo/absenService/helpers"
	"kanggo/absenService/models"
	"net/http"
)

// call mongodb
var mongoDb, _, _ = db.DbCon("map_places")

func GetPlaces(w http.ResponseWriter, r *http.Request) {
	// define places result models
	var placesArray []models.Places
	var placesSingle models.Places

	// define place search response from google map
	var places maps.PlacesSearchResponse

	// define error
	var err error

	// set json response
	w.Header().Set("Content-type", "application/json")

	// get query params
	place := r.FormValue("place")

	if place == "" {
		errHandler.ErrorResponse(w, 0, http.StatusBadRequest, "Place required")
		return
	}

	// find the keyword
	c := make(chan []models.Keywords)
	go FindKeyWord(place, c)
	cRes := <-c

	if len(cRes) == 0 {
		places, err = helpers.GmapsPlace(place)
		if err != nil {
			errHandler.ErrorResponse(w, 0, http.StatusInternalServerError, err.Error())
			return
		}

		// save keyword if data not exist
		go SaveKeyword(place)

		for _, res := range places.Results {
			// set places item
			placesSingle.Keyword = place
			placesSingle.PlaceId = res.PlaceID
			placesSingle.Name = res.Name
			placesSingle.Address = res.FormattedAddress
			placesSingle.Lat = res.Geometry.Location.Lat
			placesSingle.Lon = res.Geometry.Location.Lng

			// save keyword if data not exist
			go SavePlaces(placesSingle)

			// append to array places
			placesArray = append(placesArray, placesSingle)
		}
	} else {
		cPlace := make(chan []models.Places)
		go FindPlaces(place, cPlace)
		placesArray = <- cPlace
	}

	w.WriteHeader(http.StatusOK)
	response := models.ResultPlaces{
		Code:    1,
		Message: "Places",
		Data:    placesArray,
	}
	err = json.NewEncoder(w).Encode(response)
	errHandler.ErrHandler("Error json response: ", err)
}

/*
	FIND KEYWORD
*/
func FindKeyWord(keyword string, c chan []models.Keywords) {
	// Here's an array in which you can store the decoded documents
	var results []models.Keywords

	filter := bson.M{
		"keyword": keyword,
	}

	// Passing bson.M{} as the filter matches all documents in the collection
	cur, err := mongoDb.Collection("keywords").Find(context.Background(), filter)
	errHandler.ErrHandler("Error finding keyword: ", err)

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem models.Keywords
		err = cur.Decode(&elem)
		errHandler.ErrHandler("Error decode data: ", err)

		results = append(results, elem)
	}

	err = cur.Err()
	errHandler.ErrHandler("Error cursor: ", err)

	// Close the cursor once finished
	err = cur.Close(context.Background())
	errHandler.ErrHandler("Error close cursor: ", err)

	c <- results
}

/*
	SAVE KEYWORD
*/
func SaveKeyword(keyword string) {
	keywordData := map[string]string{
		"keyword": keyword,
	}

	// save data
	insertPlace, err := mongoDb.Collection("keywords").InsertOne(context.Background(), keywordData)
	errHandler.ErrHandler("Error save data: ", err)

	// print status
	status := fmt.Sprintf("Inserted multiple documents: %v", insertPlace.InsertedID)
	fmt.Println(status)
}

func SavePlaces(places models.Places) {
	// save data
	insertPlace, err := mongoDb.Collection("places").InsertOne(context.Background(), places)
	errHandler.ErrHandler("Error save data: ", err)

	// print status
	status := fmt.Sprintf("Inserted multiple documents: %v", insertPlace.InsertedID)
	fmt.Println(status)
}

/*
	FIND PLACES
*/
func FindPlaces(keyword string, c chan []models.Places) {
	// Here's an array in which you can store the decoded documents
	var results []models.Places

	filter := bson.M{
		"keyword": keyword,
	}

	// Passing bson.M{} as the filter matches all documents in the collection
	cur, err := mongoDb.Collection("places").Find(context.Background(), filter)
	errHandler.ErrHandler("Error finding places: ", err)

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem models.Places
		err = cur.Decode(&elem)
		errHandler.ErrHandler("Error decode data: ", err)

		results = append(results, elem)
	}

	err = cur.Err()
	errHandler.ErrHandler("Error cursor: ", err)

	// Close the cursor once finished
	err = cur.Close(context.Background())
	errHandler.ErrHandler("Error close cursor: ", err)

	c <- results
}

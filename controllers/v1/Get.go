package v1

import (
	"context"
	"encoding/json"
	"github.com/michaelwp/golang-gmap-places/configs"
	"github.com/michaelwp/golang-gmap-places/db/v1"
	"github.com/michaelwp/golang-gmap-places/errHandler"
	"github.com/michaelwp/golang-gmap-places/helpers"
	"github.com/michaelwp/golang-gmap-places/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

// call mongodb
var mongoDb, _, _ = db.DbCon("map_places")

func GetPlaces(w http.ResponseWriter, r *http.Request) {
	//define places result models
	var placesArray []models.Places

	//define error
	var (err error)

	// set json response
	w.Header().Set("Content-type", "application/json")

	// get query params
	place := strings.ToLower(r.FormValue("place"))
	cid := strings.ToLower(r.FormValue("country"))

	if place == "" {
		errHandler.ErrorResponse(
			w, configs.Code("ERROR"),
			http.StatusBadRequest,
			configs.Message("PLACE_REQUIRED"))
		return
	}

	// find the keyword
	//go FindKeyWord(place, cid, c)
	keywords, errChan := FindKeyWordAsync(r.Context(), place, cid)
	if <-errChan != nil {
		errHandler.ErrorResponse(
			w, configs.Code("ERROR"),
			http.StatusBadRequest,
			"error while find keywords")
		return
	}
	if len(<-keywords) == 0 {
		// get places list from google map
		placesArray, err = helpers.GmapsAutoComplete(place, cid)
		if err != nil {
			errHandler.ErrorResponse(
				w, configs.Code("ERROR"),
				http.StatusInternalServerError,
				err.Error())
			return
		}
	} else {
		cPlace := make(chan []models.Places)
		go FindPlaces(r.Context(), place, cid, cPlace)
		placesArray = <-cPlace
	}

	if placesArray == nil {
		errHandler.ErrorResponse(
			w, configs.Code("NOT FOUND"),
			http.StatusOK,
			configs.Message("PLACES_NOT_FOUND"))
		return
	}

	w.WriteHeader(http.StatusOK)
	response := models.ResultPlaces{
		Code:    configs.Code("SUCCESS"),
		Message: configs.Message("PLACES"),
		Data:    placesArray,
	}
	err = json.NewEncoder(w).Encode(response)
	errHandler.ErrHandler("Error json response: ", err)
}

/*
	FIND KEYWORD
*/
func FindKeyWord(ctx context.Context, keyword string, cid string) (keywords []models.Keywords, err error) {
	// Here's an array in which you can store the decoded documents
	filter := bson.M{
		"keyword": keyword,
		"country": cid,
	}

	// Passing bson.M{} as the filter matches all documents in the collection
	cur, err := mongoDb.Collection("keywords").Find(ctx, filter)
	if err != nil {
		errHandler.ErrHandler("Error finding keyword: ", err)
		return keywords, err
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem models.Keywords
		err = cur.Decode(&elem)
		if err != nil {
			errHandler.ErrHandler("Error decode data: ", err)
			return keywords, err
		}
		keywords = append(keywords, elem)
	}

	err = cur.Err()
	if err != nil {
		errHandler.ErrHandler("Error cursor: ", err)
		return keywords, err
	}


	// Close the cursor once finished
	err = cur.Close(ctx)
	if err != nil {
		errHandler.ErrHandler("Error close cursor: ", err)
		return keywords, err
	}
	return
}

func FindKeyWordAsync(ctx context.Context, keyword string, cid string) (<-chan[]models.Keywords, <-chan error){
	dataChan := make(chan []models.Keywords, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(dataChan)
		defer close(errChan)
		keywords, err := FindKeyWord(ctx, keyword, cid)
		if err != nil {
			errChan <- err
		}
		dataChan <- keywords
	}()
	return dataChan, errChan
}

/*
	FIND PLACES
*/
func FindPlaces(ctx context.Context, keyword string, cid string, c chan []models.Places) {
	// Here's an array in which you can store the decoded documents
	var results []models.Places

	filter := bson.M{
		"keyword": keyword,
		"country": cid,
	}

	// Passing bson.M{} as the filter matches all documents in the collection
	cur, err := mongoDb.Collection("places").Find(context.Background(), filter)
	errHandler.ErrHandler("Error finding places: ", err)

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var elem models.Places
		err = cur.Decode(&elem)
		errHandler.ErrHandler("Error decode data: ", err)

		results = append(results, elem)
	}

	err = cur.Err()
	errHandler.ErrHandler("Error cursor: ", err)

	// Close the cursor once finished
	err = cur.Close(ctx)
	errHandler.ErrHandler("Error close cursor: ", err)

	c <- results
}

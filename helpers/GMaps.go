package helpers

import (
	"context"
	"fmt"
	"github.com/michaelwp/golang-gmap-places/db/v1"
	"github.com/michaelwp/golang-gmap-places/errHandler"
	"github.com/michaelwp/golang-gmap-places/models"
	"googlemaps.github.io/maps"
	"os"
)

var loc maps.LatLng
var GoogleApiKey = os.Getenv("GOOGLE_API_KEY")
var mongoDb, _, _ = db.DbCon("map_places")

func Gmapsclient() (*maps.Client, error) {
	// set gmap client api key
	k := maps.WithAPIKey(GoogleApiKey)
	c, err := maps.NewClient(k)
	return c, err
}

/*
	GOOGLE MAP PLACE SEARCH
*/
func GmapsPlace(s string) (maps.PlacesSearchResponse, error) {
	// set gmap client api key
	c, err := Gmapsclient()

	// set lat lng for location
	loc.Lat = -6.264461
	loc.Lng = 106.689820
	l := &loc

	// setup option
	p := &maps.TextSearchRequest{
		Query:    s,
		Location: l,
		Radius:   200000,
		//Language:  "",
		//MinPrice:  "",
		//MaxPrice:  "",
		//OpenNow:   false,
		//Type:      "",
		//PageToken: "",
		Region: "ID",
	}

	// get the result
	places, err := c.TextSearch(context.Background(), p)

	return places, err
}

/*
	GOOGLE MAP PLACE DETAILS
*/
func GmapsPlaceDetails(s string) (maps.PlaceDetailsResult, error) {
	// set gmap client
	c, err := Gmapsclient()

	// setup field
	f := []maps.PlaceDetailsFieldMask{
		"formatted_address",
		"geometry",
	}

	// setup option
	p := &maps.PlaceDetailsRequest{
		PlaceID: s,
		//Language:     "",
		Fields: f,
		//SessionToken: maps.PlaceAutocompleteSessionToken{},
		Region: "ID",
	}

	places, err := c.PlaceDetails(context.Background(), p)

	return places, err
}

/*
	GOOGLE MAP AUTO COMPLETE
*/
func GmapsAutoComplete(s string, cid string) ([]models.Places, error) {
	// set gmap client
	c, err := Gmapsclient()

	// set lat lng for location
	//loc.Lat = -6.264461
	//loc.Lng = 106.689820
	//l := &loc

	var comp map[maps.Component][]string

	if cid != "" {
		// limit the result to region indonesia only
		country := []string{cid}
		comp = map[maps.Component][]string{
			"country": country,
		}
	} else {
		comp = nil
	}

	// setup option
	p := &maps.PlaceAutocompleteRequest{
		Input: s,
		//Offset:       3,
		//Location: l,
		//Origin:       nil,
		//Radius: 200000,
		//Language:     "",
		//Types:        "",
		Components: comp,
		//StrictBounds: false,
		//SessionToken: maps.PlaceAutocompleteSessionToken{},
	}

	// get the result
	places, err := c.PlaceAutocomplete(context.Background(), p)

	var placesArray []models.Places
	var placesSingle models.Places

	// limit the result to only 5 predictions result
	//if len(places.Predictions) > 5 {
	//	places.Predictions = places.Predictions[:5]
	//}

	for _, res := range places.Predictions {
		placeDetail, _ := GmapsPlaceDetails(res.PlaceID)

		placesSingle.Keyword = s
		placesSingle.PlaceId = res.PlaceID
		placesSingle.Name = res.StructuredFormatting.MainText
		placesSingle.Address = placeDetail.FormattedAddress
		placesSingle.Country = cid
		placesSingle.Lat = placeDetail.Geometry.Location.Lat
		placesSingle.Lon = placeDetail.Geometry.Location.Lng

		// save keyword if data not exist
		go SavePlaces(placesSingle)

		// append to array places
		placesArray = append(placesArray, placesSingle)
	}

	go func() {
		if placesArray != nil {
			// save keyword if data not exist
			SaveKeyword(s, cid)
		}
	}()

	return placesArray, err
}

/*
	SAVE KEYWORD
*/
func SaveKeyword(keyword string, cid string) {
	keywordData := map[string]string{
		"keyword": keyword,
		"country": cid,
	}

	// save data
	insertPlace, err := mongoDb.Collection("keywords").InsertOne(context.Background(), keywordData)
	errHandler.ErrHandler("Error save data: ", err)

	// print status
	status := fmt.Sprintf("Inserted multiple documents: %v", insertPlace.InsertedID)
	fmt.Println(status)
}

/*
	SAVE PLACES
*/

func SavePlaces(places models.Places) {
	// save data
	insertPlace, err := mongoDb.Collection("places").InsertOne(context.Background(), places)
	errHandler.ErrHandler("Error save data: ", err)

	// print status
	status := fmt.Sprintf("Inserted multiple documents: %v", insertPlace.InsertedID)
	fmt.Println(status)
}

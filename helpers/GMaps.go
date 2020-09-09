package helpers

import (
	"context"
	"googlemaps.github.io/maps"
	"os"
)

func GmapsPlace(s string) (maps.PlacesSearchResponse, error) {
	// set gmap client api key
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))

	// set request
	p := &maps.TextSearchRequest{
		Query: s,
	}

	// get the result
	places, err := c.TextSearch(context.Background(), p)

	return places, err
}

package helpers

import (
	"context"
	"googlemaps.github.io/maps"
)

func GmapsPlace(s string) (maps.PlacesSearchResponse, error) {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyCBCNOKmd3VrxQaaPqxgNuGgfuZ1Idjryg"))

	p := &maps.TextSearchRequest{
		Query: s,
	}
	places, err := c.TextSearch(context.Background(), p)

	return places, err
}

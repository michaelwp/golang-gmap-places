package models

type Keywords struct {
	Keyword string   `json:"keyword"`
	Places  []Places `json:"places"`
}

type Places struct {
	Keyword string  `json:"keyword"`
	PlaceId string  `json:"place_id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type ResultKeywords struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Keywords `json:"data"`
}

type ResultPlaces struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []Places `json:"data"`
}

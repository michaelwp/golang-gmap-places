package configs

func Message(s string) (message string) {
	m := map[string]string{
		"PLACE_REQUIRED":   "place is required",
		"PLACES":           "places",
		"PLACES_NOT_FOUND": "places not found",
	}
	return m[s]
}

func Code(s string) (code int) {
	c := map[string]int{
		"ERROR":     0,
		"SUCCESS":   1,
		"DUPLICATE": 2,
		"NOT FOUND": 3,
	}
	return c[s]
}

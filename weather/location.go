package weather

import (
	"encoding/json"
)

type LocationService service

type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
}

func (s *LocationService) FetchLatLonForCity() (Location, error) {
	var err error
	var locations []Location
	var location Location

	url := locationURL + s.client.apiKey + "&q=" + s.client.city

	res, err := s.client.NewRequest(url)
	json.Unmarshal(res, &locations)

	if err != nil {
		return location, err
	}

	locationLength := len(locations)
	if locationLength == 0 {
		return location, errorNoCityFound
	} else {
		location = locations[0]
	}
	return location, err
}

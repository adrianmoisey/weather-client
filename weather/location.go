package weather

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
}

func (s *WeatherClient) FetchLatLonForCity(city string) (*Location, error) {
	var locations []Location

	url := locationURL + s.apiKey + "&q=" + city

	// Resty
	res, err := s.Fetch(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &locations)
	if err != nil {
		return nil, err
	}

	if len(locations) == 0 {
		cause := errors.New(errorNoCityFound)
		return nil, errors.WithStack(cause)
	}

	return &locations[0], nil
}

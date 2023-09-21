package weather

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type LocationService service

type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
}

func (s *LocationService) FetchLatLonForCity(city string) (*Location, error) {
	var locations []Location

	url := locationURL + s.client.apiKey + "&q=" + city

	res, err := s.client.NewRequest(url)
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

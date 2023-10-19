package weather

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Weather struct {
	WeatherConditions WeatherConditions `json:"main"`
}

type WeatherConditions struct {
	Temperature float64 `json:"temp"`
	Pressure    int32   `json:"pressure"`
}

func (s *WeatherClient) FetchWeatherForCity(city string) (*WeatherConditions, error) {
	var weather Weather

	location, err := s.FetchLatLonForCity(city)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	units := fmt.Sprintf("units=%s", s.unit)
	latlonURL := fmt.Sprintf("&lat=%v&lon=%v", location.Latitude, location.Longitude)

	url := weatherURL + units + latlonURL

	resp, err := s.Fetch(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = json.Unmarshal(resp, &weather)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &weather.WeatherConditions, nil
}

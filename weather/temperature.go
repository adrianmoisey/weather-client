package weather

import (
	"encoding/json"
	"fmt"
)

type Weather struct {
	Main WeatherConditions `json:"main"`
}

type WeatherConditions struct {
	Temperature float64 `json:"temp"`
	Pressure    int32   `json:"pressure"`
}

func (s *WeatherClient) FetchWeatherForCity(city string) (*WeatherConditions, error) {
	var weather Weather

	// Fetch the Latitide and Longitide for city
	location, err := s.FetchLatLonForCity(city)
	if err != nil {
		return nil, err
	}

	units := fmt.Sprintf("units=%s", s.unit)
	latlonURL := fmt.Sprintf("&lat=%v&lon=%v", location.Latitude, location.Longitude)

	url := weatherURL + units + latlonURL

	res, err := s.Fetch(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &weather)
	if err != nil {
		return nil, err
	}

	// This weather.Main stuff is weird, is there a better way to handle it?
	main := &weather.Main

	return main, nil
}

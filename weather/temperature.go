package weather

import (
	"encoding/json"
	"fmt"
)

type TemperatureService service

type Weather struct {
	Main WeatherConditions `json:"main"`
}

type WeatherConditions struct {
	Temperature float64 `json:"temp"`
	Pressure    int32   `json:"pressure"`
}

func (s *TemperatureService) FetchWeatherForCity(city string) (*WeatherConditions, error) {
	var weather Weather

	// Fetch the Latitide and Longitide for city
	location, err := s.client.Location.FetchLatLonForCity(city)
	if err != nil {
		return nil, err
	}

	latlonURL := fmt.Sprintf("&lat=%v&lon=%v", location.Latitude, location.Longitude)
	units := fmt.Sprintf("&units=%s", s.client.unit)
	url := weatherURL + s.client.apiKey + units + latlonURL

	res, err := s.client.NewRequest(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &weather)
	if err != nil {
		return nil, err
	}

	// This weather.Main stuff is weird, is there a better way to handle it?
	main := &weather.Main

	return main, err
}

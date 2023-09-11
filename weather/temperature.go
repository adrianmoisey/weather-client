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

func (s *TemperatureService) FetchWeatherForCity(city string) (WeatherConditions, error) {
	var err error
	var weather Weather
	var temperature WeatherConditions

	// Fetch the Latitide and Longitide for city
	location, err := s.client.Location.FetchLatLonForCity(city)
	if err != nil {
		return temperature, err
	}

	latlonURL := fmt.Sprintf("&lat=%v&lon=%v", location.Latitude, location.Longitude)
	units := fmt.Sprintf("&units=%s", s.client.unit)
	url := weatherURL + s.client.apiKey + units + latlonURL

	res, err := s.client.NewRequest(url)
	json.Unmarshal(res, &weather)

	return weather.Main, err
}

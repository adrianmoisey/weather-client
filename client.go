package weather

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const APIKey = "0ae0f3bac305cd9b45a4a324e0d3b88a"
const baseURL = "https://api.openweathermap.org/geo/1.0/direct?limit=1&appid=" + APIKey
const weatherURL = "https://api.openweathermap.org/data/2.5/weather?units=metric&appid=" + APIKey

type Client struct {
	restyClient *resty.Client
}

type GeoLocation struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
}

type Weather struct {
	Main Main `json:"main"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

func NewClient() *Client {
	client := resty.New()
	return &Client{restyClient: client}
}

// How do I make this private and also be testable
func (c *Client) fetchLatLonForCity(city string) (float64, float64, error) {
	var err error
	var locations []GeoLocation

	url := baseURL + "&q=" + city

	_, err = c.restyClient.R().SetResult(&locations).Get(url)
	if err != nil {
		return 0, 0, err
	}

	locationLength := len(locations)
	if locationLength == 0 {
		return 0, 0, ErrorNoCityFound
	}

	return locations[0].Latitude, locations[0].Longitude, err
}

func (c *Client) FetchWeatherForCity(city string) (float64, error) {
	var err error
	var temp Weather

	// Fetch the Latitide and Longitide for city
	lat, lon, err := c.fetchLatLonForCity(city)
	if err != nil {
		return 0, err
	}

	latlonURL := fmt.Sprintf("&lat=%v&lon=%v", lat, lon)
	url := weatherURL + latlonURL

	_, err = c.restyClient.R().SetResult(&temp).Get(url)
	if err != nil {
		return 0, err
	}
	return temp.Main.Temp, err
}

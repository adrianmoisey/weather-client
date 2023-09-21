package weather

import (
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	locationURL = "https://api.openweathermap.org/geo/1.0/direct?limit=1&appid="
	weatherURL  = "https://api.openweathermap.org/data/2.5/weather?appid="
)

type Client struct {
	apiKey string
	unit   string

	common service

	Temperature *TemperatureService
	Location    *LocationService
}

type service struct {
	client *Client
}

func NewClient(apiKey string, unit string) *Client {
	c := &Client{
		apiKey: apiKey,
		unit:   unit,
	}
	if c.unit == "" {
		c.unit = "metric"
	}
	c.initialize()
	return c
}

func (c *Client) initialize() {
	c.common.client = c

	c.Temperature = (*TemperatureService)(&c.common)
	c.Location = (*LocationService)(&c.common)
}

// TODO Switch to using resty
func (c *Client) NewRequest(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		cause := errors.New(invalidAPIKey)
		return nil, errors.WithStack(cause)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

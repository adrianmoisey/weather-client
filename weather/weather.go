package weather

import (
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	locationURL = "https://api.openweathermap.org/geo/1.0/direct?limit=1&appid="
	weatherURL  = "https://api.openweathermap.org/data/2.5/weather?appid="
)

type weatherClient struct {
	apiKey string
	unit   string

	httpClient *resty.Client
}

type WeatherConfig struct {
	apiKey string
	units  string "metric"
}

func NewClient(config WeatherConfig) (*weatherClient, error) {

	if config.apiKey == "" {
		cause := errors.New(apiKeyNotSupplied)
		return nil, errors.WithStack(cause)
	}

	c := &weatherClient{
		apiKey: config.apiKey,
		unit:   config.units,
	}
	c.httpClient = resty.New()
	return c, nil
}

func (c *weatherClient) Fetch(url string) ([]byte, error) {

	resp, err := c.httpClient.R().
		EnableTrace().
		Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == 401 {
		cause := errors.New(invalidAPIKey)
		return nil, errors.WithStack(cause)
	}

	return resp.Body(), err
}

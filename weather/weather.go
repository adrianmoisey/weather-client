package weather

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	locationURL = "https://api.openweathermap.org/geo/1.0/direct?limit=1&appid="
	weatherURL  = "https://api.openweathermap.org/data/2.5/weather?appid="
)

type WeatherClient struct {
	apiKey string
	unit   string

	httpClient *resty.Client
}

type WeatherConfig struct {
	ApiKey string
	Units  string "metric" // This isn't setting a default
}

func NewClient(config WeatherConfig) (*WeatherClient, error) {

	if config.ApiKey == "" {
		cause := errors.New(apiKeyNotSupplied)
		return nil, errors.WithStack(cause)
	}

	if config.Units == "" {
		config.Units = "metric"
	}

	c := &WeatherClient{
		apiKey: config.ApiKey,
		unit:   config.Units,
	}
	c.httpClient = resty.New()
	return c, nil
}

func (c *WeatherClient) Fetch(url string) ([]byte, error) {

	resp, err := c.httpClient.R().
		EnableTrace().
		Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return nil, errors.New(invalidAPIKey)
	}

	return resp.Body(), nil
}

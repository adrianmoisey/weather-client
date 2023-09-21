package weather

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestTemperature(t *testing.T) {
	test_city := "testville"
	api_key := "apikey"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", locationURL+api_key+"&q="+test_city,
		httpmock.NewStringResponder(200, `[{"lat": 0.1337, "lon": -1337.1, "country": "test-country", "state": "test-state"}]`))

	httpmock.RegisterResponder("GET", weatherURL+api_key+"&lat=0.1337&lon=-1337.1&units=metric",
		httpmock.NewStringResponder(200, `{"main":{"temp":13.37,"pressure":9999}}`))

	client := NewClient(api_key, "metric")
	temperature, err := client.Temperature.FetchWeatherForCity(test_city)

	assert.Equal(t, err, nil)

	assert.Equal(t, temperature.Pressure, int32(9999))
	assert.Equal(t, temperature.Temperature, 13.37)
}

func TestTemperatureFail(t *testing.T) {
	test_city := "testville"
	api_key := "apikey"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", locationURL+api_key+"&q="+test_city,
		httpmock.NewStringResponder(200, `[]`))

	client := NewClient(api_key, "metric")

	temperature, err := client.Temperature.FetchWeatherForCity(test_city)

	assert.EqualError(t, err, errorNoCityFound)

	assert.Equal(t, temperature.Pressure, int32(0))
	assert.Equal(t, temperature.Temperature, 0.0)
}

func TestTemperatureFailApiKey(t *testing.T) {
	test_city := "atlanta"
	api_key := "asdasd"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", locationURL+api_key+"&q="+test_city,
		httpmock.NewStringResponder(401, `{"cod":401, "message": "Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."}a`))

	client := NewClient(api_key, "metric")

	_, err := client.Temperature.FetchWeatherForCity(test_city)

	assert.EqualError(t, err, invalidAPIKey)
}

package weather

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestTemperature(t *testing.T) {

	httpmock.RegisterResponder("GET", locationURL+apiKey+"&q="+testCity,
		httpmock.NewStringResponder(200, `[{"lat": 0.1337, "lon": -1337.1, "country": "test-country", "state": "test-state"}]`))

	httpmock.RegisterResponder("GET", weatherURL+apiKey+"&lat=0.1337&lon=-1337.1&units=metric",
		httpmock.NewStringResponder(200, `{"main":{"temp":13.37,"pressure":9999}}`))

	temperature, err := testClient.FetchWeatherForCity(testCity)

	assert.NoError(t, err)

	assert.Equal(t, temperature.Pressure, int32(9999))
	assert.Equal(t, temperature.Temperature, 13.37)
}

func TestTemperatureFail(t *testing.T) {

	httpmock.RegisterResponder("GET", locationURL+apiKey+"&q="+testCity,
		httpmock.NewStringResponder(200, `[]`))

	temperature, err := testClient.FetchWeatherForCity(testCity)

	assert.Nil(t, temperature)
	assert.EqualError(t, err, errorNoCityFound)
}

func TestTemperatureFailApiKey(t *testing.T) {

	httpmock.RegisterResponder("GET", locationURL+apiKey+"&q="+testCity,
		httpmock.NewStringResponder(401, `{"cod":401, "message": "Invalid API key. Please see https://openweathermap.org/faq#error401 for more info."}`))

	city, err := testClient.FetchWeatherForCity(testCity)

	assert.Nil(t, city)
	assert.EqualError(t, err, invalidAPIKey)
}

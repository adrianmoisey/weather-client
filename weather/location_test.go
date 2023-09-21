package weather

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	test_city := "testville"
	api_key := "apikey"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", locationURL+api_key+"&q="+test_city,
		httpmock.NewStringResponder(200, `[{"lat": 0.1337, "lon": -1337.0, "country": "test-country", "state": "test-state"}]`))

	client := NewClient(api_key, "metric")
	location, err := client.Location.FetchLatLonForCity(test_city)

	assert.Equal(t, err, nil)

	assert.Equal(t, location.Latitude, 0.1337)
	assert.Equal(t, location.Longitude, -1337.0)
	assert.Equal(t, location.Country, "test-country")
	assert.Equal(t, location.State, "test-state")
}

func TestLocationFail(t *testing.T) {
	test_city := "testville"
	api_key := "apikey"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", locationURL+api_key+"&q="+test_city,
		httpmock.NewStringResponder(200, `[]`))

	client := NewClient(api_key, "metric")
	location, err := client.Location.FetchLatLonForCity(test_city)

	var emptyLocation Location

	assert.EqualError(t, err, errorNoCityFound)
	assert.Equal(t, location, emptyLocation)
}

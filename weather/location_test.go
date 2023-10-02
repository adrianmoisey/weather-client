package weather

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {

	httpmock.RegisterResponder("GET", locationURL+apiKey+"&q="+testCity,
		httpmock.NewStringResponder(200, `[{"lat": 0.1337, "lon": -1337.0, "country": "test-country", "state": "test-state"}]`))

	location, err := testClient.FetchLatLonForCity(testCity)

	assert.Equal(t, location.Latitude, 0.1337)
	assert.Equal(t, location.Longitude, -1337.0)
	assert.Equal(t, location.Country, "test-country")
	assert.Equal(t, location.State, "test-state")

	assert.NoError(t, err)
}

func TestLocationFail(t *testing.T) {

	httpmock.RegisterResponder("GET", locationURL+apiKey+"&q="+testCity,
		httpmock.NewStringResponder(200, `[]`))

	location, err := testClient.FetchLatLonForCity(testCity)

	assert.Nil(t, location)
	assert.EqualError(t, err, errorNoCityFound)
}

package weather

import (
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var (
	apiKey   string
	testCity string

	testClient *WeatherClient
)

func TestMain(m *testing.M) {
	apiKey = "apikey"
	testCity = "testville"

	config := WeatherConfig{
		ApiKey: apiKey,
		Units:  "metric",
	}

	testClient, _ = NewClient(config)

	httpmock.ActivateNonDefault(testClient.httpClient.GetClient())
	defer httpmock.DeactivateAndReset()

	os.Exit(m.Run())
}

func TestNewClient(t *testing.T) {
	assert.Equal(t, "metric", testClient.unit)
}

func TestNewClient_WithoutApiUnits(t *testing.T) {

	config := WeatherConfig{ApiKey: "apikey"}
	testClient, err := NewClient(config)

	assert.NoError(t, err)
	assert.Equal(t, "metric", testClient.unit)
}

func TestNewClientWithoutApiKey(t *testing.T) {

	config := WeatherConfig{}
	testClient, err := NewClient(config)

	assert.EqualError(t, err, apiKeyNotSupplied)
	assert.Nil(t, testClient)
}

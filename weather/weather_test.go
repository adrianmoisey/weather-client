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
		apiKey: apiKey,
		units:  "metric",
	}

	testClient, _ = NewClient(config)

	httpmock.ActivateNonDefault(testClient.httpClient.GetClient())
	defer httpmock.DeactivateAndReset()

	os.Exit(m.Run())
}

func TestNewClient(t *testing.T) {
	assert.Equal(t, testClient.unit, "metric")
}

func TestNewClientWithoutApiKey(t *testing.T) {

	config := WeatherConfig{}
	testClient, err := NewClient(config)

	assert.EqualError(t, err, apiKeyNotSupplied)
	assert.Nil(t, testClient)
}

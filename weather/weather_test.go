package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient("apikey", "atlanta", "metric")

	assert.Equal(t, client.city, "atlanta")
	assert.Equal(t, client.unit, "metric")

}

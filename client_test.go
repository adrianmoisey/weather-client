package weather

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchWeatherForCity(t *testing.T) {

	client := NewClient()
	res, _ := client.FetchWeatherForCity("atlanta")
	fmt.Println(res)
}

func TestFetchWeatherForCityFail(t *testing.T) {

	client := NewClient()
	res, err := client.FetchWeatherForCity("adasdsdasd")
	assert.Equal(t, err, ErrorNoCityFound)
	assert.Equal(t, res, float64(0))
}

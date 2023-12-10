package utils

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// HitGeoApi sends a request to the ipgeolocation.io API and returns the response.
func HitGeoApi(ip string) (*resty.Response, error) {
	apiKey := "5895a686c4b242aeb8782723ab7a301c"

	response, err := resty.New().R().Get(fmt.Sprintf("https://api.ipgeolocation.io/ipgeo?apiKey=%s&ip=%s", apiKey, ip))
	if err != nil {
		return nil, err
	}

	return response, nil
}

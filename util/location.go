package util

import (
	"alta/airbnb/config"

	"github.com/kelvins/geocoder"
)

// function untuk generate latitude, longitude menggunakan api geocode (google)
func GetGeocodeLocations(fulladdress string) ([]geocoder.Address, float64, float64, error) {
	env := config.GetConfig()
	apikey := env["APIKEYS"]
	geocoder.ApiKey = apikey
	var lng, lat float64
	var address geocoder.Address
	address.City = fulladdress
	location, err := geocoder.Geocoding(address)
	if err != nil {
		return nil, 0, 0, err
	} else {
		lng = location.Latitude
		lat = location.Longitude
	}
	alamats, e := geocoder.GeocodingReverse(location)
	if e != nil {
		return nil, 0, 0, err
	}
	return alamats, lng, lat, nil
}

package util

import (
	"alta/airbnb/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// function untuk generate latitude, longitude menggunakan api geocode (google)
func GetGeocodeLocations(alamat string) (float64, float64, error) {
	env := config.GetConfig()
	apikey := env["APIKEYS"]
	location := strings.ReplaceAll(alamat, " ", "%")
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%v&key=%v", location, apikey)
	response, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		return 0, 0, err
	}
	var result Responses
	json.Unmarshal([]byte(body), &result)
	var lat float64 = result.Result[0].Geometry.Location.Lat
	var lng float64 = result.Result[0].Geometry.Location.Lng
	return lat, lng, nil
}

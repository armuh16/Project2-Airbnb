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
	location := strings.ReplaceAll(alamat, " ", "+")
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%v&key=%v", location, apikey)
	response, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		return 0, 0, err
	}

	var longitude float64
	var latitude float64
	var values map[string]interface{}

	json.Unmarshal(body, &values)
	for _, value := range values["results"].([]interface{}) {
		for str, arr := range value.(map[string]interface{}) {
			if str == "geometry" {
				latitude = arr.(map[string]interface{})["location"].(map[string]interface{})["lat"].(float64)
				longitude = arr.(map[string]interface{})["location"].(map[string]interface{})["lng"].(float64)
				break
			}
		}
	}
	return latitude, longitude, nil
}

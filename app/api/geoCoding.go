package api

import (
	"covidapi/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Document struct {
	ResponseCode int                 `json:"responseCode"`
	Results      []map[string]string `json:"results"`
	Version      string              `json:"version"`
}

func GetByGeoCoordinates(lat string, lng string) (string, error) {
	configuration := config.GetConfigurations()
	url := configuration.Geocoding_api
	api := configuration.Geocoding_apikey
	query := fmt.Sprintf("%s%s/rev_geocode?lat=%s&lng=%s&region=IND", url, api, lat, lng)
	fmt.Println(query)
	res, err := http.Get(query)
	if err != nil {
		log.Println(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var json_data Document
	json.Unmarshal([]byte(body), &json_data)
	fmt.Println(json_data.Results[0]["state"])
	result := json_data.Results[0]
	state := result["state"]
	if state == "Jammu & Kashmir" {
		state = "Jammu and Kashmir"
	}
	if state == "Andaman & Nicobar" {
		state = "Andaman and Nicobar"
	}
	return state, err
}

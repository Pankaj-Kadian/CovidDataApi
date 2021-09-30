package externalapis

import (
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

	query := fmt.Sprintf("https://apis.mapmyindia.com/advancedmaps/v1/ffc5994b9c25a64dc5267125382805b5/rev_geocode?lat=%s&lng=%s&region=IND", lat, lng)
	res, err := http.Get(query)
	if err != nil {
		log.Println(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var json_data Document
	json.Unmarshal([]byte(body), &json_data)
	fmt.Println(json_data)
	result := json_data.Results[0]
	state := result["state"]
	return state, err
}

func main() {
	lat := "28.569548"
	lng := "74.856954"
	GetByGeoCoordinates(lat, lng)
}

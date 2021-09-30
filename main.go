package main

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

func getByGeoCoordinates(lat string, lng string) {

	query := fmt.Sprintf("https://apis.mapmyindia.com/advancedmaps/v1/ffc5994b9c25a64dc5267125382805b5/rev_geocode?lat=%s&lng=%s&region=IND", lat, lng)
	res, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var json_data Document
	json.Unmarshal([]byte(body), &json_data)
	location_data := json_data.Results
	state := location_data[0]["state"]
	fmt.Print(state)

}

func main() {
	// clientOptions := options.Client().
	// 	ApplyURI("mongodb+srv://pankaj:jc420931@cluster.q37h2.mongodb.net/test?retryWrites=true&w=majority")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db := client.Database("covid")
	// collection := db.Collection("statewise")
	// fmt.Println(collection)
	lat := "28.569548"
	lng := "74.856954"
	getByGeoCoordinates(lat, lng)
}

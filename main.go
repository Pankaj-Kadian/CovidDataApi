package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getStateCodes() map[string]string {
	var stateCode map[string]string = map[string]string{"Total": "TT",
		"Andaman and Nicobar Islands": "AN",
		"Andhra Pradesh":              "AP",
		"Arunachal Pradesh":           "AR",
		"Assam":                       "AS",
		"Bihar":                       "BR",
		"Chandigarh":                  "CH",
		"Chhattisgarh":                "CT",
		"Dadra and Nagar Haveli and Daman and Diu": "DN",
		"Delhi":             "DL",
		"Goa":               "GA",
		"Gujarat":           "GJ",
		"Haryana":           "HR",
		"Himachal Pradesh":  "HP",
		"Jammu and Kashmir": "JK",
		"Jharkhand":         "JH",
		"Karnataka":         "KA",
		"Kerala":            "KL",
		"Ladakh":            "LA",
		"Lakshadweep":       "LD",
		"Madhya Pradesh":    "MP",
		"Maharashtra":       "MH",
		"Manipur":           "MN",
		"Meghalaya":         "ML",
		"Mizoram":           "MZ",
		"Nagaland":          "NL",
		"Odisha":            "OR",
		"Puducherry":        "PY",
		"Punjab":            "PB",
		"Rajasthan":         "RJ",
		"Sikkim":            "SK",
		"Tamil Nadu":        "TN",
		"Telangana":         "TG",
		"Tripura":           "TR",
		"Uttar Pradesh":     "UP",
		"Uttarakhand":       "UT",
		"West Bengal":       "WB"}
	return stateCode
}

func routing() {
	// e := echo.New()
}

type ResponseData struct {
	// ID                primitive.ObjectID `json:"_id" bson:"_id" `
	Total_cases       int       `json:"total_cases" bson:"total_cases"`
	Total_recovered   int       `json:"total_recovered" bson:"total_recovered"`
	Total_death       int       `json:"total_death" bson:"total_death"`
	Total_vaccinated1 int       `json:"total_vaccinated1" bson:"total_vaccinated1"`
	Total_vaccinated2 int       `json:"total_vaccinated2" bson:"total_vaccinated2"`
	Total_tested      int       `json:"total_tested" bson:"total_tested"`
	Last_updated      time.Time `json:"last_updated" bson:"last_updated"`
}

func conversion(covid_data interface{}) int {
	var totalInt int64 = int64(covid_data.(float64))
	total := strconv.FormatInt(totalInt, 10)
	t, _ := strconv.Atoi(total)
	return t
}

func main() {
	res, err := http.Get("https://data.covid19india.org/v4/min/data.min.json")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var json_data map[string]interface{}
	json.Unmarshal([]byte(body), &json_data)

	final_data := make(map[string]ResponseData)
	for _, v := range getStateCodes() {
		d := json_data[v].(map[string]interface{})
		// deltaData := d["delta"].(map[string]interface{})
		// today_confirmed_raw := deltaData["confirmed"]
		// fmt.Printf("%T", today_confirmed_raw)
		metaDataOfState := d["meta"].(map[string]interface{})
		totalCasesInState := d["total"].(map[string]interface{})
		last_updated := metaDataOfState["last_updated"]
		total_confirmed := totalCasesInState["confirmed"]
		total_recovered := totalCasesInState["recovered"]
		total_death := totalCasesInState["deceased"]
		total_vaccinated1 := totalCasesInState["vaccinated1"]
		total_vaccinated2 := totalCasesInState["vaccinated2"]
		total_test := totalCasesInState["tested"]
		last_updated_date, _ := time.Parse(time.RFC3339, last_updated.(string))

		total_confirmed_int := conversion(total_confirmed)
		total_recovered_int := conversion(total_recovered)
		total_death_int := conversion(total_death)
		total_vaccinated1_int := conversion(total_vaccinated1)
		total_vaccinated2_int := conversion(total_vaccinated2)
		total_test_int := conversion(total_test)

		var resData ResponseData

		resData.Total_cases = total_confirmed_int
		resData.Total_death = total_death_int
		resData.Total_recovered = total_recovered_int
		resData.Total_vaccinated1 = total_vaccinated1_int
		resData.Total_vaccinated2 = total_vaccinated2_int
		resData.Total_tested = total_test_int
		resData.Last_updated = last_updated_date
		final_data[v] = resData
		// fmt.Printf("%T", metaDataOfState)
		// fmt.Printf("%T   ", last_updated)
		// fmt.Printf("%T   ", result)
		// fmt.Println(resData)
	}
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://pankaj:jc420931@cluster.q37h2.mongodb.net/test?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("covid")
	collection := db.Collection("statewise")

	for _, v := range getStateCodes() {
		// InsertOne using json
		_, err2 := collection.InsertOne(context.Background(), final_data[v])

		// res, err := collection.InsertOne(context.Background(),)
		if err2 != nil {
			fmt.Println(err)
		}
		// fmt.Println(res1)
		fmt.Println(final_data[v])
	}
}

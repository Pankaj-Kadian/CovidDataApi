package api

import (
	"covidapi/config"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func conversion(covid_data interface{}) int {
	var totalInt int64 = int64(covid_data.(float64))
	total := strconv.FormatInt(totalInt, 10)
	t, _ := strconv.Atoi(total)
	return t
}

func GettingData() map[string]config.ResponseData {
	res, err := http.Get("https://data.covid19india.org/v4/min/data.min.json")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var json_data map[string]interface{}
	json.Unmarshal([]byte(body), &json_data)

	final_data := make(map[string]config.ResponseData)
	for _, v := range config.GetStateCodes() {
		d := json_data[v].(map[string]interface{})
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

		var resData config.ResponseData
		resData.State_code = v
		resData.Total_cases = total_confirmed_int
		resData.Total_death = total_death_int
		resData.Total_recovered = total_recovered_int
		resData.Total_vaccinated1 = total_vaccinated1_int
		resData.Total_vaccinated2 = total_vaccinated2_int
		resData.Total_tested = total_test_int
		resData.Last_updated = last_updated_date
		final_data[v] = resData
	}
	// fmt.Println(final_data)
	return final_data
}

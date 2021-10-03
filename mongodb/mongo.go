package mongodb

import (
	"context"
	"covidapi/config"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertingData(final_data map[string]config.ResponseData) {
	collection, err := config.ConnectionMongoDb()
	if err != nil {
		log.Fatal("Unable to Insert")
		return
	}

	for _, v := range config.GetStateCodes() {
		// InsertOne using json
		_, err := collection.InsertOne(context.Background(), final_data[v])
		if err != nil {
			fmt.Print(err)
		}

	}
}
func InsertingNewData(final_data map[string]config.ResponseData) {
	collection, err := config.ConnectionMongoDb()
	if err != nil {
		log.Fatal("Unable to Insert")
		return
	}
	for _, v := range config.GetStateCodes() {
		_, err2 := collection.UpdateOne(context.Background(), bson.M{"state_code": v}, bson.M{"$set": bson.M{
			"Total_cases":       final_data[v].Total_cases,
			"Total_recovered":   final_data[v].Total_recovered,
			"Total_death":       final_data[v].Total_death,
			"Total_vaccinated1": final_data[v].Total_vaccinated1,
			"Total_vaccinated2": final_data[v].Total_vaccinated2,
			"Total_tested":      final_data[v].Total_tested,
			"Last_updated":      final_data[v].Last_updated,
		}})
		if err2 != nil {
			fmt.Println(err2)
		}
	}
}

func GetData(state_code string) (config.ResponseData, error) {
	collection, err := config.ConnectionMongoDb()
	if err != nil {
		log.Fatal("Unable to Fetch Data")
	}
	var findOne config.ResponseData
	err = collection.FindOne(context.Background(), bson.M{"state_code": state_code}).Decode(&findOne)
	if err != nil {
		log.Fatal("No Record Found")
	}
	// TO convert time in IST
	loc, _ := time.LoadLocation("Asia/Kolkata")
	local := findOne.Last_updated
	if err == nil {
		local = local.In(loc)
	}
	findOne.Last_updated = local
	// fmt.Println(findOne)
	return findOne, err
}

func GetAllData() ([]config.ResponseData, error) {
	collection, err := config.ConnectionMongoDb()
	if err != nil {
		log.Fatal("Unable to Insert")
	}
	var findOne config.ResponseData
	var find []config.ResponseData
	find_cursor, err := collection.Find(context.Background(), bson.M{})
	for find_cursor.Next(context.Background()) {
		err := find_cursor.Decode(&findOne)
		if err != nil {
			fmt.Println(err)
		}
		loc, _ := time.LoadLocation("Asia/Kolkata")
		local := findOne.Last_updated
		if err == nil {
			local = local.In(loc)
		}
		findOne.Last_updated = local
		find = append(find, findOne)

	}
	if err != nil {
		log.Fatal(err)
	}
	return find, err
}

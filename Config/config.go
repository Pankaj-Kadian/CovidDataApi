package Config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var state_code map[string]string = map[string]string{"Total": "TT",
	"Andaman & Nicobar Islands": "AN",
	"Andhra Pradesh":            "AP",
	"Arunachal Pradesh":         "AR",
	"Assam":                     "AS",
	"Bihar":                     "BR",
	"Chandigarh":                "CH",
	"Chhattisgarh":              "CT",
	"Delhi":                     "DL",
	"Goa":                       "GA",
	"Gujarat":                   "GJ",
	"Haryana":                   "HR",
	"Himachal Pradesh":          "HP",
	"Jammu and Kashmir":         "JK",
	"Jharkhand":                 "JH",
	"Karnataka":                 "KA",
	"Kerala":                    "KL",
	"Ladakh":                    "LA",
	"Lakshadweep":               "LD",
	"Madhya Pradesh":            "MP",
	"Maharashtra":               "MH",
	"Manipur":                   "MN",
	"Meghalaya":                 "ML",
	"Mizoram":                   "MZ",
	"Nagaland":                  "NL",
	"Odisha":                    "OR",
	"Puducherry":                "PY",
	"Punjab":                    "PB",
	"Rajasthan":                 "RJ",
	"Sikkim":                    "SK",
	"Tamil Nadu":                "TN",
	"Telangana":                 "TG",
	"Tripura":                   "TR",
	"Uttar Pradesh":             "UP",
	"Uttarakhand":               "UT",
	"West Bengal":               "WB"}

func GetStateCodes() map[string]string {
	return state_code
}

func StateCodeFromStateName() map[string]string {
	state_name := make(map[string]string)
	for k, v := range state_code {
		state_name[v] = k
	}
	return state_name
}
func ConnectionMongoDb() *mongo.Collection {
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
	return collection
}

// func StateCode() map[string]string{

// 	return stateCode
// }

package main

import (
	handlers "Covid-Data-Api/Handlers"

	"github.com/labstack/echo/v4"
)

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
	// final_data := externalapis.GettingData()
	// externalapis.InsertingData(final_data)

	e := echo.New()

	e.GET("/getCases", handlers.GetCases)
	e.GET("/getAllCases", handlers.GetAllCases)
	e.GET("/getbyCoordinates", handlers.GetDataFromGeoLocation)
	e.Start("localhost:8080")
}

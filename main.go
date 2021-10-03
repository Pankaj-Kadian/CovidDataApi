package main

import (
	"covidapi/app/api"
	"covidapi/app/handlers"
	"covidapi/mongodb"
	"fmt"
	"time"

	_ "covidapi/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {

	e := echo.New()
	go func() {
		for {
			final_data := api.GettingData()
			fmt.Println(final_data)
			mongodb.InsertingNewData(final_data)
			time.Sleep(30 * time.Minute)
		}
	}()
	e.GET("GetStateData", handlers.GetCases)
	e.GET("/GetAllData", handlers.GetAllCases)
	e.GET("/GetByGeoLocation", handlers.GetDataFromGeoLocation)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

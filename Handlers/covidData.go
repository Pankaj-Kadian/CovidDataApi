package handlers

import (
	externalapis "Covid-Data-Api/ExternalApis"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCases(c echo.Context) error {

	state := c.QueryParam("state")

	fmt.Println(state)
	data, err := externalapis.GetData(state)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Invaild StateName")
	}
	return c.JSON(http.StatusOK, data)
}

func GetAllCases(c echo.Context) error {
	data, err := externalapis.GetAllData()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Invaild StateName")
	}
	return c.JSON(http.StatusOK, data)
}

func GetDataFromGeoLocation(c echo.Context) error {
	lat := c.QueryParam("latitude")
	lng := c.QueryParam("longitude")
	fmt.Println(lat, lng)
	state, err := externalapis.GetByGeoCoordinates(lat, lng)
	if err != nil {
		log.Println(err)
	}
	data, err := externalapis.GetData(state)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Invaild StateName")
	}
	return c.JSON(http.StatusOK, data)

}

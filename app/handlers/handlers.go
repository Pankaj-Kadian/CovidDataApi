package handlers

import (
	"covidapi/app/api"
	"covidapi/config"
	"covidapi/mongodb"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCases godoc
// @Summary Get a State Cases
// @Description get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param state query string true "state"
// @Success 200 {object} model.OneState
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /GetStateData [get]

func GetCases(c echo.Context) error {
	valid := c.Request().URL.Query().Has("state")

	if !valid {
		return c.JSON(http.StatusBadRequest, "Invalid query string key name it should be state")
	}
	state := c.QueryParam("state")

	fmt.Println(state)
	state_code, present := config.GetStateCodes()[state]
	if !present {
		states := config.GetStateCodes()
		state_list := ""
		for k := range states {
			state_list = state_list + k + ","
		}
		list := fmt.Sprint("Please provide a valid state name... Here is the list of availabe state name........", state_list)
		return c.JSON(http.StatusBadRequest, list)
	}
	fmt.Println(state_code)
	data, err := mongodb.GetData(state_code)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, data)
}

// GetAllCases godoc
// @Summary Get a State Cases
// @Description get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} model.AllState
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /GetAllData [get]

func GetAllCases(c echo.Context) error {
	data, err := mongodb.GetAllData()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Invaild StateName")
	}
	return c.JSON(http.StatusOK, data)
}

// GetByGeoLocation godoc
// @Summary Get a State Cases
// @Description get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param latitude query string true "latitude"
// @Param longitude query string true "longitude"
// @Success 200 {object} model.OneState
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /GetByGeoLocation [get]

func GetDataFromGeoLocation(c echo.Context) error {
	latitude := c.Request().URL.Query().Has("latitude")
	longitude := c.Request().URL.Query().Has("longitude")
	if !latitude || !longitude {
		return c.JSON(http.StatusBadRequest, "Invalid query string key name it should be lat and lng")
	}
	lat := c.QueryParam("latitude")
	lng := c.QueryParam("longitude")
	state, err := api.GetByGeoCoordinates(lat, lng)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Please give coordinates of India")
	}
	data, err := mongodb.GetData(state)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Invaild StateName")
	}
	return c.JSON(http.StatusOK, data)

}

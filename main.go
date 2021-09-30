package main

import (
	"/ExternalApis"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCases(c echo.Context) error {

	state, err := strconv.Atoi(c.Param("state"))
	if err != nil {
		return err
	}
	data, err := ExternalApis.GetData(state)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Invaild StateName")
	}
	return c.JSON(http.StatusOK, data)
}

func getAllCases(c echo.Context) error {
	data, err := ExternalApis.GetData()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Invaild StateName")
	}
	return c.JSON(http.StatusOK, data)
}

func main() {
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
	fmt.Println(collection)

	e := echo.New()

	e.GET("/getCases:states", getCases)
	e.GET("/getCases", getCases)
}

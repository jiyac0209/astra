package main

import (
	"astraSecurity/core"
	"astraSecurity/repository"
	"fmt"

	"astraSecurity/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

// todo add logging and commetns
func main() {

	//connecto to db
	db := repository.Connect()
	repository.CreateSchema(db)
	defer db.Close()

	// Start a goroutine to process files, if any file is found it will push it to db and delete the file
	go core.ProcessFiles(db, "data")

	// Create a new Echo instance
	e := echo.New()

	// Define a route
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.POST("/push", Push)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
	select {}

}

// push handler
func Push(c echo.Context) error {
	var astra domain.Astra

	// parse HTTP body to astra
	if err := c.Bind(&astra); err != nil {
		return err
	}

	if astra.Data == "" {
		return fmt.Errorf("invalid payload")
	}
	// push data to a file in /data
	err := core.Push(&astra)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, astra)

}

package main

import (
	"math/rand"
	"net/http"
	"rosatomcase/backend/temperature"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	MaxRange = 10.0
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	tempSensors := temperature.TempArray{make([]*temperature.Temperature, 10)}

	for i := 0; i < 10; i++ {
		name := "id" + strconv.Itoa(i)
		tempSensors.Array[i] = &temperature.Temperature{}
		go tempSensors.Array[i].Generate(name, MaxRange)
	}

	// Route => handler
	e.GET("/temp", func(c echo.Context) error {
		return c.JSON(http.StatusOK, tempSensors.Retrieve())
	})

	// Start server
	e.Logger.Fatal(e.Start(":8090"))
}

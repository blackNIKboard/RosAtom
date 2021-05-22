package main

import (
	"math/rand"
	"net/http"
	"rosatomcase/backend/model"
	"rosatomcase/backend/sensor"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	TempWarn = model.ValueWarning{
		Minimum: 15,
		Maximum: 32,
	}
	EnergyWarn = model.ValueWarning{
		Minimum: 0,
		Maximum: 1000,
	}
	TempSensorsCount = 10
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	sensors := sensor.Array{make([]*sensor.Sensor, TempSensorsCount)}

	for i := 0; i < TempSensorsCount; i++ {
		name := "id" + strconv.Itoa(i)
		sensors.Array[i] = &sensor.Sensor{}
		go sensors.Array[i].Generate(name, TempWarn, EnergyWarn)
	}

	// Route => handler
	e.GET("/temp", func(c echo.Context) error {
		return c.JSON(http.StatusOK, sensors.Retrieve())
	})

	// Start server
	e.Logger.Fatal(e.Start(":8090"))
}

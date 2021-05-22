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
	MaxRange         = 10.0
	TempSensorsCount = 10
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete}, MaxAge: 300}))
	e.Use(middleware.CORS())

	tempSensors := temperature.TempArray{make([]*temperature.Temperature, TempSensorsCount)}

	for i := 0; i < TempSensorsCount; i++ {
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

package main

import (
	"math/rand"
	"net/http"
	"rosatomcase/backend/temperature"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	temp := temperature.Temperature{}

	go temp.Generate()

	// Route => handler
	e.GET("/temp", func(c echo.Context) error {
		return c.JSON(http.StatusOK, temp.Last())
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

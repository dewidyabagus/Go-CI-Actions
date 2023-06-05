package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

const defAddress = ":8080"

func main() {
	e := echo.New()

	// TO DO: Tambahkan endpoint lainnya
	e.GET("/welcome", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome, Echo Web Framework")
	})

	// TO DO: Gracefully shutdown
	e.Logger.Fatal(e.Start(defAddress))
}

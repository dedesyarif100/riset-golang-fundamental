package main

import (
	"fmt"
	"net/http"
	"github.com/rs/cors"
	"github.com/labstack/echo"
)

func main() {
    e := echo.New()

    // ...
	corsMiddleware := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://novalagung.com", "https://www.google.com"},
		// AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		// AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		// Debug:          true,

		AllowedOrigins: []string{"https://linktr.ee/novalagung", "https://www.google.com"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))


    e.GET("/index", func(c echo.Context) error {
		fmt.Println("MASUK URL INDEX")
		return c.String(http.StatusOK, "hello")
    })

    e.Logger.Fatal(e.Start(":9000"))
}
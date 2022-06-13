package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        fmt.Println("from middleware one")
        return next(c)
    }
}

func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        fmt.Println("from middleware two")
        return next(c)
    }
}

func middlewareSomething(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("from middleware something")
        next.ServeHTTP(w, r)
    })
}

func main() {
    e := echo.New()

    // middleware here
	// C.8.1. Custom Middleware
		e.Use(middlewareOne)
		e.Use(middlewareTwo)

		// C.8.2. Integrasi Middleware ber-skema Non-Echo-Middleware
		e.Use(echo.WrapMiddleware(middlewareSomething))

		// C.8.3. Echo Middleware: Logger
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		}))

		e.GET("/index", func(c echo.Context) (err error) {
			fmt.Println("threeeeee!")
			return c.JSON(http.StatusOK, true)
		})

    e.Logger.Fatal(e.Start(":9000"))
}
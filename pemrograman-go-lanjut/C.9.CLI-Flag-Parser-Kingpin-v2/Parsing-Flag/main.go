package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app         = kingpin.New("App", "Simple app")
	flagAppName = app.Flag("name", "Application name").Required().String()
	flagPort    = app.Flag("port", "Web server port").Short('p').Default("9000").Int()
)

func main() {
    // kingpin.Parse()

	// more code here ...
	// C.9.1. Parsing Argument
	// C.9.2. Penggunaan Kingpin Application Instance
		flagName := *flagAppName
		port := fmt.Sprintf(":%d", *flagPort)
		// app.MustParse(app.Parse(os.Args[1:]))
		fmt.Printf("Starting %s at %s", flagName, port)
		kingpin.MustParse(app.Parse(os.Args[1:]))
		e := echo.New()
		e.GET("/index", func(c echo.Context) (err error) {
			fmt.Println("MASUK URL INDEX")
			return c.JSON(http.StatusOK, true)
		})

	e.Logger.Fatal(e.Start(port))
}
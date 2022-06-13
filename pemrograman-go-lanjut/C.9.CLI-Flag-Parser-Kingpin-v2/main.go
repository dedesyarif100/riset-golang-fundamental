package main

import (
	"fmt"
	"net/http"
	// "os"
	"github.com/labstack/echo"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app        = kingpin.New("App", "Simple app")
    argAppName = app.Arg("name", "laravel").Required().String()
    argPort    = app.Arg("port", "Web server port").Default("9000").Int()
)

func main() {
    // kingpin.Parse()

	// more code here ...
	// C.9.1. Parsing Argument
	// C.9.2. Penggunaan Kingpin Application Instance
		appName := *argAppName
		port := fmt.Sprintf(":%d", *argPort)
		// app.MustParse(app.Parse(os.Args[1:]))
		fmt.Printf("Starting %s at %s", appName, port)

		e := echo.New()
		e.GET("/index", func(c echo.Context) (err error) {
			fmt.Println("MASUK URL INDEX")
			return c.JSON(http.StatusOK, true)
		})

	



	e.Logger.Fatal(e.Start(port))
}
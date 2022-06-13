package main

import (
	"fmt"
	"net/http"
	"github.com/spf13/viper"
	"github.com/labstack/echo"
	"github.com/fsnotify/fsnotify"
)

func main() {
    e := echo.New()

    viper.SetConfigType("json")
    viper.AddConfigPath(".")
    viper.SetConfigName("app.config")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

    err := viper.ReadInConfig()
    if err != nil {
        e.Logger.Fatal(err)
    }

    // ...
	e.GET("/index", func(c echo.Context) (err error) {
		fmt.Println("MASUK URL INDEX")
		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Print("Starting", viper.GetString("name_appName"))
	e.Logger.Fatal(e.Start(":" + viper.GetString("server.port")))
}
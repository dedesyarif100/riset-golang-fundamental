package main

import (
    // "fmt"
    "github.com/labstack/echo"
    "net/http"
    // "strings"
)

type M map[string]interface{}

func main() {
    r := echo.New()

    r.GET("/", func(ctx echo.Context) error {
        data := "Hello from /index"
        return ctx.String(http.StatusOK, data)
    })

    r.Start(":9000")
}
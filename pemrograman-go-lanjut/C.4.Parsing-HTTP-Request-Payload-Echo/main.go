package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)

type User struct {
    Name  string `json:"name" form:"name" query:"name"`
    Email string `json:"email" form:"email" query:"email"`
}

func main() {
    r := echo.New()

    // C.4.1. Parsing Request Payload
	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		fmt.Println("MASUK USER")
		fmt.Println(u)
		return c.JSON(http.StatusOK, u)
	})

    fmt.Println("server started at :9000")
    r.Start(":9000")
}
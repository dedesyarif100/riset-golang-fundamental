package main

import (
	// "fmt"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type M map[int][]int

func main() {
	request := echo.New()

    // C.3.3. Response Method milik ctx
		// Method .String()
		request.GET("/", func(ctx echo.Context) error {
			data := "HELLO FROM /INDEX"
			return ctx.String(http.StatusOK, data)
		})

		// Method .HTML()
		request.GET("/html", func(ctx echo.Context) error {
			data := "<u>HELLO </u> FROM /HTML"
			return ctx.HTML(http.StatusOK, data)
		})

		// Method .Redirect()
		request.GET("/index", func(ctx echo.Context) error {
			return ctx.Redirect(http.StatusTemporaryRedirect, "/")
		})

		// Method .JSON()
		request.GET("/json", func(ctx echo.Context) error {
			data := M {
						0: []int{1, 2, 3}, 
						1: []int{1, 2, 3},
					}
			return ctx.JSON(http.StatusOK, data)
		})

		
    // C.3.4. Parsing Request
		// Parsing Query String
		request.GET("page1", func(ctx echo.Context) error {
			fmt.Println("MASUK PAGE 1")
			name := ctx.QueryParam("name")
			data := fmt.Sprintf("HELLO %s\n", name)
			return ctx.String(http.StatusOK, data)
		})

		// Parsing URL Path Param
		request.GET("/page2/:name", func(ctx echo.Context) error {
			fmt.Println("MASUK PAGE 2")
			name := ctx.Param("name")
			data := fmt.Sprintf("HELLO %s\n", name)
			return ctx.String(http.StatusOK, data)
		})

    request.Start(":9000")
}
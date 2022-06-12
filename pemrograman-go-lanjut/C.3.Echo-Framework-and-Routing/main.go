package main

import (
	// "fmt"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	// "strings"
)

type M map[string]interface{}

func main() {
    request := echo.New()

    // C.3.3. Response Method milik ctx
        // • Method .String()
        request.GET("/", func(ctx echo.Context) error {
            data := "Hello from /index"
            return ctx.String(http.StatusOK, data)
        })

        // • Method .HTML()
        request.GET("/html", func(ctx echo.Context) error {
            data := "Hello from /html"
            return ctx.HTML(http.StatusOK, data)
        })

        // • Method .Redirect()
        request.GET("/index", func(ctx echo.Context) error {
            return ctx.Redirect(http.StatusTemporaryRedirect, "/")
        })

        // • Method .JSON()
        request.GET("/json", func(ctx echo.Context) error {
            data := M{"Message": "Hello", "Counter": 2}
            return ctx.JSON(http.StatusOK, data)
        })


    // C.3.4. Parsing Request
        // • Parsing Query String
        request.GET("page1", func(ctx echo.Context) error {
            fmt.Println("MASUK PAGE 1")
            name := ctx.QueryParam("name")
            data := fmt.Sprintf("Hello %s\n", name)
            return ctx.String(http.StatusOK, data)
        })

        // • Parsing URL Path Param
        request.GET("/page2/:name", func(ctx echo.Context) error {
            fmt.Println("MASUK PAGE 2")
            name := ctx.Param("name")
            data := fmt.Sprintf("Hello %s\n", name)
            return ctx.String(http.StatusOK, data)
        })

        // • Parsing URL Path Param dan Setelahnya
        request.GET("/page3/:name/*", func(ctx echo.Context) error {
            fmt.Println("MASUK PAGE 3")
            name := ctx.Param("name")
            message := ctx.Param("*")
            data := fmt.Sprintf("Hello %s, I have message for you: %s\n", name, message)
            return ctx.String(http.StatusOK, data)
        })

        // • Parsing Form Data
        request.POST("/page4", func(ctx echo.Context) error {
            fmt.Println("MASUK PAGE 4")
            name := ctx.FormValue("name")
            message := ctx.FormValue("message")
        
            data := fmt.Sprintf(
                "Hello %s, I have message for you: %s\n",
                name,
                strings.Replace(message, "/", "", 1),
            )
        
            return ctx.String(http.StatusOK, data)
        })


    // C.3.5. Penggunaan echo.WrapHandler Untuk Routing Handler Bertipe
        request.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
        request.GET("/home", echo.WrapHandler(ActionHome))
        request.GET("/about", ActionAbout)

    request.Start(":9000")
}


// C.3.5. Penggunaan echo.WrapHandler Untuk Routing Handler Bertipe
var ActionIndex = func(response http.ResponseWriter, request *http.Request) {
    fmt.Println("MASUK INDEX")
    response.Write([]byte("from action index"))
}

var ActionHome = http.HandlerFunc(
    func(response http.ResponseWriter, request *http.Request) {
        fmt.Println("MASUK HOME")
        response.Write([]byte("from action home"))
    },
)

var ActionAbout = echo.WrapHandler(
    http.HandlerFunc(
        func(response http.ResponseWriter, request *http.Request) {
            fmt.Println("MASUK ABOUT")
            response.Write([]byte("from action about"))
        },
    ),
)
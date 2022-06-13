package main

import (
	"io"
	"fmt"
	"net/http"
	"html/template"
	"github.com/labstack/echo"
)

type M map[string]interface{}

type Renderer struct {
    template *template.Template
    location string
    debug    bool
}

func NewRenderer(location string, debug bool) *Renderer {
	fmt.Println("MASUK NEW RENDERER")
    tpl := new(Renderer)
    tpl.location = location
    tpl.debug = debug

    tpl.ReloadTemplates()

    return tpl
}

func (t *Renderer) ReloadTemplates() {
	fmt.Println("MASUK RELOAD TEMPLATES")
    t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
    w io.Writer, 
    name string, 
    data interface{}, 
    c echo.Context,
) error {
	fmt.Println("MASUK RENDER")
    if t.debug {
        t.ReloadTemplates()
    }

    return t.template.ExecuteTemplate(w, name, data)
}

func main() {
    e := echo.New()

    e.Renderer = NewRenderer("template/*.html", true)

    e.GET("/index", func(ctx echo.Context) error {
        data := M{"message": "Hello World!"}
        return ctx.Render(http.StatusOK, "index.html", data)
    })

    e.Logger.Fatal(e.Start(":9000"))
}
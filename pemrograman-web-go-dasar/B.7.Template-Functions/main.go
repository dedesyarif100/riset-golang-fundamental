package main

import (
    "html/template"
    "net/http"
    "fmt"
)

type Superhero struct {
    Name    string
    Alias   string
    Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
    return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
    http.HandleFunc("/", func(response http.ResponseWriter, r *http.Request) {
        var person = Superhero{
            Name:    "Bruce Wayne",
            Alias:   "Batman",
            Friends: []string{"Superman", "Flash", "Green Lantern"},
        }

        var tmpl = template.Must(template.ParseFiles("view.html"))
        if err := tmpl.Execute(response, person); err != nil {
            http.Error(response, err.Error(), http.StatusInternalServerError)
        }
    })

    fmt.Println("server started at localhost:7000")
    http.ListenAndServe(":7000", nil)
}
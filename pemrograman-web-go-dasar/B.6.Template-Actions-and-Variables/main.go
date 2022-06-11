package main

import (
    "net/http"
    "fmt"
    "html/template"
)

type Info struct {
    Affiliation string
    Address     string
}

type Person struct {
    Name    string
    Gender  string
    Hobbies []string
    Info    Info
    IsTrue  bool
    Abjad   string
}

func (t Info) GetAffiliationDetailInfo() string {
    return "have 31 divisions"
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var person = Person{
            Name:    "Bruce Wayne",
            Gender:  "male",
            Hobbies: []string{"Reading Books", "Traveling", "Buying things"},
            Info:    Info{"Wayne Enterprises", "Gotham City"},
            IsTrue:  true,
            Abjad :  "B",
        }

        var tmpl = template.Must(template.ParseFiles("view.html"))
        if err := tmpl.Execute(w, person); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    fmt.Println("server started at localhost:7000")
    http.ListenAndServe(":7000", nil)
}
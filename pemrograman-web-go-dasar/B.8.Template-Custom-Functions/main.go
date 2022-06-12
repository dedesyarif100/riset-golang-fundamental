package main

import (
    "fmt"
    "net/http"
    "html/template"
)

var num = 0
var funcMap = template.FuncMap{
    "unescape": func(s string) template.HTML {
        num++
        fmt.Println(num)
        return template.HTML(s)
    },
    "avg": func(n ...int) int {
        var total = 0
        for _, each := range n {
            total += each
        }
        return total / len(n)
    },
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var tmpl = template.Must(template.New("view.html").
                    Funcs(funcMap).
                    ParseFiles("view.html"));

        if err := tmpl.Execute(w, nil); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    });

    fmt.Println("server started at localhost:7000")
    http.ListenAndServe(":7000", nil)
}
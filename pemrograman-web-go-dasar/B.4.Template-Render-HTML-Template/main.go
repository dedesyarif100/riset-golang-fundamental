package main

import (
	"fmt"
	"path"
	"net/http"
	"html/template"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var data = map[string]interface{}{
			"title": "Learning Golang Web",
			"name":  "Batman",
			"age": 22,
			"vocation" : "SOFTWARE ENGINEER",
			"birth" : "SURABAYA, INDONESIA",
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// -----------------------------------------------------------
    })


	http.HandleFunc("/ExampleOne", func(response http.ResponseWriter, request *http.Request) {
		var FilePathTwo = path.Join("views", "ExampleOne.html")
		TempletaTwo, err := template.ParseFiles(FilePathTwo)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		data2 := map[string]any{
			"numberOne" : 11,
			"numberTwo" : 22,
		}
		err = TempletaTwo.Execute(response, data2)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}
	})

    // http.Handle("/static/", 
    //     http.StripPrefix("/static/", 
    //         http.FileServer(http.Dir("assets"))))

    fmt.Println("server started at localhost:8000")
    http.ListenAndServe(":8000", nil)
}
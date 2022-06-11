package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	// var tmpl, err = template.ParseGlob("views/*")
	// if err != nil {
	//     panic(err.Error())
	//     return
	// }

	http.HandleFunc("/index", func(response http.ResponseWriter, request *http.Request) {
		var data = M{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))
		err := tmpl.ExecuteTemplate(response, "index", data)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(response http.ResponseWriter, request *http.Request) {
		var data = M{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))
		err := tmpl.ExecuteTemplate(response, "about", data)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:7000")
	http.ListenAndServe(":7000", nil)
}

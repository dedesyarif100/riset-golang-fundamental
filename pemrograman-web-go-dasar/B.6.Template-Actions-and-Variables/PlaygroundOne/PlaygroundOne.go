package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address		string
}

type Person struct {
	Name	string
	Gender	string
	Hobbies	[]string
	Info	Info
	IsTrue	bool
	Abjad	string
}

func (t Info) GetAffiliationDetailInfo() string {
	return "HAVE 31 DIVISIONS"
}

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		var person = Person{
			Name	: "DEDE SYARIF",
			Gender	: "MALE",
			Hobbies	: []string{"PROGRAMMING", "FOOTBALL"},
			Info	: Info{"Surabaya", "INDONESIA"},
			IsTrue	: true,
			Abjad	: "B",
		}

		Template := template.Must(template.ParseFiles("view.html"))
		if err := Template.Execute(response, person); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("SERVER STARTED AT LOCALHOST:7000")
	http.ListenAndServe(":7000", nil)
}
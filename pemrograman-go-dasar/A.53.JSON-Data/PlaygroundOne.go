package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	FullName string	`json:"Name"`
	Age 	 int	`json:"umur"`
	Vocation string	`json:"pekerjaan"`
	Skill	 string	`json:"keahlian"`
}

type Merchant struct {
	FullName string
	Age 	 int
	Vocation string
	Skill 	 string
}

func main() {
	fmt.Println("# - A.53.4. Encode Objek Ke JSON String");
		var object = []Merchant{
			{"TES", 22, "tes", "tes"},
			{"TES", 22, "tes", "tes"},
		}

		jsonDataExFour, errExFour := json.Marshal(object)
		if errExFour != nil {
			fmt.Println(errExFour.Error())
			return
		}

		jsonMarshalExFour := string(jsonDataExFour)
		fmt.Println(jsonMarshalExFour)
}
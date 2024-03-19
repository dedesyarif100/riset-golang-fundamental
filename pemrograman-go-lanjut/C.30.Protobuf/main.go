package main

import (
	// sesuaikan dengan struktuk folder project masing2
	"encoding/json"
	"fmt"
	"os"

	"C.30.Protobuf/model"
)

var user1 = &model.User{
	Id:       "u001",
	Name:     "Sylvana Windrunner",
	Password: "f0r Th3 H0rD3",
	Gender:   model.UserGender_FEMALE,
}

var userList = &model.UserList{
	List: []*model.User{
		user1,
	},
}

var garage1 = &model.Garage{
	Id:   "g001",
	Name: "Kalimdor",
	Coordinate: &model.GarageCoordinate{
		Latitude:  23.2212847,
		Longitude: 53.22033123,
	},
}

var garageList = &model.GarageList{
	List: []*model.Garage{
		garage1,
	},
}

var garageListByUser = &model.GarageListByUser{
	List: map[string]*model.GarageList{
		user1.Id: garageList,
	},
}

func main() {
	// more code here ...
	fmt.Println("Original : ", user1)

	fmt.Println("AS String : ", user1.String())

	// Konversi objek proto ke json string
	jsonb, err1 := json.Marshal(garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	fmt.Println("AS JSON String : ", string(jsonb))

	// Konversi json string ke objek proto
	protoObject := new(model.GarageList)
	err2 := json.Unmarshal(jsonb, protoObject)

	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Println("AS String : ", protoObject.String())

}

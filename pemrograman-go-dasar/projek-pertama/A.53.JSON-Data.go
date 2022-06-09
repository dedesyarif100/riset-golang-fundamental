package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
    FullName string `json:"Name"`
    Age      int	`json:"umur"`
	Vocation string `json:"pekerjaan"`
	Skill	 string `json:"keahlian"`
}

type Merchant struct {
	FullName string
    Age      int
	Vocation string
	Skill	 string
}

func main() {
	// A.53.1. Decode JSON Ke Variabel Objek Struct
	fmt.Println("# - A.53.1. Decode JSON Ke Variabel Objek Struct");
		var jsonString = `{"Name": "john wick", "umur": 27, "pekerjaan": "SOFTWARE ENGINEER", "keahlian": "GOLANG, NODEJS, LARAVEL"}`;
		var jsonData = []byte(jsonString);

		var data User;
		// fmt.Println(jsonString);

		var err = json.Unmarshal(jsonData, &data);
		// fmt.Println(err);
		if err != nil {
			fmt.Println(err.Error());
			return;
		}

		fmt.Println("USER 		:", data.FullName);
		fmt.Println("AGE  		:", data.Age);
		fmt.Println("VOCATION	:", data.Vocation);
		fmt.Println("SKILL  		:", data.Skill);
		fmt.Println();

	// A.53.2. Decode JSON Ke map[string]interface{} & interface{}
	fmt.Println("# - A.53.2. Decode JSON Ke map[string]interface{} & interface{}");
		var data1 map[string]interface{};
		json.Unmarshal(jsonData, &data1);

		fmt.Println("USER 		:", data1["Name"]);
		fmt.Println("AGE  		:", data1["umur"]);
		fmt.Println("VOCATION  	:", data1["pekerjaan"]);
		fmt.Println("SKILL  		:", data1["keahlian"]);

		fmt.Println("-----------------------------------------");
		var data2 interface{};
		json.Unmarshal(jsonData, &data2);

		var decodedData = data2.(map[string]interface{});
		fmt.Println("USER 		:", decodedData["Name"]);
		fmt.Println("AGE  		:", decodedData["umur"]);
		fmt.Println("VOCATION  	:", decodedData["pekerjaan"]);
		fmt.Println("SKILL  		:", decodedData["keahlian"]);
		fmt.Println();

	// A.53.3. Decode Array JSON Ke Array Objek
	fmt.Println("# - A.53.3. Decode Array JSON Ke Array Objek");
		var jsonStringExThree = `[
			{"FullName": "john wick", "Age": 27, "Vocation": "SOFTWARE ENGINEER", "Skill": "GOLANG"},
			{"FullName": "ethan hunt", "Age": 32, "Vocation": "BACKEND ENGINEER", "Skill": "LARAVEL"}
		]`;

		var dataExThree []Merchant;
		
		var errExThree = json.Unmarshal([]byte(jsonStringExThree), &dataExThree);
		fmt.Println( dataExThree );
		if errExThree != nil {
			fmt.Println(errExThree.Error());
			return;
		}

		fmt.Println("NAME		:", dataExThree[0].FullName);
		fmt.Println("AGE		:", dataExThree[0].Age);
		fmt.Println("VOCATION	:", dataExThree[0].Vocation);
		fmt.Println("SKILL		:", dataExThree[0].Skill);
		fmt.Println();

		fmt.Println("NAME		:", dataExThree[1].FullName);
		fmt.Println("AGE		:", dataExThree[1].Age);
		fmt.Println("VOCATION	:", dataExThree[1].Vocation);
		fmt.Println("SKILL		:", dataExThree[1].Skill);
		fmt.Println();

	// A.53.4. Encode Objek Ke JSON String
	fmt.Println("# - A.53.4. Encode Objek Ke JSON String");
		var object = []Merchant{
						{"john wick", 27, "SOFTWARE ENGINEER", "GOLANG"}, 
						{"ethan hunt", 32, "BACKEND ENGINEER", "LARAVEL"},
					};
		// fmt.Println( reflect.TypeOf(object) );
		fmt.Println( object )
		fmt.Println();

		var jsonDataExFour, errExFour = json.Marshal(object);
		if errExFour != nil {
			fmt.Println(errExFour.Error());
			return;
		}

		var jsonStringExFour = string(jsonDataExFour);
		fmt.Println( jsonStringExFour );
		// fmt.Println( reflect.TypeOf(jsonStringExFour) );
		fmt.Println();
}
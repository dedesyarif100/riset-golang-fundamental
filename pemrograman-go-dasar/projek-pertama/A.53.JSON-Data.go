package main

import "encoding/json"
import "fmt"

type User struct {
    FullName string `json:"Name"`
    Age      int
}

func main() {
	// A.53.1. Decode JSON Ke Variabel Objek Struct
	fmt.Println("# - A.53.1. Decode JSON Ke Variabel Objek Struct");
		var jsonString = `{"Name": "john wick", "Age": 27}`;
		var jsonData = []byte(jsonString);

		var data User;

		var err = json.Unmarshal(jsonData, &data);
		// fmt.Println(err);
		if err != nil {
			fmt.Println(err.Error());
			return;
		}

		fmt.Println("user :", data.FullName);
		fmt.Println("age  :", data.Age);

	// A.53.2. Decode JSON Ke map[string]interface{} & interface{}
	fmt.Println("# - A.53.2. Decode JSON Ke map[string]interface{} & interface{}");
		var data1 map[string]interface{};
		json.Unmarshal(jsonData, &data1);

		fmt.Println("user :", data1["Name"]);
		fmt.Println("age  :", data1["Age"]);

		fmt.Println("-----------------------------------------");
		var data2 interface{};
		json.Unmarshal(jsonData, &data2);

		var decodedData = data2.(map[string]interface{});
		fmt.Println("user :", decodedData["Name"]);
		fmt.Println("age  :", decodedData["Age"]);

	// A.53.3. Decode Array JSON Ke Array Objek
	fmt.Println("# - A.53.3. Decode Array JSON Ke Array Objek");
		var jsonStringExThree = `[
			{"Name": "john wick", "Age": 27},
			{"Name": "ethan hunt", "Age": 32}
		]`;

		var dataExThree []User;

		var errExThree = json.Unmarshal([]byte(jsonStringExThree), &dataExThree);
		if errExThree != nil {
			fmt.Println(errExThree.Error());
			return;
		}

		fmt.Println("user 1:", dataExThree[0].FullName);
		fmt.Println("user 2:", dataExThree[1].FullName);

	// A.53.4. Encode Objek Ke JSON String
	fmt.Println("# - A.53.4. Encode Objek Ke JSON String");
		var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
		var jsonDataExFour, errExFour = json.Marshal(object)
		if errExFour != nil {
			fmt.Println(errExFour.Error())
			return
		}

		var jsonStringExFour = string(jsonDataExFour)
		fmt.Println(jsonStringExFour)
}
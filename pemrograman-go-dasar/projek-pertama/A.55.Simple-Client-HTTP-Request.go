package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"net/url"
)

var baseURL = "http://localhost:8080"

type student struct {
    ID    string
    Name  string
    Grade int
}

func main() {
	// A.55.1. Penggunaan HTTP Request
    fmt.Println("# - A.55.1. Penggunaan HTTP Request");
	var users, errExOne = fetchUsers();
    if errExOne != nil {
        fmt.Println("Error!", errExOne.Error());
        return;
    }

    for _, each := range users {
        fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade);
    }

	// A.55.2. HTTP Request Dengan Form Data
	fmt.Println("# - A.55.2. HTTP Request Dengan Form Data");
    var user1, errExTwo = fetchUser("E001");
    if errExTwo != nil {
        fmt.Println("Error!", errExTwo.Error());
        return;
    }

    fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade);
}

func fetchUsers() ([]student, error) {
    var err error
    var client = &http.Client{}
    var data []student

    request, err := http.NewRequest("POST", baseURL+"/users", nil)
	// fmt.Println(err != nil);
    if err != nil {
        return nil, err
    }

    response, err := client.Do(request)
	// fmt.Println(response);
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    err = json.NewDecoder(response.Body).Decode(&data)
	// fmt.Println(data);
    if err != nil {
        return nil, err
    }

    return data,nil
}

func fetchUser(ID string) (student, error) {
    var err error
    var client = &http.Client{}
    var data student

    var param = url.Values{}
    param.Set("id", ID)
    var payload = bytes.NewBufferString(param.Encode())

    request, err := http.NewRequest("POST", baseURL+"/user", payload)
    if err != nil {
        return data, err
    }
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    response, err := client.Do(request)
    if err != nil {
        return data, err
    }
    defer response.Body.Close()

    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        return data, err
    }

    return data, nil
}
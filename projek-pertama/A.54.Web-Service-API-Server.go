package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type student struct {
    ID    string
    Name  string
    Grade int
}

var data = []student{
    student{"E001", "ethan", 21},
    student{"W001", "wick", 22},
    student{"B001", "bourne", 23},
    student{"B002", "bond", 23},
}

func main() {
    http.HandleFunc("/users", users);
    http.HandleFunc("/user", user);

    fmt.Println("starting web server at http://localhost:8080/");
    http.ListenAndServe(":8080", nil);
}

func users(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json");

    if r.Method == "GET" {
        var result, err = json.Marshal(data);
		fmt.Println(json.Marshal(data));

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError);
            return;
        }

        w.Write(result);
        return;
    }

    http.Error(w, "", http.StatusBadRequest);
}

func user(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json");
	// fmt.Println(r);

    if r.Method == "GET" {
        var id = r.FormValue("E001");
        var result []byte;
        var err error;
		fmt.Println(id);

        for _, each := range data {
			// fmt.Println(each.ID);
            if each.ID == id {
                result, err = json.Marshal(each);

                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError);
                    return;
                }

                w.Write(result);
                return;
            }
        }

        http.Error(w, "User not found", http.StatusNotFound);
        return;
    }

    http.Error(w, "", http.StatusBadRequest);
}
package main

import (
	"fmt"
    "strings"
    // "reflect"
)

func main() {
    // A.28.1. Penggunaan interface{}
    fmt.Println("# - A.28.1. Penggunaan interface{}");
        var secret interface{};
        secret = "ethan hunt";
        fmt.Println(secret);
        secret = []string{"apple", "manggo", "banana"};
        // fmt.Println(reflect.TypeOf(secret));
        fmt.Println(secret);
        secret = 12.4;
        fmt.Println(secret);
        fmt.Println("----------------------------------------");
        var data map[string]interface{};
        data = map[string]interface{} {
            "breakfast" : []string{"apple", "manggo", "banana"},
            "grade"     : 2,
            "name"      : "ethan hunt",
        };
        fmt.Println(data);
        fmt.Println();

    // A.28.2. Type Alias Any
    fmt.Println("# - A.28.2. Type Alias Any");
        // TYPE ANY, HANYA BISA DIGUNAKAN PADA GOLANG VERSI v1.18
        var exampleTwo map[string]any
        exampleTwo = map[string]any{
            "name":      "ethan hunt",
            "grade":     2,
            "breakfast": []string{"apple", "manggo", "banana"},
        }
        fmt.Println(exampleTwo);
        fmt.Println();

    // A.28.3. Casting Variabel Interface Kosong
    fmt.Println("# - A.28.3. Casting Variabel Interface Kosong");
        var exampleThree interface{};

        exampleThree = 2;
        var number = exampleThree.(int) * 10;
        fmt.Println(exampleThree, "multiplied by 10 is :", number);

        exampleThree = []string{"apple", "manggo", "banana"};
        var gruits = strings.Join(exampleThree.([]string), ", ");
        fmt.Println(gruits, "is my favorite fruits");

        exampleThree = "DEDE SYARIFUDIN";
        var result = exampleThree.(string) +" HIDAYAT";
        fmt.Println(result);
        fmt.Println();

    // A.28.4. Casting Variabel Interface Kosong Ke Objek Pointer
    fmt.Println("# - A.28.4. Casting Variabel Interface Kosong Ke Objek Pointer");
        type structExampleFour struct {
            name string
            age  int
        };
        var exampleFour interface{} = &structExampleFour{name: "wick", age: 27};
        var name = exampleFour.(*structExampleFour).name;
        age := exampleFour.(*structExampleFour).age;
        fmt.Println("NAMA : "+name);
        fmt.Println("AGE  : ",age);
        fmt.Println();

    // A.28.5. Kombinasi Slice, map, dan interface{}
    fmt.Println("# - A.28.5. Kombinasi Slice, map, dan interface{}");
        fmt.Println("----------------------------------------");
        var exampleFiveSecOne = []map[string]interface{} {
            {"name": "Wick", "age": 23},
            {"name": "Ethan", "age": 23},
            {"name": "Bourne", "age": 22},
        }
        for _, each := range exampleFiveSecOne {
            fmt.Println(each["name"], "age is", each["age"])
        }

        fmt.Println("----------------------------------------");
        var exampleFiveSecTwo = []map[any]any {
            {"name": "Wick", "age": 23},
            {"name": "Ethan", "age": 23},
            {"name": "Bourne", "age": 22},
        }
        for _, each := range exampleFiveSecTwo {
            fmt.Println(each["name"], "age is", each["age"])
        }

        fmt.Println("----------------------------------------");
        var fruitsOne = [] interface{} {
            map[ string ] interface{} {"name": "strawberry", "total": 10},
            [] string {"manggo", "pineapple", "papaya"},
            "orange",
        }
        for _, each := range fruitsOne {
            fmt.Println(each)
        }

        fmt.Println("----------------------------------------");
        var fruitsTwo = [] any {
            map[ string ] any {"name": "strawberry", "total": 10},
            [] string {"manggo", "pineapple", "papaya"},
            "orange",
        }
        for _, each := range fruitsTwo {
            fmt.Println(each)
        }

        fmt.Println("----------------------------------------");
        var fruitsThree = [] interface{} {
            map[ interface{} ] interface{} {"name": "strawberry", "total": 10},
            [] interface{} {"manggo", "pineapple", "papaya"},
            "orange",
        }
        for _, each := range fruitsThree {
            fmt.Println(each)
        }

        fmt.Println("----------------------------------------");
        var fruitsFour = [] any {
            map[ any ] any {"name": "strawberry", "total": 10},
            [] any {"manggo", "pineapple", "papaya"},
            "orange",
        }
        for _, each := range fruitsFour {
            fmt.Println(each)
        }
        fmt.Println();
}
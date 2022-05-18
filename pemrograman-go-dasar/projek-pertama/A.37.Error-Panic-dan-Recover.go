package main

import (
	"fmt"
    "strconv"
    "errors"
    "strings"
)

func main() {
	// A.37.1. Pemanfaatan Error
	fmt.Println("# - A.37.1. Pemanfaatan Error");
    var input string;
    fmt.Print("Type some number: ");
    fmt.Scanln(&input);
	
    var number int;
    var err error;
    number, err = strconv.Atoi(input);
	// fmt.Println(err.Error());
    if err == nil {
        fmt.Println(number, "is number");
    } else {
        fmt.Println(input, "is not number");
        fmt.Println(err.Error());
    }

	// A.37.2. Membuat Custom Error
	fmt.Println("# - A.37.2. Membuat Custom Error");
	var nameTwo string;
    fmt.Print("Type your name: ");
    fmt.Scanln(&nameTwo);

    if err, valid := validate(nameTwo); !valid {
		fmt.Println(err.Error());
	} else {
		fmt.Println(validate(nameTwo));
        fmt.Println("halo", nameTwo);
    }

	// A.37.3. Penggunaan panic
	fmt.Println("# - A.37.3. Penggunaan panic");
    var nameThree string;
    fmt.Print("Type your name: ");
    fmt.Scanln(&nameThree);

    if err, valid := validate(nameThree); !valid {
        panic(err.Error());
        fmt.Println("end");
	} else {
        fmt.Println("halo", nameThree);
    }

	// A.37.4. Penggunaan recover
	fmt.Println("# - A.37.4. Penggunaan recover");
    defer catch();

    var name string;
    fmt.Print("Type your name: ");
    fmt.Scanln(&name);

    if err, valid := validate(name); !valid {
        panic(err.Error());
        fmt.Println("end");
	} else {
        fmt.Println("halo", name);
    }

	// A.37.5. Pemanfaatan recover pada IIFE
	fmt.Println("# - A.37.5. Pemanfaatan recover pada IIFE");
	defer fmt.Println("INI DIJALANKAN AKHIR");
	// runErrorPanic();
	// defer func() {
    //     if err := recover(); err != nil {
    //         fmt.Println("Panic occured", err);
    //     } else {
    //         fmt.Println("Application running perfectly");
    //     }
    // }();
    // panic("some error happen");
	fmt.Println("----------------------------------------");

    data := []string{"superman", "aquaman", "wonder woman"};
    for _, each := range data {
        func() {
            // recover untuk IIFE dalam perulangan
            defer func() {
                if err := recover(); err != nil {
                    fmt.Println("Panic occured on looping", each, "| message:", err);
                } else {
                    fmt.Println("APLIKASI BERJALAN DENGAN SEMPURNA");
                }
            }()
            panic("BEBERAPA ERROR TERJADI");
        }();
    }
}

func validate(input string) (error, bool) {
    if strings.TrimSpace(input) == "" {
        return errors.New("cannot be empty"), false;
    }
    return nil, true;
}

func catch() {
	// fmt.Println("RECOVER : ",recover());
    if err := recover(); err != nil {
        fmt.Println("Error occured", err);
    } else {
        fmt.Println("APLIKASI BERJALAN DENGAN SEMPURNA, DIJALANKAN DARI FUNGSI CATCH");
    }
}

// func runErrorPanic() {
//     defer func() {
//         if err := recover(); err != nil {
//             fmt.Println("Panic occured", err);
//         } else {
//             fmt.Println("Application running perfectly");
//         }
//     }();
//     panic("some error happen");
// }
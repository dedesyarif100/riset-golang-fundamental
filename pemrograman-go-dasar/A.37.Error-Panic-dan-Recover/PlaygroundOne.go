package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("# - A.37.1. Pemanfaatan Error");
		var InputOne string
		fmt.Print("TYPE SOME NUMBER : ")
		fmt.Scanln(&InputOne)

		var number int
		var err error
		number, err = strconv.Atoi(InputOne)
		if err == nil {
			fmt.Println(number, "IS NUMBER")
		} else {
			fmt.Println(InputOne, "IS NOT NUMBER")
			fmt.Println(err.Error())
		}
		println()


	fmt.Println("# - A.37.2. Membuat Custom Error");
		var InputTwo string
		fmt.Print("TYPE YOUR NAME : ")
		fmt.Scanln(&InputTwo)

		if err, valid := Validate(InputTwo); !valid {
			fmt.Println(err.Error())
		} else {
			fmt.Println(Validate(InputTwo))
			fmt.Println("HALLO :", InputTwo)
		}
		println()


	fmt.Println("# - A.37.3. Penggunaan panic");
		var InputThree string
		fmt.Print("TYPE YOUR NAME : ")
		fmt.Scanln(&InputThree)

		if err, valid := Validate(InputThree); !valid {
			panic(err.Error())
            fmt.Println("end");
		} else {
			fmt.Println("HALLO :", InputThree)
		}
		println()


	fmt.Println("# - A.37.4. Penggunaan recover");
		var InputFour string
		fmt.Print("TYPE YOUR NAME : ")
		fmt.Scanln(&InputFour)

		defer Catch()
		if err, valid := Validate(InputFour); !valid {
			panic(err.Error())
            fmt.Println("end");
		} else {
			fmt.Println("HALLO :", InputFour)
		}
		println()


	fmt.Println("# - A.37.5. Pemanfaatan recover pada IIFE");
		defer fmt.Println("6. INI DIJALANKAN AKHIR");


		fmt.Println("----------------------------------------");
		data := []string{"superman", "aquaman", "wonder woman"};
		for _, each := range data {
			func() {
				// recover untuk IIFE dalam perulangan
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err);
						fmt.Println("1. Panic occured on looping", each, "| message:", err);
					} else {
						fmt.Println("2. APLIKASI BERJALAN DENGAN SEMPURNA");
					}
				}()
				panic("3. BEBERAPA ERROR TERJADI");
			}();
		}
		fmt.Println();
}

func Validate(input string) (error, bool) {
	if strings.TrimSpace(input) == "" {
		return errors.New("CANNOT BE EMPTY"), false
	}
	return nil, true
}

func Catch() {
	if err := recover(); err != nil {
		fmt.Println("ERROR OCCURED |", err)
	} else {
        fmt.Println("APLIKASI BERJALAN DENGAN SEMPURNA, DIJALANKAN DARI FUNGSI CATCH");
	}
}
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// A.43.1. Konversi Menggunakan strconv
	fmt.Println("# - A.43.1. Konversi Menggunakan strconv");
		// • Fungsi strconv.Atoi()
		fmt.Println("• Fungsi strconv.Atoi()");
			var strExOneSecOne = "124"
			fmt.Println("STRING TO INT :", reflect.TypeOf(strExOneSecOne))
			var numExOneSecOne, errExOneSecOne = strconv.Atoi(strExOneSecOne)
			// fmt.Println(strconv.Atoi(strExOneSecOne))
			if errExOneSecOne == nil {
				fmt.Println("STRING TO INT :", reflect.TypeOf(numExOneSecOne)) // 124
			}
			fmt.Println();

		// • Fungsi strconv.Itoa()
		fmt.Println("• Fungsi strconv.Itoa()");
			var numExOneSecTwo = 124
			fmt.Println("INT TO STRING :", reflect.TypeOf(numExOneSecOne))
			var strExOneSecTwo = strconv.Itoa(numExOneSecTwo)
			fmt.Println("INT TO STRING :", reflect.TypeOf(strExOneSecTwo)) // "124"
			fmt.Println();

		// • Fungsi strconv.ParseInt()
		fmt.Println("• Fungsi strconv.ParseInt()");
			var strExOneSecThree = "124"
			fmt.Println("STRING TO INT64 :", reflect.TypeOf(strExOneSecThree))
			var numExOneSecThree, errExOneSecThree = strconv.ParseInt(strExOneSecThree, 10, 64)
			if errExOneSecThree == nil {
				fmt.Println("STRING TO INT64 :", reflect.TypeOf(numExOneSecThree)) // 124
			}
			fmt.Println("-----------------------")
			var str = "1010"
			fmt.Println("STRING TO INT64 :", reflect.TypeOf(str))
			var num, err = strconv.ParseInt(str, 2, 8)
			if err == nil {
				fmt.Println("STRING TO INT64 :", reflect.TypeOf(num)) // 10
			}
			fmt.Println();

		// • Fungsi strconv.FormatInt()
		fmt.Println("• Fungsi strconv.FormatInt()");
			var numExOneSecFour = int64(24)
			fmt.Println("INT64 TO STRING :", reflect.TypeOf(numExOneSecFour))
			var strExOneSecFour = strconv.FormatInt(numExOneSecFour, 10)
			fmt.Println("INT64 TO STRING :", reflect.TypeOf(strExOneSecFour)) // 24
			fmt.Println();

		// • Fungsi strconv.ParseFloat()
		fmt.Println("• Fungsi strconv.ParseFloat()");
			var strExOneSecFive = "24.12"
			fmt.Println("STRING TO FLOAT64 :", reflect.TypeOf(strExOneSecFive))
			var numExOneSecFive, errExOneSecFive = strconv.ParseFloat(strExOneSecFive, 32)
			if errExOneSecFive == nil {
				fmt.Println("STRING TO FLOAT64 :", reflect.TypeOf(numExOneSecFive)) // 24.1200008392334
			}
			fmt.Println();

		// • Fungsi strconv.FormatFloat()
		fmt.Println("• Fungsi strconv.FormatFloat()");
			var numExOneSecSix = float64(24.12)
			fmt.Println("FLOAT64 TO STRING :", reflect.TypeOf(numExOneSecSix))
			var strExOneSecSix = strconv.FormatFloat(numExOneSecSix, 'f', 6, 64)
			fmt.Println("FLOAT64 TO STRING :", reflect.TypeOf(strExOneSecSix)) // 24.120000
			fmt.Println();

		// • Fungsi strconv.ParseBool()
		fmt.Println("• Fungsi strconv.ParseBool()");
			var strExOneSecSeven = "true"
			fmt.Println("STRING TO BOOL :", reflect.TypeOf(strExOneSecSeven))
			var bulExOneSecSeven, errExOneSecSeven = strconv.ParseBool(strExOneSecSeven)
			if errExOneSecSeven == nil {
				fmt.Println("STRING TO BOOL :", reflect.TypeOf(bulExOneSecSeven)) // true
			}
			fmt.Println();

		// • Fungsi strconv.FormatBool()
		fmt.Println("• Fungsi strconv.FormatBool()");
			var bulExOneSecEight = true
			fmt.Println("BOOL TO STRING :", reflect.TypeOf(bulExOneSecEight))
			var strExOneSecEight = strconv.FormatBool(bulExOneSecEight)
			fmt.Println("BOOL TO STRING :", reflect.TypeOf(strExOneSecEight)) // true
			fmt.Println();

	fmt.Println("----------------------------------------------------")
	// A.43.2. Konversi Data Menggunakan Teknik Casting
	fmt.Println("# - A.43.2. Konversi Data Menggunakan Teknik Casting");
		var castingOne float64 = float64(24);
		fmt.Println("FLOAT64 TO INT32 :", reflect.TypeOf(castingOne)); // 24
		var castingTwo int32 = int32(24.00);
		fmt.Println("FLOAT64 TO INT32 :", reflect.TypeOf(castingTwo)); // 24
		fmt.Println();

	// A.43.3. Casting string ↔ byte
	fmt.Println("# - A.43.3. Casting string ↔ byte");
		var text1 = "halo";
		fmt.Println("STRING TO []UINT8 :", reflect.TypeOf(text1))
		var b = []byte(text1);
		fmt.Println("STRING TO []UINT8 :", reflect.TypeOf(b))
		fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3]);
		// 104 97 108 111
		fmt.Println("----------------")

		var byte1 = []byte{104, 97, 108, 111}
		fmt.Println("[]UINT8 TO STRING :", reflect.TypeOf(byte1))
		var s = string(byte1)
		fmt.Println("[]UINT8 TO STRING :", reflect.TypeOf(s))
		fmt.Printf("%s \n", s)
		// halo
		fmt.Println("----------------")

		var c int64 = int64('h')
		fmt.Println("INT64 	:", reflect.TypeOf(c))
		fmt.Println(c) // 104
		fmt.Println("----------------")

		var d string = string([]byte{104})
		fmt.Println("STRING 	:", reflect.TypeOf(d))
		fmt.Println(d) // h
		fmt.Println();

	// A.43.4. Type Assertions Pada Interface Kosong (interface{})
	fmt.Println("# - A.43.4. Type Assertions Pada Interface Kosong (interface{})");
		var data = map[string]interface{}{
			"nama":    "john wick",
			"grade":   2,
			"height":  156.5,
			"isMale":  true,
			"hobbies": []string{"eating", "sleeping"},
		}

		fmt.Println("NAMA	:", data["nama"].(string));
		fmt.Println("GRADE	:", data["grade"].(int));
		fmt.Println("HEIGHT	:", data["height"].(float64));
		fmt.Println("ISMALE	:", data["isMale"].(bool));
		fmt.Println("HOBBIES	:", data["hobbies"].([]string));

		fmt.Println("---------------------------------------");
		for _, val := range data {
			switch val.(type) {
				case string:
					fmt.Println("NAMA	:", val.(string));
				case int:
					fmt.Println("GRADE	:", val.(int));
				case float64:
					fmt.Println("HEIGHT	:", val.(float64));
				case bool:
					fmt.Println("ISMALE	:", val.(bool));
				case []string:
					fmt.Println("HOBBY	:", val.([]string));
				default:
					fmt.Println("DEFAULT	:",val.(int));
			}
		}
		println()
}
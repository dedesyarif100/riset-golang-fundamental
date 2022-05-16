package main

import (
    "fmt"
	"strconv"
)

func main() {
	// A.43.1. Konversi Menggunakan strconv
	fmt.Println("# - A.43.1. Konversi Menggunakan strconv");
		// • Fungsi strconv.Atoi()
		fmt.Println("• Fungsi strconv.Atoi()");
		var strExOneSecOne = "124"
		var numExOneSecOne, errExOneSecOne = strconv.Atoi(strExOneSecOne)
		// fmt.Println(strconv.Atoi(strExOneSecOne))
		if errExOneSecOne == nil {
			fmt.Println(numExOneSecOne) // 124
		}

		// • Fungsi strconv.Itoa()
		fmt.Println("• Fungsi strconv.Itoa()");
		var numExOneSecTwo = 124
		var strExOneSecTwo = strconv.Itoa(numExOneSecTwo)
		fmt.Println(strExOneSecTwo) // "124"

		// • Fungsi strconv.ParseInt()
		fmt.Println("• Fungsi strconv.ParseInt()");
		var strExOneSecThree = "124"
		var numExOneSecThree, errExOneSecThree = strconv.ParseInt(strExOneSecThree, 10, 64)
		if errExOneSecThree == nil {
			fmt.Println(numExOneSecThree) // 124
		}
		var str = "1010"
		var num, err = strconv.ParseInt(str, 2, 8)
		if err == nil {
			fmt.Println(num) // 10
		}

		// • Fungsi strconv.FormatInt()
		fmt.Println("• Fungsi strconv.FormatInt()");
		var numExOneSecFour = int64(24)
		var strExOneSecFour = strconv.FormatInt(numExOneSecFour, 8)
		fmt.Println(strExOneSecFour) // 30

		// • Fungsi strconv.ParseFloat()
		fmt.Println("• Fungsi strconv.ParseFloat()");
		var strExOneSecFive = "24.12"
		var numExOneSecFive, errExOneSecFive = strconv.ParseFloat(strExOneSecFive, 32)
		if errExOneSecFive == nil {
			fmt.Println(numExOneSecFive) // 24.1200008392334
		}

		// • Fungsi strconv.FormatFloat()
		fmt.Println("• Fungsi strconv.FormatFloat()");
		var numExOneSecSix = float64(24.12)
		var strExOneSecSix = strconv.FormatFloat(numExOneSecSix, 'f', 6, 64)
		fmt.Println(strExOneSecSix) // 24.120000

		// • Fungsi strconv.ParseBool()
		fmt.Println("• Fungsi strconv.ParseBool()");
		var strExOneSecSeven = "true"
		var bulExOneSecSeven, errExOneSecSeven = strconv.ParseBool(strExOneSecSeven)
		if errExOneSecSeven == nil {
			fmt.Println(bulExOneSecSeven) // true
		}

		// • Fungsi strconv.FormatBool()
		fmt.Println("• Fungsi strconv.FormatBool()");
		var bulExOneSecEight = true
		var strExOneSecEight = strconv.FormatBool(bulExOneSecEight)
		fmt.Println(strExOneSecEight) // true

	// A.43.2. Konversi Data Menggunakan Teknik Casting
	fmt.Println("# - A.43.2. Konversi Data Menggunakan Teknik Casting");
		var castingOne float64 = float64(24);
		fmt.Println(castingOne); // 24
		var castingTwo int32 = int32(24.00);
		fmt.Println(castingTwo); // 24

	// A.43.3. Casting string ↔ byte
	fmt.Println("# - A.43.3. Casting string ↔ byte");
		var text1 = "halo";
		var b = []byte(text1);
		fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3]);
		// 104 97 108 111

		var byte1 = []byte{104, 97, 108, 111}
		var s = string(byte1)
		fmt.Printf("%s \n", s)
		// halo

		var c int64 = int64('h')
		fmt.Println(c) // 104

		var d string = string(104)
		fmt.Println(d) // h

	// A.43.4. Type Assertions Pada Interface Kosong (interface{})
	fmt.Println("# - A.43.4. Type Assertions Pada Interface Kosong (interface{})");
		var data = map[string]interface{}{
			"nama":    "john wick",
			"grade":   2,
			"height":  156.5,
			"isMale":  true,
			"hobbies": []string{"eating", "sleeping"},
		}

		fmt.Println(data["nama"].(string));
		fmt.Println(data["grade"].(int));
		fmt.Println(data["height"].(float64));
		fmt.Println(data["isMale"].(bool));
		fmt.Println(data["hobbies"].([]string));

		fmt.Println("---------------------------------------");
		for _, val := range data {
			switch val.(type) {
				case string:
					fmt.Println(val.(string));
				case int:
					fmt.Println(val.(int));
				case float64:
					fmt.Println(val.(float64));
				case bool:
					fmt.Println(val.(bool));
				case []string:
					fmt.Println(val.([]string));
				default:
					fmt.Println(val.(int));
			}
		}
}
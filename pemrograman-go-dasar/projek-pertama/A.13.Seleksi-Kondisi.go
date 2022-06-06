package main

import (
	"fmt"
	"reflect"
)

func main() {
	// A.13.1. Seleksi Kondisi Menggunakan Keyword if, else if, & else
	fmt.Println("# - A.13.1. Seleksi Kondisi Menggunakan Keyword if, else if, & else");
		var pointOne = 8;
		if pointOne == 10 {
			fmt.Println("lulus dengan nilai sempurna");
		} else if pointOne > 5 {
			fmt.Println("lulus");
		} else if pointOne == 4 {
			fmt.Println("hampir lulus");
		} else {
			fmt.Printf("tidak lulus. nilai anda %d\n", pointOne);
		}
		println()

	// A.13.2. Variabel Temporary Pada if - else
	fmt.Println("# - A.13.2. Variabel Temporary Pada if - else");
		var pointTwo = 8840.0
		if percent := pointTwo / 100; percent >= 100 {
			fmt.Printf("%.1f%s perfect!\n", percent, "%")
		} else if percent >= 70 {
			fmt.Printf("%.1f%s good\n", percent, "%")
		} else {
			fmt.Printf("%.1f%s not bad\n", percent, "%")
		}
		
		// fmt.Println(reflect.Slice)
		var tesOne = 2
		if tesTwo := tesOne; tesOne != 2 {
			fmt.Printf("VAL TES TWO : %d", tesTwo)
		} else if tesThree := 3; tesThree != 3 {
			fmt.Printf("VAL TES THREE : %d", tesThree)
		} else if tesFour := 4; tesFour == 4 && tesFour > 10 {
			fmt.Printf("VAL TES FOUR : %d", tesFour)
		} else if tesFive := []int{1, 2, 3}; reflect.TypeOf(tesFive).Kind() == reflect.Slice {
			fmt.Println("VAL TES FIVE :", tesFive)
			fmt.Printf("TYPE OF : %T\n", tesFive)
		}
		println()

		if tesSix := []any{1, "ONE"}; reflect.TypeOf(tesSix).Kind() == reflect.Slice {
			fmt.Println("VAL TES SIX :", tesSix)
			fmt.Printf("TYPE OF : %T\n", tesSix)
			fmt.Printf("TYPE OF : %v\n\n", reflect.TypeOf(tesSix))
		}

	// A.13.3. Seleksi Kondisi Menggunakan Keyword switch - case
	fmt.Println("# - A.13.3. Seleksi Kondisi Menggunakan Keyword switch - case");
		var pointThree = 7
		switch pointThree {
			case 8:
				fmt.Println("perfect")
			case 7:
				fmt.Println("awesome")
			default:
				fmt.Println("not bad")
		}
		println()

	// A.13.4. Pemanfaatan case Untuk Banyak Kondisi
	fmt.Println("# - A.13.4. Pemanfaatan case Untuk Banyak Kondisi");
		var pointFour = 10
		switch pointFour {
			// case tmp > 8:
			// 	fmt.Println("perfect")
			case 8:
				fmt.Println("perfect")
			case 7, 6, 5, 4:
				fmt.Println("awesome")
			default:
				fmt.Println("not bad")
		}
		println()

	// A.13.5. Kurung Kurawal Pada Keyword case & default
	fmt.Println("# - A.13.5. Kurung Kurawal Pada Keyword case & default");
		var point5 = 3
		switch point5 {
			case 8:
				fmt.Println("perfect")
			case 7, 6, 5, 4:
				fmt.Println("awesome")
			default:
				{
					fmt.Println("not bad")
					fmt.Println("you can be better!")
				}
		}
		println()

	// A.13.6. Switch Dengan Gaya if - else
	fmt.Println("# - A.13.6. Switch Dengan Gaya if - else");
		var pointSix = 6
		switch {
			case pointSix == 8:
				fmt.Println("perfect")
			case (pointSix < 8) && (pointSix > 3):
				fmt.Println("awesome")
			default:
				{
					fmt.Println("not bad")
					fmt.Println("you need to learn more")
				}
		}
		println()

	// A.13.7. Penggunaan Keyword fallthrough Dalam switch
	fmt.Println("# - A.13.7. Penggunaan Keyword fallthrough Dalam switch");
		var pointSeven = 6
		switch {
			case pointSeven == 8:
				fmt.Println("perfect")
			case (pointSeven < 8) && (pointSeven > 3):
				fmt.Println("awesome")
				fallthrough
			case pointSeven < 5:
				fmt.Println("you need to learn more")
			default:
				{
					fmt.Println("not bad")
					fmt.Println("you need to learn more")
				}
		}
		println()

	// A.13.8. Seleksi Kondisi Bersarang
	fmt.Println("# - A.13.8. Seleksi Kondisi Bersarang");
		var pointEigth = 10
		if pointEigth > 7 {
			switch pointEigth {
				case 10:
					fmt.Println("perfect!")
				default:
					fmt.Println("nice!")
			}
		} else {
			if pointEigth == 5 {
				fmt.Println("not bad")
			} else if pointEigth == 3 {
				fmt.Println("keep trying")
			} else {
				fmt.Println("you can do it")
				if pointEigth == 0 {
					fmt.Println("try harder!")
				}
			}
		}
		println()
}
package main

import "fmt"

func main() {
	// A.13.1. Seleksi Kondisi Menggunakan Keyword if, else if, & else
	fmt.Println("A.13.1. Seleksi Kondisi Menggunakan Keyword if, else if, & else");
	var pointOne = 8
	if pointOne == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if pointOne > 5 {
		fmt.Println("lulus")
	} else if pointOne == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", pointOne)
	}

	// A.13.2. Variabel Temporary Pada if - else
	fmt.Println("A.13.2. Variabel Temporary Pada if - else");
	var pointTwo = 8840.0
	if percent := pointTwo / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// A.13.3. Seleksi Kondisi Menggunakan Keyword switch - case
	fmt.Println("A.13.3. Seleksi Kondisi Menggunakan Keyword switch - case");
	var pointThree = 7
	switch pointThree {
		case 8:
			fmt.Println("perfect")
		case 7:
			fmt.Println("awesome")
		default:
			fmt.Println("not bad")
	}

	// A.13.4. Pemanfaatan case Untuk Banyak Kondisi
	fmt.Println("A.13.4. Pemanfaatan case Untuk Banyak Kondisi");
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

	// A.13.5. Kurung Kurawal Pada Keyword case & default
	fmt.Println("A.13.5. Kurung Kurawal Pada Keyword case & default");
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

	// A.13.6. Switch Dengan Gaya if - else
	fmt.Println("A.13.6. Switch Dengan Gaya if - else");
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

	// A.13.7. Penggunaan Keyword fallthrough Dalam switch
	fmt.Println("A.13.7. Penggunaan Keyword fallthrough Dalam switch");
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

	// A.13.8. Seleksi Kondisi Bersarang
	fmt.Println("A.13.8. Seleksi Kondisi Bersarang");
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
}
package main

import "fmt"

func main() {
	// A.14.1. Perulangan Menggunakan Keyword for
	fmt.Println("A.14.1. Perulangan Menggunakan Keyword for");
	for i := 0; i < 12; i++ {
		fmt.Println("Angka", i+1)
	}
	fmt.Println("----------------------")
	for j := 5; j > 0; j-- {
		fmt.Println("Angka", j)
	}

	// A.14.2. Penggunaan Keyword for Dengan Argumen Hanya Kondisi
	fmt.Println("A.14.2. Penggunaan Keyword for Dengan Argumen Hanya Kondisi");
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i + 1)
		i++
	}
	fmt.Println("----------------------")
	j := 5
	for j > 0 {
		fmt.Println("Angka", j)
		j--
	}

	// A.14.3. Penggunaan Keyword for Tanpa
	fmt.Println("A.14.3. Penggunaan Keyword for Tanpa Argumen");
	var loopingThree = 0
	for {
		fmt.Println("Angka", loopingThree + 1)

		loopingThree++
		if loopingThree == 8 {
			break
		}
	}

	// A.14.4. Penggunaan Keyword for - range
	fmt.Println("A.14.4. Penggunaan Keyword for - range");

	// A.14.5. Penggunaan Keyword break & 
	fmt.Println("A.14.5. Penggunaan Keyword break & continue");
	for loopingFive := 1; loopingFive <= 10; loopingFive++ {
		if loopingFive % 2 == 1 {
			continue
		}
		if loopingFive > 8 {
			break
		}
		fmt.Println("Angka", loopingFive)
	}

	// A.14.6. Perulangan Bersarang
	fmt.Println("A.14.6. Perulangan Bersarang");
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// A.14.7. Pemanfaatan Label Dalam Perulangan
	fmt.Println("A.14.7. Pemanfaatan Label Dalam Perulangan");
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
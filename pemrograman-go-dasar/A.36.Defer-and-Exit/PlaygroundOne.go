package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("# - A.36.1. Penerapan keyword defer");
		defer fmt.Println("INI DIJALANKAN PALING AKHIR")
		fmt.Println("SELAMAT DATANG")
		ExampleOne("pizza")
		println()

	fmt.Println("# - A.36.2. Kombinasi defer dan IIFE");
		numberTwoSecOne := 3
		if numberTwoSecOne == 3 {
			fmt.Println("HALLO 1")
			defer fmt.Println("HALLO 3")
		}
		fmt.Println("HALLO 2")
		// ------------------------------------------------
		numberTwoSecTwo := 3
		if numberTwoSecTwo == 3 {
			fmt.Println("HAI 1")
			func ()  {
				defer fmt.Println("HAI 3 -----------------")
			}()
		}
		fmt.Println("HAI 2")
		println()

	fmt.Println("# - A.36.3. Penerapan Fungsi os.Exit()");
		defer fmt.Println("TES 1")
		os.Exit(1)
}

func ExampleOne(menu string) {
	defer fmt.Println("TERIMA KASIH, SILAHKAN TUNGGU")
	if menu == "pizza" {
		fmt.Println("PILIHAN TEPAT !!!")
		return
	}
	fmt.Println("PESANAN ANDA :", menu)
}
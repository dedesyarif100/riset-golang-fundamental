package main

import "fmt"

func main() {
	// fmt.Println("COBA TIPE DATA")
	// var a float32 = 12
	// fmt.Println("value A : ", a)

	// A.10.1. Tipe Data Numerik Non-Desimal
	fmt.Println("A.10.1. Tipe Data Numerik Non-Desimal");
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	firstName := "DEDE"
	lastName := "SYARIF"

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	fmt.Printf("hallo %s %s!\n", firstName, lastName)

	// A.10.2. Tipe Data Numerik Desimal
	fmt.Println("A.10.2. Tipe Data Numerik Desimal");
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// A.10.3. Tipe Data bool (Boolean)
	fmt.Println("A.10.3. Tipe Data bool (Boolean)");
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// A.10.4. Tipe Data string
	fmt.Println("A.10.4. Tipe Data string");
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	// A.10.5. Nilai nil & Zero Value
}
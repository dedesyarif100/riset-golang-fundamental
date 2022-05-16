package main

import (
	"fmt"
	"math"
)

func main() {
	// A.19.1 Penerapan Fungsi Multiple Return
	fmt.Println("# - A.19.1 Penerapan Fungsi Multiple Return");
		var diameter float64 = 15
		var area, circumference = calculate(diameter)
		fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
		fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

// A.19.1 Penerapan Fungsi Multiple Return
// func calculate(d float64) (float64, float64) {
//     // hitung luas
//     var area = math.Pi * math.Pow(d / 2, 2)
//     // hitung keliling
//     var circumference = math.Pi * d
//     // kembalikan 2 nilai
// 	fmt.Println(area)
// 	fmt.Println(circumference)
//     return area, circumference
// }

// A.19.2 Fungsi Dengan Predefined Return Value
func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d / 2, 2)
	circumference = math.Pi * d
	fmt.Println(area)
	fmt.Println(circumference)
	// math.
	return
}
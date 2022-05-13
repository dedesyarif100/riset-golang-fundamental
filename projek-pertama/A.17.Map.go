package main

import (
	"fmt"
)

func main() {
	// A.17.1. Penggunaan Map
	fmt.Println("A.17.1. Penggunaan Map");
	var exampleOne map[string]int
	exampleOne = map[string]int{}

	exampleOne["A"] = 10
	exampleOne["B"] = 20
	exampleOne["C"] = 30
	exampleOne["D"] = 40
	exampleOne["E"] = 50

	fmt.Println("A :", exampleOne["A"])
	fmt.Println("B :", exampleOne["B"])
	fmt.Println("C :", exampleOne["C"])
	fmt.Println("D :", exampleOne["D"])
	fmt.Println("E :", exampleOne["E"])

	// A.17.2. Inisialisasi Nilai Map
	fmt.Println("A.17.2. Inisialisasi Nilai Map");
	// CARA HORIZONTAL
	var exampleTwoSectionA = map[string]int{"A" : 10, "B" : 20}
	// CARA VERTICAL
	var exampleTwoSectionB = map[string]int{
		"A" : 10,
		"B" : 20,
	}
	fmt.Println("SECTION A : ", exampleTwoSectionA["A"])
	fmt.Println("SECTION A : ", exampleTwoSectionA["B"])
	fmt.Println("-------------------------------------")
	fmt.Println("SECTION B : ", exampleTwoSectionB["A"])
	fmt.Println("SECTION B : ", exampleTwoSectionB["B"])

	// KETIGA CARA INI UNTUK MENG-INISIALISASI MAP DENGAN CARA YANG SAMA
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)
	// KETIGA CARA INI UNTUK MENG-INISIALISASI MAP DENGAN CARA YANG SAMA

	// A.17.3. Iterasi Item Map Menggunakan for - range
	fmt.Println("A.17.3. Iterasi Item Map Menggunakan for - range");
	var exampleThree = map[string]int{
		"A" : 10,
		"B" : 20,
		"C" : 30,
		"D" : 40,
	}

	for key, val := range exampleThree {
		fmt.Println(key, ":\t", val)
	}

	// A.17.4. Menghapus Item Map
	fmt.Println("A.17.4. Menghapus Item Map");
	var exampleFour = map[string]int{"A":10, "B":20, "C":30}
	fmt.Println(len(exampleFour))
	fmt.Println(exampleFour)
	fmt.Println("-------------------------")
	delete(exampleFour, "A")
	fmt.Println(len(exampleFour))
	fmt.Println(exampleFour)

	// A.17.5. Deteksi Keberadaan Item Dengan Key Tertentu
	fmt.Println("A.17.5. Deteksi Keberadaan Item Dengan Key Tertentu");
	var exampleFive = map[string]int{"A":10, "B":20, "C":30}
	var value, isExist = exampleFive["D"]

	fmt.Println(isExist)
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}

	// A.17.6. Kombinasi Slice & Map
	fmt.Println("A.17.6. Kombinasi Slice & Map");
	var exampleSix = []map[string]string{
		map[string]string{"name": "chicken blue",   "gender": "male"},
		map[string]string{"name": "chicken red",    "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	// var exampleSix = []map[string]string{
	// 	map[string]string{"name": "chicken blue",   "gender": "male"},
	// 	map[string]string{"name": "chicken red",    "gender": "male"},
	// 	map[string]string{"name": "chicken yellow", "gender": "female"},
	// }
	for _, chicken := range exampleSix {
		fmt.Println(chicken["gender"] +" : "+ chicken["name"])
	}
	
	fmt.Println("-------------------------")
	var exampleSixSectionTwo = []map[int]string{
		{0: "A", 1: "B", 2: "C", 3: "D"},
		{0: "A", 1: "B", 2: "C", 3: "D"},
		{0: "A", 1: "B", 2: "C", 3: "D"},
		{0: "A", 1: "B", 2: "C", 3: "D"},
	}

	for _, val := range exampleSixSectionTwo {
		fmt.Println(val[0]+" : "+val[1]+" : "+val[2]+" : "+val[3])
	}
}
package main

import "fmt"

func main() {
	fmt.Println("# - A.15.Array.go");
		var names [4]string
		names[0] = "trafalgar"
		names[1] = "d"
		names[2] = "water"
		names[3] = "law"

		fmt.Println(names[0], names[1], names[2], names[3])
		println()

	// A.15.1. Pengisian Elemen Array yang Melebihi Alokasi Awal
	fmt.Println("# - A.15.1. Pengisian Elemen Array yang Melebihi Alokasi Awal");
		var exampleOne [4]string
		exampleOne[0] = "trafalgar"
		exampleOne[1] = "d"
		exampleOne[2] = "water"
		exampleOne[3] = "law"
		fmt.Println(exampleOne[0], exampleOne[1], exampleOne[2], exampleOne[3])
		// names[4] = "ez" // baris kode ini menghasilkan error
		println()

	// A.15.2. Inisialisasi Nilai Awal Array
	fmt.Println("# - A.15.2. Inisialisasi Nilai Awal Array");
		var exampleTwo = [4]string{"apple", "grape", "banana", "melon"}
		fmt.Println("Jumlah element \t\t :", len(exampleTwo))
		fmt.Println("Isi semua element \t :", exampleTwo)
		println()

	// A.15.3. Inisialisasi Nilai Array Dengan Gaya Vertikal
	fmt.Println("# - A.15.3. Inisialisasi Nilai Array Dengan Gaya Vertikal");
		var exampleThree [4]string
		// cara horizontal
		exampleThree  = [4]string{"apple", "grape", "banana", "melon"}
		// cara vertikal
		exampleThree  = [4]string{
			"apple",
			"grape",
			"banana",
			"melon",
		}
		fmt.Println("Isi semua element \t :", exampleThree)
		println()

	// A.15.4. Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	fmt.Println("# - A.15.4. Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen");
		var exampleFour = [...]int{2, 3, 2, 4, 3}
		fmt.Println("data array \t:", exampleFour)
		fmt.Println("jumlah elemen \t:", len(exampleFour))
		println()

	// A.15.5. Array Multidimensi
	fmt.Println("# - A.15.5. Array Multidimensi");
		var numbers1 = [2][3]int{
							[3]int{3, 2, 3}, 
							[3]int{3, 4, 5},
						}
		var numbers2 = [2][3]int{
							{3, 2, 3}, 
							{3, 4, 5},
						}
		exFiveSecOne := 2
		var numbers3 = [2][3]interface{}{
							{1, exFiveSecOne, 3},
							{"A", true, 1.2},
						}
		fmt.Println("numbers1 :", numbers1)
		fmt.Println("numbers2 :", numbers2)
		fmt.Println("numbers3 :", numbers3)
		println()

	// A.15.6. Perulangan Elemen Array Menggunakan Keyword for
	fmt.Println("# - A.15.6. Perulangan Elemen Array Menggunakan Keyword for");
		var exampleSix = [4]string{"apple", "grape", "banana", "melon"}
		for i := 0; i < len(exampleSix); i++ {
			fmt.Printf("elemen %d : %s\n", i, exampleSix[i])
		}
		println()

	// A.15.7. Perulangan Elemen Array Menggunakan Keyword for - range
	fmt.Println("# - A.15.7. Perulangan Elemen Array Menggunakan Keyword for - range");
		var exampleSeven = [4]string{"apple", "grape", "banana", "melon"}
		for key, seven := range exampleSeven {
			fmt.Println("element :", key, seven)
		}
		println()

	// A.15.8. Penggunaan Variabel Underscore _ Dalam for - range
	fmt.Println("# - A.15.8. Penggunaan Variabel Underscore _ Dalam for - range");
		var exampleEight = [4]string{"apple", "grape", "banana", "melon"}
		for _, eight := range exampleEight {
			fmt.Println("nama buah :", eight)
		}
		// for eight := range exampleEight {
		// 	fmt.Println("nama buah : ", eight)
		// }
		println()

	// A.15.9. Alokasi Elemen Array Menggunakan Keyword make
	fmt.Println("# - A.15.9. Alokasi Elemen Array Menggunakan Keyword make");
		var exampleNine = make([]string, 2)
		exampleNine[0] = "apple"
		exampleNine[1] = "manggo"
		fmt.Println(exampleNine)  // [apple manggo]
		println()
}
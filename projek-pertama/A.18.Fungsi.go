package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
)

func main() {
	// A.18.1. Penerapan Fungsi
	fmt.Println("A.18.1. Penerapan Fungsi");
	var names = []string{"DEDE", "HEND"}
	exampleOne("RIAN", names)

	// A.18.2. Fungsi Dengan Return Value / Nilai Balik
	fmt.Println("A.18.2. Fungsi Dengan Return Value / Nilai Balik");
    rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = exampleTwo(2, 10)
    fmt.Println("random number:", randomValue)
	randomValue = exampleTwo(2, 10)
    fmt.Println("random number:", randomValue)
	randomValue = exampleTwo(2, 10)
    fmt.Println("random number:", randomValue)

	// A.18.6. Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi
	fmt.Println("A.18.6. Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi");
	exampleSix(10, 2)
	exampleSix(4, 0)
	exampleSix(8, -4)
}

// A.18.1. Penerapan Fungsi
func exampleOne(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// A.18.2. Fungsi Dengan Return Value / Nilai Balik
func exampleTwo(min, max int) int {
	fmt.Println(rand.Int())
	var value = rand.Int() % (max - min + 1) + min
	return value
}

// A.18.6. Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi
func exampleSix(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
        return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
package main

import (
	"fmt"
    "math/rand"
	"time"
)

func main() {
	// A.39.2. Package math/rand
	fmt.Println("# - A.39.2. Package math/rand");
		rand.Seed(10)
		fmt.Println("random ke-1:", rand.Int()) // 5221277731205826435
		fmt.Println("random ke-2:", rand.Int()) // 3852159813000522384
		fmt.Println("random ke-3:", rand.Int()) // 8532807521486154107
		fmt.Println();

	// A.39.3. Unique Seed
	fmt.Println("# - A.39.3. Unique Seed");
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Println(rand.Int())
		fmt.Println(rand.Int())
		fmt.Println(rand.Int())
		fmt.Println();

	// A.39.4. Random Tipe Data Numerik Lainnya
	fmt.Println("# - A.39.4. Random Tipe Data Numerik Lainnya");
		fmt.Println(rand.Float64());
		fmt.Println(rand.Int31());
		fmt.Println(rand.NormFloat64());
		fmt.Println();

	// A.39.5. Angka Random Index Tertentu
	fmt.Println("# - A.39.5. Angka Random Index Tertentu");
		fmt.Println(rand.Intn(5));
		fmt.Println();

	// A.39.6. Random Tipe Data String
	fmt.Println("# - A.39.6. Random Tipe Data String");
		fmt.Println("RANDOM STRING :",randomString(10));
		fmt.Println();
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ");
var valA = []rune("1");

func randomString(length int) string {
	fmt.Println("CEK VAL LETTERS :",letters,"\n");
	fmt.Println("CEK VAL A :",valA,"\n");
    b := make([]rune, length)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
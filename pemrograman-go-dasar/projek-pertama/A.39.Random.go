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

	// A.39.3. Unique Seed
	fmt.Println("# - A.39.3. Unique Seed");
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())

	// A.39.4. Random Tipe Data Numerik Lainnya
	fmt.Println("# - A.39.4. Random Tipe Data Numerik Lainnya");
	// A.39.5. Angka Random Index Tertentu
	fmt.Println("# - A.39.5. Angka Random Index Tertentu");
	// A.39.6. Random Tipe Data String
	fmt.Println("# - A.39.6. Random Tipe Data String");
	randomString(10)
}
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ");

func randomString(length int) string {
    b := make([]rune, length)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
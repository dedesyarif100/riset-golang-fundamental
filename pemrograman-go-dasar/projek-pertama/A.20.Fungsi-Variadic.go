package main

import (
	"fmt"
	"strings"
)

func main() {
	// A.20.1. Penerapan Fungsi Variadic
	fmt.Println("# - A.20.1. Penerapan Fungsi Variadic");
        var avgOne = exampleOne(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
        var msgOne = fmt.Sprintf("Rata-rata : %.2f", avgOne)
        fmt.Println(msgOne)
        println()

    // A.20.2. Pengisian Parameter Fungsi Variadic Menggunakan Data Slice
	fmt.Println("# - A.20.2. Pengisian Parameter Fungsi Variadic Menggunakan Data Slice");
        var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
        var avgTwo = exampleOne(numbers...)
        var msgTwo = fmt.Sprintf("Rata-rata : %.2f", avgTwo)
        fmt.Println(msgTwo)
        println()

	// A.20.3. Fungsi Dengan Parameter Biasa & Variadic
	fmt.Println("# - A.20.3. Fungsi Dengan Parameter Biasa & Variadic");
        exampleTwo("wick", "sleeping", "eating")
        fmt.Println("--------------------------")
        var hobbies = []string{"sleeping", "eating"}
        exampleTwo("wick", hobbies...)
        println()
}

func exampleOne(numbers ...int) float64 {
    var total int = 0
    for _, number := range numbers {
        total += number
    }
	fmt.Println(float64(total))
    var avg = float64(total) / float64(len(numbers))
    return avg
}

func exampleTwo(name string, hobbies ...string) {
    var hobbiesAsString = strings.Join(hobbies, ", ")
    fmt.Printf("Hello, my name is: %s\n", name)
    fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
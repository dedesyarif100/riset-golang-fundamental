package main

import (
	"fmt"
    "runtime"
)

func print(till int, message string) {
    for i := 0; i < till; i++ {
        fmt.Println((i + 1), message);
    }
}

func main() {
	// A.30.1. Penerapan Goroutine
	fmt.Println("# - A.30.1. Penerapan Goroutine");
    runtime.GOMAXPROCS(2);

    go print(5, "halo");
    print(5, "apa kabar");

    var input string;
    fmt.Scanln(&input);

	// A.30.1.1. Penggunaan Fungsi runtime.GOMAXPROCS();
	fmt.Println("# - A.30.1.1. Penggunaan Fungsi runtime.GOMAXPROCS()");

	// A.30.1.2. Penggunaan Fungsi fmt.Scanln()
	fmt.Println("# - A.30.1.2. Penggunaan Fungsi fmt.Scanln()");
	// var s1, s2, s3 string
	// fmt.Scanln(&s1, &s2, &s3)

	// // user inputs: "trafalgar d law"

	// fmt.Println(s1) // trafalgar
	// fmt.Println(s2) // d
	// fmt.Println(s3) // law
}
package main

import "fmt"

func main() {
	// A.11.1. Penggunaan Konstanta
	fmt.Println("# - A.11.1. Penggunaan Konstanta");
		const firstName string = "john";
		fmt.Print("halo ", firstName, "!\n");

		const lastName = "wick";
		fmt.Print("nice to meet you ", lastName, "!\n");

		fmt.Println("john wick");
		fmt.Println("john", "wick");

		fmt.Print("john wick\n");
		fmt.Print("john ", "wick\n");
		fmt.Print("john", " ", "wick\n");
}
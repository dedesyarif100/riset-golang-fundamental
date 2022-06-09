package main

import "fmt"

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	go print(5, "TRES")
	go print(10, "APA KABAR")
	print(5, "HALLO")

	var input string
	fmt.Scanln(&input)
}
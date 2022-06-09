package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("# - A.34.1. Penerapan for - range - close Pada Channel");
		runtime.GOMAXPROCS(2)
		message := make(chan string)
		go SendMessage(message)
		PrintMessageExampleOne(message)
		println()

	fmt.Println("# - EXAMPLE TWO");
		number := make(chan int)
		go SendData(number, 10)
		PrintDataExampleTwo(number)
		println()
}


// A.34.1. Penerapan for - range - close Pada Channel
func SendMessage(channel chan<- string) {
	for i := 1; i <= 20; i++ {
		channel <- fmt.Sprintf("DATA %d", i)
	}
	close(channel)
}
func PrintMessageExampleOne(channel <-chan string) {
	fmt.Println("PRINT MESSAGE :")
	for message := range channel {
		fmt.Println(message)
	}
}


// EXAMPLE TWO
func SendData(channel chan<- int, count int) {
	for i := 1; i <= count; i++ {
		channel <- i
	}
	close(channel)
}
func PrintDataExampleTwo(channel <-chan int) {
    fmt.Println("PRINT DATA EXAMPLE TWO :")
	for msg := range channel {
		fmt.Print(msg," | ")
	}
	fmt.Println()
}
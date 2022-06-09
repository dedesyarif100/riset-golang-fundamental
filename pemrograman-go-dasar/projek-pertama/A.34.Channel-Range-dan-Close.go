package main

import (
	"fmt"
    "runtime"
)

func main() {
	// A.34.1. Penerapan for - range - close Pada Channel
	fmt.Println("# - A.34.1. Penerapan for - range - close Pada Channel");
        runtime.GOMAXPROCS(2);
        var messages = make(chan string);
        go sendMessage(messages);
        printMessage(messages);
        fmt.Println();

    // EXAMPLE TWO
    fmt.Println("# - EXAMPLE TWO");
        var number = make(chan int);
        go sendDataExTwo(number, 10);
        receiveDataExTwo(number);
        fmt.Println();
}

// A.34.1. Penerapan for - range - close Pada Channel
func sendMessage(ch chan<- string) {
    for i := 0; i < 20; i++ {
        ch <- fmt.Sprintf("data %d", i);
    }
    close(ch);
}
func printMessage(ch <-chan string) {
    fmt.Println("PRINT MESSAGE :")
    for message := range ch {
        fmt.Println(message);
    }
}

// EXAMPLE TWO
func sendDataExTwo(channel chan<- int, count int) {
    for i := 0; i < count; i++ {
        channel <- i
    }
    close(channel);
}
func receiveDataExTwo(channel <-chan int) {
    fmt.Println("RECEIVE DATA EX TWO :")
    for msg := range channel {
        fmt.Print(msg," | ");
    }
    fmt.Println();
}
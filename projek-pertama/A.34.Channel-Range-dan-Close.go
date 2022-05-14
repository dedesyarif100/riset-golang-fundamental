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
}

func sendMessage(ch chan<- string) {
    for i := 0; i < 20; i++ {
        ch <- fmt.Sprintf("data %d", i);
    }
    close(ch);
}
func printMessage(ch <-chan string) {
    for message := range ch {
        fmt.Println(message);
    }
}
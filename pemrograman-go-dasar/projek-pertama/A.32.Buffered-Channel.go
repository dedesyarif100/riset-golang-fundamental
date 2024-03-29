package main

import (
	"fmt"
    "runtime"
)

func main() {
	// A.32.1. Penerapan Buffered Channel
	fmt.Println("# - A.32.1. Penerapan Buffered Channel");
        runtime.GOMAXPROCS(2);
        messages := make(chan int, 4);
        go func() {
            for {
                i := <-messages;
                fmt.Println("receive data", i);
            }
        }();
        for i := 0; i < 8; i++ {
            fmt.Println("send data", i);
            messages <- i;
        }
}
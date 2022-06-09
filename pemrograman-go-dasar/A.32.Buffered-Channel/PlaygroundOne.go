package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("# - A.32.1. Penerapan Buffered Channel");
		runtime.GOMAXPROCS(1)
		message := make(chan int, 1)
		go func ()  {
			for {
				i := <-message
				fmt.Println("RECEIVE DATA", i)
			}
		}()
		for j := 0; j < 20; j++ {
			fmt.Println("SEND DATA", j)
			message <- j
		}
		println()
}
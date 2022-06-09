package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fmt.Println("# - A.35.1. Penerapan Channel Timeout");
		rand.Seed(time.Now().Unix())
		runtime.GOMAXPROCS(2)
		message := make(chan int)
		go SendData(message)
		RetreiveData(message)
		println()
}

func SendData(channel chan<- int) {
	// fmt.Println("SEND DATA ---")
	// fmt.Println( time.Duration( rand.Int()%10+1 ) * time.Second )
	for i := 0; true; i++ {
		channel <- i
		time.Sleep( time.Duration( rand.Int()%10+1 ) * time.Second )
	}
}
func RetreiveData(channel <-chan int) {
	loop:
	for {
		select {
			case data := <-channel:
				fmt.Println("RECEIVE DATA :", data)
			case <-time.After(time.Second * 5):
				fmt.Println("TIMEOUT")
				break loop
		}
	}
}
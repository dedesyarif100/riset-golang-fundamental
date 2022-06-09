package main

import (
	"fmt"
	"math/rand"
    "runtime"
	"time"
)

func main() {
	// A.35.1. Penerapan Channel Timeout
	fmt.Println("# - A.35.1. Penerapan Channel Timeout");
        rand.Seed(time.Now().Unix());
        runtime.GOMAXPROCS(2);
        // fmt.Println(  );

        var messages = make(chan int);

        go sendData(messages);
        retreiveData(messages);
        println()
}


func sendData(channel chan<- int) {
	fmt.Println( time.Duration( rand.Int()%10+1 ) * time.Second );
	// fmt.Println( time.Sleep( time.Duration( rand.Int()%10+1 ) * time.Second ) );
    for i := 0; true; i++ {
        channel <- i;
        time.Sleep( time.Duration( rand.Int()%10+1 ) * time.Second );
    }
}
func retreiveData(channel <-chan int) {
    loop:
    for {
        select {
			case data := <-channel:
				// fmt.Println(data);
				fmt.Print(`receive data "`, data, `"`, "\n");
			case <-time.After(time.Second * 5):
				fmt.Println("timeout. no activities under 5 seconds");
				break loop;
        }
    }
}
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// A.41.1. Fungsi time.Sleep()
	fmt.Println("# - A.41.1. Fungsi time.Sleep()");

    fmt.Println("start")
    time.Sleep(time.Second * 4)
    fmt.Println("after 4 seconds")

	// A.41.2. Scheduler Menggunakan time.Sleep()
	fmt.Println("# - A.41.2. Scheduler Menggunakan time.Sleep()");
	// for true {
	// 	fmt.Println("Hello !!")
	// 	time.Sleep(1 * time.Second)
	// }

	// A.41.3. Fungsi time.NewTimer()
	fmt.Println("# - A.41.3. Fungsi time.NewTimer()");
	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("start")
	<-timer.C
	fmt.Println("finish")

	// A.41.4. Fungsi time.AfterFunc()
	fmt.Println("# - A.41.4. Fungsi time.AfterFunc()");
	var ch = make(chan bool)
	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})
	fmt.Println("start")
	<-ch
	fmt.Println("finish")

	// A.41.5. Fungsi time.After()
	fmt.Println("# - A.41.5. Fungsi time.After()");
	<-time.After(4 * time.Second)
	fmt.Println("EXPIRED 5")

	// A.41.6. Scheduler Menggunakan Ticker
	fmt.Println("# - A.41.6. Scheduler Menggunakan Ticker");

    // done := make(chan bool)
    // ticker := time.NewTicker(time.Second)

    // go func() {
    //     time.Sleep(10 * time.Second) // wait for 10 seconds
    //     done <- true
    // }()

    // for {
    //     select {
    //     case <-done:
    //         ticker.Stop()
    //         return
    //     case t := <-ticker.C:
    //         fmt.Println("Hello !!", t)
    //     }
    // }

	// A.41.7. Kombinasi Timer & Goroutine
	fmt.Println("# - A.41.7. Kombinasi Timer & Goroutine");
    var timeout = 5;
    var chSeven = make(chan bool);

    go runTimer(timeout, chSeven);
    go runWatcher(timeout, chSeven);

    var input string;
    fmt.Print("what is 725/25 ? ");
    fmt.Scan(&input);

    if input == "29" {
        fmt.Println("the answer is right!");
    } else {
        fmt.Println("the answer is wrong!");
    }
}


// A.41.7. Kombinasi Timer & Goroutine
func runTimer(timeout int, ch chan<- bool) {
    time.AfterFunc(time.Duration(timeout)*time.Second, func() {
        ch <- true
    })
}
func runWatcher(timeout int, ch <-chan bool) {
    <-ch
    fmt.Println("\ntime out! no answer more than", timeout, "seconds")
    os.Exit(0)
}
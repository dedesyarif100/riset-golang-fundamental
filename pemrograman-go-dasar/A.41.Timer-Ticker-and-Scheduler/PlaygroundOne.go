package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("# - A.41.1. Fungsi time.Sleep()");
		fmt.Println("START")
		time.Sleep(time.Second * 3)
		fmt.Println("AFTER 3 SECONDS")
		println()


	fmt.Println("# - A.41.2. Scheduler Menggunakan time.Sleep()");
		println()


	fmt.Println("# - A.41.3. Fungsi time.NewTimer()");
		fmt.Println("-----------------------------------");
		timer1 := time.NewTimer(3 * time.Second)
		<-timer1.C
		fmt.Println("FINISH")
		timer2 := time.NewTimer(3 * time.Second)
		<-timer2.C
		println()


	fmt.Println("# - A.41.4. Fungsi time.AfterFunc()");
		start := make(chan bool)
		time.AfterFunc(3 * time.Second, func() {
			fmt.Println("EXPIRED")
			time.Sleep(time.Second * 2)
			start <- true
		});
		fmt.Println("START")
		<-start

		println()
		ChOne := make(chan bool)
		time.AfterFunc(3 * time.Second, func() {
			fmt.Println("EXPIRED")
			time.Sleep(time.Second * 2)
			ChOne <- true
		});
		fmt.Println("VAL A")
		<-ChOne
		println()


	fmt.Println("# - A.41.5. Fungsi time.After()");
		<-time.After(3 * time.Second)
		fmt.Println("EXPIRED 5")
		println()


	fmt.Println("# - A.41.7. Kombinasi Timer & Goroutine");
		timeout := 3
		ChSeven := make(chan bool)

		go RunTimer(timeout, ChSeven)
		go RunWatcher(timeout, ChSeven)

		var input string;
        fmt.Println("what is 725/25 ? ");
        fmt.Scanln(&input);
}


// A.41.7. Kombinasi Timer & Goroutine
func RunTimer(timeout int, channel chan<- bool) {
    fmt.Println("FUNGSI RUN TIMER");
	time.AfterFunc(time.Duration(timeout) * time.Second, func() {
		channel <- true
	})
}
func RunWatcher(timeout int, channel <-chan bool) {
    fmt.Println("FUNGSI RUN WATCHER");
	<-channel
    fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}
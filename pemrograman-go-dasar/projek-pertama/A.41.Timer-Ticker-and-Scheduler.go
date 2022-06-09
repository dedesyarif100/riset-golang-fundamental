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
        time.Sleep(time.Second * 3)
        fmt.Println("after 3 seconds")
        fmt.Println();

	// A.41.2. Scheduler Menggunakan time.Sleep()
	fmt.Println("# - A.41.2. Scheduler Menggunakan time.Sleep()");
        // for true {
        // 	fmt.Println("Hello !!")
        // 	time.Sleep(1 * time.Second)
        // }
        fmt.Println();

	// A.41.3. Fungsi time.NewTimer()
	fmt.Println("# - A.41.3. Fungsi time.NewTimer()");
        fmt.Println("-----------------------------------");
        fmt.Println("start")
        var timer1 = time.NewTimer(3 * time.Second)
        <-timer1.C
        fmt.Println("finish")
        var timer2 = time.NewTimer(3 * time.Second)
        <-timer2.C
        fmt.Println();

        fmt.Println("-----------------------------------");
        fmt.Println("PERSIAPAN");
        var timerB = time.NewTimer(3 * time.Second);
        <-timerB.C
        fmt.Println("SIAP");
        var timerC = time.NewTimer(3 * time.Second);
        <-timerC.C
        fmt.Println("MULAI");
        var timerD = time.NewTimer(3 * time.Second);
        <-timerD.C
        fmt.Println("FINISH");
        var timerE = time.NewTimer(3 * time.Second);
        <-timerE.C
        fmt.Println();

	// A.41.4. Fungsi time.AfterFunc()
	fmt.Println("# - A.41.4. Fungsi time.AfterFunc()");
        var start = make(chan bool)
        time.AfterFunc(3 * time.Second, func() {
            fmt.Println("expired\n");
            time.Sleep(time.Second * 2);
            start <- true;
        });
        fmt.Println("start");
        <-start;

        var chOne = make(chan bool);
        time.AfterFunc(3 * time.Second, func() {
            fmt.Println("expired\n");
            time.Sleep(time.Second * 2);
            chOne <- true;
        });
        fmt.Println("VAL A");
        <-chOne;

        var chTwo = make(chan bool);
        time.AfterFunc(3 * time.Second, func() {
            fmt.Println("expired\n");
            time.Sleep(time.Second * 2);
            chTwo <- true;
        });
        fmt.Println("VAL B");
        <-chTwo;

        fmt.Println("finish");
        fmt.Println();

	// A.41.5. Fungsi time.After()
	fmt.Println("# - A.41.5. Fungsi time.After()");
        <-time.After(3 * time.Second)
        fmt.Println("EXPIRED 5")
        fmt.Println();

	// A.41.6. Scheduler Menggunakan Ticker
	fmt.Println("# - A.41.6. Scheduler Menggunakan Ticker");
        // JIKA TIDAK DIJALANKAN, MAKA DI COMMENT SAJA
        // done := make(chan bool)
        // ticker := time.NewTicker(time.Second)

        // go func() {
        //     time.Sleep(10 * time.Second) // wait for 10 seconds
        //     done <- true
        // }()

        // for {
        //     select {
        //         case <-done:
        //             ticker.Stop()
        //             return
        //         case t := <-ticker.C:
        //             fmt.Println("Hello !!", t)
        //     }
        // }
        // JIKA TIDAK DIJALANKAN, MAKA DI COMMENT SAJA

        fmt.Println();

	// A.41.7. Kombinasi Timer & Goroutine
	fmt.Println("# - A.41.7. Kombinasi Timer & Goroutine");
        var timeout = 3;
        var chSeven = make(chan bool);
        // var chSevenSecTwo = make(chan bool);

        go runTimer(timeout, chSeven);
        go runWatcher(timeout, chSeven);

        var input string;
        fmt.Println("what is 725/25 ? ");
        fmt.Scanln(&input);
}

// A.41.7. Kombinasi Timer & Goroutine
func runTimer(timeout int, ch chan<- bool) {
    fmt.Println("FUNGSI RUN TIMER");
    time.AfterFunc(time.Duration(timeout) * time.Second, func() {
        ch <- true
    })
}
func runWatcher(timeout int, ch <-chan bool) {
    fmt.Println("FUNGSI RUN WATCHER");
    <-ch
    fmt.Println("\ntime out! no answer more than", timeout, "seconds")
    os.Exit(0)
}
package main

import "sync"
import "runtime"
import "fmt"

func doPrint(wg *sync.WaitGroup, message string) {
	// fmt.Println(wg)
    defer wg.Done()
    fmt.Println(message)
}

func main() {
	// A.59.1. Penerapan sync.WaitGroup
	fmt.Println("# - A.59.1. Penerapan sync.WaitGroup");
    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        var data = fmt.Sprintf("data %d", i)

        wg.Add(1)
        go doPrint(&wg, data)
    }

    wg.Wait()
}
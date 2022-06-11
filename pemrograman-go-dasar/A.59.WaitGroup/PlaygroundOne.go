package main

import (
	"fmt"
	"runtime"
	"sync"
)

func DoPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	fmt.Println("# - A.59.1. Penerapan sync.WaitGroup");
		runtime.GOMAXPROCS(2)

		var input int64
		fmt.Print("INPUT : ")
		fmt.Scanln(&input)
		fmt.Println("------------")

		var wg sync.WaitGroup
		var i int64
		for i = 0; i <= input; i++ {
			data := fmt.Sprintf("DATA %d", i)
			wg.Add(1)
			go DoPrint(&wg, data)
		}

		fmt.Println("WAIT")
		wg.Wait()
}
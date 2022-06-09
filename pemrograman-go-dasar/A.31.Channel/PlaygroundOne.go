package main

import (
	"fmt"
	"runtime"
)

// # - A.31.2. Channel Sebagai Tipe Data Parameter
func printMessage(data chan any) {
	fmt.Println(<-data)
}
// # - A.31.2. Channel Sebagai Tipe Data Parameter

func main() {
	fmt.Println("# - A.31.1. Penerapan Channel");
		runtime.GOMAXPROCS(2)
		messageOne := make(chan string)
		ExampleOne := func(who string) {
			data := fmt.Sprintf("HELLO %s", who)
			messageOne <-data
		}

		go ExampleOne("DEDE")
		go ExampleOne("SYARIF")
		go ExampleOne("HIDAYAT")

		msg1 := <-messageOne
		fmt.Println(msg1)

		msg2 := <-messageOne
		fmt.Println(msg2)

		msg3 := <-messageOne
		fmt.Println(msg3)
		println()

	fmt.Println("# - A.31.2. Channel Sebagai Tipe Data Parameter");
		messageTwo := make(chan any)
		words := []int{1, 2, 3, 4, 5}
		for _, each := range words {
			go func(val any)  {
				data := fmt.Sprintf("HELLO %d", val)
				messageTwo <-data
			}(each)
		}

		for i := 0; i < len(words); i++ {
			printMessage(messageTwo)
		}
		println()

	fmt.Println("# - EXAMPLE THREE");
		var ChanExampleThree = make(chan any)
		var ExampleThree = func(val any) {
			ChanExampleThree <-val
		}

		go ExampleThree(1);
		go ExampleThree(2);
		go ExampleThree(3);

		val1 := <-ChanExampleThree
		fmt.Println(val1)

		val2 := <-ChanExampleThree
		fmt.Println(val2)

		val3 := <-ChanExampleThree
		fmt.Println(val3)
		println()

	fmt.Println("# - EXAMPLE FOUR");
		ChanExampleFour := make(chan int)
		ExampleFour := func(val int)  {
			ChanExampleFour <-val
		}

		for j := 0; j <= 10; j++ {
			go ExampleFour(j)
		}

		var num []int
		for k := 0; k <= 10; k++ {
			num = append(num, <-ChanExampleFour)
		}

		for n := 0; n < len(num); n++ {
			fmt.Printf("DATA : %d\n", num[n])
		}
		println()
}
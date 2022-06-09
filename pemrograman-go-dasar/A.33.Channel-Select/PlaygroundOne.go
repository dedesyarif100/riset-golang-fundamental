package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("# - A.32.1. Penerapan Buffered Channel");
		runtime.GOMAXPROCS(2)

		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Printf("NUMBER : %d\n", numbers)

		ch1 := make(chan float64)
		go GetAverage(numbers, ch1)

		ch2 := make(chan int)
		go GetMax(numbers, ch2)

		ch3 := make(chan int)
		go GetNumber(numbers, ch3)

		data := []any{ch1, ch2, ch3}

		fmt.Println("LOOP -----------------------------")
		for i := 1; i <= len(data); i++ {
			fmt.Printf("START : %d\n", i)
			select {
				case average := <-ch1:
					fmt.Printf("AVERAGE : %.2f\n", average)
				case max := <-ch2:
					fmt.Printf("AVERAGE : %d\n", max)
				case number := <-ch3:
					fmt.Printf("AVERAGE : %d\n", number)
			}
			fmt.Printf("NEXT  : %d\n", i)
			println()
		}
		println()
}

func GetAverage(numbers []int, channel chan float64) {
	fmt.Println("GET AVERAGE :")
	var sum = 0;
	for _, val := range numbers {
		sum += val
	}
	result := float64(sum) / float64(len(numbers))
	channel <- result
}

func GetMax(numbers []int, channel chan int) {
	fmt.Println("GET MAX :")
	var max = numbers[0]
	for _, val := range numbers {
		if max < val {
			max = val
		}
	}
	channel <- max
}

func GetNumber(numbers []int, channel chan int) {
	fmt.Println("GET NUMBER :")
	val := numbers[0]
	channel <- val
}
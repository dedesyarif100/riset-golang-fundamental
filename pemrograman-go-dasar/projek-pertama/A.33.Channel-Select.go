package main

import (
	"fmt"
    "runtime"
)

func main() {
	// A.33.1. Penerapan Keyword select
	fmt.Println("# - A.33.1. Penerapan Keyword select");
        runtime.GOMAXPROCS(2);

        var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3};
        fmt.Println("numbers :", numbers);

        var ch1 = make(chan float64);
        go getAverage(numbers, ch1);

        var ch2 = make(chan int);
        go getMax(numbers, ch2);

        var ch3 = make(chan int);
        go getNumber(numbers, ch3);

        for i := 0; i < 3; i++ {
            select {
                case avg := <-ch1:
                    fmt.Printf("Avg \t: %.2f \n", avg);
                case max := <-ch2:
                    fmt.Printf("Max \t: %d \n", max);
                case val := <-ch3:
                    fmt.Println("Val \t:", val);
            }
        }
        println()
}

func getAverage(numbers []int, ch chan float64) {
    var sum = 0;
    for _, e := range numbers {
        sum += e;
    }
    ch <- float64(sum) / float64(len(numbers));
}

func getMax(numbers []int, ch chan int) {
    var max = numbers[0];
    for _, e := range numbers {
        if max < e {
            max = e;
        }
    }
    ch <- max;
}

func getNumber(numbers []int, ch chan int) {
	val := numbers[0];
	ch <- val;
}
package main

import (
	"fmt"
	// "reflect"
	"runtime"
)

func main() {
	// A.33.1. Penerapan Keyword select
	fmt.Println("# - A.33.1. Penerapan Keyword select");
        runtime.GOMAXPROCS(2);

        var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3};
        fmt.Println("numbers :", numbers);

        var ch1 = make(chan float64);
        go GetAverage(numbers, ch1);
        
        var ch2 = make(chan int);
        go GetMax(numbers, ch2);

        var ch3 = make(chan int);
        go GetNumber(numbers, ch3);
        
        data := []any{ch1, ch2, ch3};

        fmt.Println("LOOP ---------------------------------");
        for i := 1; i <= len(data); i++ {
            fmt.Println("START :", i)
            select {
                case average := <-ch1:
                    fmt.Printf("AVERAGE \t: %.2f \n", average);
                case max := <-ch2:
                    fmt.Printf("MAX \t\t: %d \n", max);
                case number := <-ch3:
                    fmt.Println("NUMBER \t\t:", number);
            }
            fmt.Println("NEXT  :", i)
            println()
        }
        println()
}

func GetAverage(numbers []int, ch chan float64) {
    fmt.Println("GET AVERAGE \t:")
    var sum = 0;
    for _, e := range numbers {
        sum += e;
    }
    ch <- float64(sum) / float64(len(numbers));
}

func GetMax(numbers []int, ch chan int) {
    fmt.Println("GET MAX \t:")
    var max = numbers[0];
    for _, e := range numbers {
        if max < e {
            max = e;
        }
    }
    ch <- max;
}

func GetNumber(numbers []int, ch chan int) {
    fmt.Println("GET NUMBER \t:")
    val := numbers[0];
	ch <- val;
}
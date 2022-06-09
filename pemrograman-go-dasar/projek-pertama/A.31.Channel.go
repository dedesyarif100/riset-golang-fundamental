package main

import (
	"fmt"
    "runtime"
)

func main() {
	// A.31.1. Penerapan Channel
	fmt.Println("# - A.31.1. Penerapan Channel");
        runtime.GOMAXPROCS(3);
        var messagesOne = make(chan string);
        // fmt.Println(messagesOne);
        var sayHelloTo = func(who string) {
            var data = fmt.Sprintf("hello %s", who);
            messagesOne <-data;
            // fmt.Println("CEK DATA :",data);
        }

        // EKSEKUSI BERSIFAT ASYNCRONOUS
        go sayHelloTo("john wick");
        go sayHelloTo("ethan hunt");
        go sayHelloTo("jason bourne");

        var message1 = <-messagesOne;
        fmt.Println(message1);

        var message2 = <-messagesOne;
        fmt.Println(message2);

        var message3 = <-messagesOne;
        fmt.Println(message3);
        fmt.Println();
        // EKSEKUSI BERSIFAT ASYNCRONOUS

	// A.31.2. Channel Sebagai Tipe Data Parameter
	fmt.Println("# - A.31.2. Channel Sebagai Tipe Data Parameter");
        var messagesTwo = make(chan any);

        words := []string{"wick", "hunt", "bourne", "dede", "tes"}
        for _, each := range words {
            go func(who string) {
                var data = fmt.Sprintf("hello %s", who);
                messagesTwo <-data;
            }(each)
        }
        // fmt.Println(messagesTwo);

        // fmt.Println()
        // EKSEKUSI BERSIFAT SYNCRONOUS
        for i := 0; i < len(words); i++ {
            printMessage(messagesTwo);
        }
        // EKSEKUSI BERSIFAT SYNCRONOUS
        fmt.Println();

    // EXAMPLE THREE
    fmt.Println("# - EXAMPLE THREE");
        var chanExampleThree = make(chan int);
        var exampleThree = func(val int) {
            // var number = fmt.Sprintln(val);
            // var number = fmt.Printf("hello %o", val);
            chanExampleThree <-val;
        }

        go exampleThree(1);
        go exampleThree(2);
        go exampleThree(3);

        val1 := <-chanExampleThree;
        fmt.Println(val1);

        val2 := <-chanExampleThree;
        fmt.Println(val2);

        val3 := <-chanExampleThree;
        fmt.Println(val3);
        fmt.Println();

    fmt.Println("# - EXAMPLE FOUR");
		ChanExampleFour := make(chan any)
		ExampleFour := func(val any) {
			ChanExampleFour <-val
		}

		for j := 0; j <= 10; j++ {
			go ExampleFour(j)
		}

		var num []any
		for k := 0; k <= 10; k++ {
			num = append(num, <-ChanExampleFour)
		}

        for n := 0; n < len(num); n++ {
			fmt.Printf("DATA : %d\n", num[n])
		}
		println()
}

func printMessage(what chan any) {
    fmt.Println(<-what)
}
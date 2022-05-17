package main

import (
	"fmt"
    "runtime"
)

func main() {
	// A.31.1. Penerapan Channel
	fmt.Println("# - A.31.1. Penerapan Channel");
        runtime.GOMAXPROCS(2);
        var messagesOne = make(chan string);
        // fmt.Println(messagesOne);
        var sayHelloTo = func(who string) {
            var data = fmt.Sprintf("hello %s", who);
            messagesOne <- data;
            fmt.Println("CEK DATA :",data);
        }

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

	// A.31.2. Channel Sebagai Tipe Data Parameter
	fmt.Println("# - A.31.2. Channel Sebagai Tipe Data Parameter");
        var messagesTwo = make(chan string);

        for _, each := range []string{"wick", "hunt", "bourne"} {
            go func(who string) {
                var data = fmt.Sprintf("hello %s", who);
                messagesTwo <- data;
            }(each)
        }

        for i := 0; i < 3; i++ {
            printMessage(messagesTwo);
        }
}


func printMessage(what chan string) {
    fmt.Println(<-what)
}
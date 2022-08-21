package main

import (
	"fmt"
	"runtime"
)

func main() {
	case4()
}

func case1() {
	var jumlahData int
	fmt.Print("JUMLAH DATA : ")
	fmt.Scanln(&jumlahData)	
	
	channelOne := make(chan string)
	var sendData = func(param string) {
		var data = fmt.Sprintf("DATA %s", param)
		channelOne <-data
	}

	var input string
	for i := 1; i <= jumlahData; i++ {
		fmt.Print("INPUT : ")
		fmt.Scanln(&input)
		go sendData(input)
	}

	for j := 1; j <= jumlahData; j++ {
		var msg = <-channelOne
		fmt.Println(msg)
	}
}

func case2() {
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
}

func case3() {
	runtime.GOMAXPROCS(2);
	messages := make(chan any, 3);
	go func() {
		for {
			i := <-messages;
			fmt.Println("receive data", i);
		}
	}();
	for i := 0; i < 5; i++ {
		fmt.Println("send data", i);
		messages<- i;
	}
	fmt.Println();
}

func case4() {
	runtime.GOMAXPROCS(2);
	messages := make(chan any, 3);
	go test(messages)
	for i := 0; i < 5; i++ {
		fmt.Println("send data", i);
		messages<- i;
	}
	fmt.Println();
}

func test(data chan any) {
	for {
		fmt.Println("receive data", <-data);
	}
}

func printMessage(what chan any) {
    fmt.Println(<-what)
}

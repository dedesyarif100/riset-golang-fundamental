package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
    var message = "Welcome"
    w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
    var message = "Hello world!"
    w.Write([]byte(message))
}

func main() {
	// B.1.1. Pembuatan Aplikasi
	// fmt.Println("# - B.1.1. Pembuatan Aplikasi");
	// B.1.2. Web Server Menggunakan http.Server
	fmt.Println("# - B.1.2. Web Server Menggunakan http.Server");

    http.HandleFunc("/", handlerIndex)
    http.HandleFunc("/index", handlerIndex)
    http.HandleFunc("/hello", handlerHello)

    // var address = "localhost:9000"
    // fmt.Printf("server started at %s\n", address)
    // err := http.ListenAndServe(address, nil)
    // if err != nil {
    //     fmt.Println(err.Error())
    // }
	
	var address = ":9000"
	fmt.Printf("server started at %s\n", address)

	server := new(http.Server)
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10
	server.Addr = address
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
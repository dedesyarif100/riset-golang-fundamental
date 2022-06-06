package main

import (
    "fmt"
	"os"
	"flag"
)

func main() {
	// A.48.1. Penggunaan Arguments
	fmt.Println("# - A.48.1. Penggunaan Arguments");
		var argsRaw = os.Args
		fmt.Printf("-> %#v\n", argsRaw)
		// -> []string{".../bab45", "banana", "potato", "ice cream"}

		var args = argsRaw[1:]
		fmt.Printf("-> %#v\n", args)
		// -> []string{"banana", "potatao", "ice cream"}
		fmt.Println();

	// A.48.2. Penggunaan Flag
	fmt.Println("# - A.48.2. Penggunaan Flag");
		var name = flag.String("name", "anonymous", "type your name")
		var age = flag.Int64("age", 25, "type your age")

		flag.Parse()
		fmt.Printf("name\t: %s\n", *name)
		fmt.Printf("age\t: %d\n", *age)

		// var dataName = flag.String("name", "anonymous", "type your name")
		// fmt.Println(*dataName)
		fmt.Println();

	// A.48.3. Deklarasi Flag Dengan Cara Passing Reference Variabel Penampung Data
	fmt.Println("# - A.48.3. Deklarasi Flag Dengan Cara Passing Reference Variabel Penampung Data");
		// cara ke-1
		// var data1 = flag.String("name", "anonymous", "type your name")
		// fmt.Println(*data1)

		// cara ke-2
		var data2 string
		flag.StringVar(&data2, "gender", "male", "type your gender")
		fmt.Println(data2)
}
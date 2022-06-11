package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("# - A.48.1. Penggunaan Arguments");
		ArgsRaw := os.Args
		fmt.Printf("-> %#v\n", ArgsRaw)

		Args := ArgsRaw[1:]
		fmt.Printf("-> %#v\n", Args)
		println()

	fmt.Println("# - A.48.2. Penggunaan Flag");
		name := flag.String("NAME", "DEDE", "TYPE YOUR NAME")
		age := flag.Int64("AGE", 25, "TYPE YOUR AGE")

		fmt.Printf("NAME\t : %s\n", *name)
		fmt.Printf("AGE\t : %d\n", *age)
		println()

	fmt.Println("# - A.48.3. Deklarasi Flag Dengan Cara Passing Reference Variabel Penampung Data");
		var ExampleThree string
		flag.StringVar(&ExampleThree, "NAMA", "DEDE", "tes")
		fmt.Println(ExampleThree)

		var Ex3secOne int64
		flag.Int64Var(&Ex3secOne, "NUMBER", 22, "TES")
		fmt.Println(Ex3secOne)
		println()
}
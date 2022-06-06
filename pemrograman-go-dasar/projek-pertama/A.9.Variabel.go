package main

import "fmt"

type typeOne interface{}
type typeTwo any

func main() {
	// A.9.2 Deklarasi Variabel Menggunakan Keyword var
	fmt.Println("# - A.9.2 Deklarasi Variabel Menggunakan Keyword var");
		println()

	// A.9.3 Deklarasi Variabel Tanpa Tipe Data
	fmt.Println("# - A.9.3 Deklarasi Variabel Tanpa Tipe Data");
		var firstName string = "DEDE";
		// var lastName string;
		var lastName = "SYARIF";

		var number1 int = 12;
		var number2 int = 2
		var tesOne []interface{} = []interface{}{1, 2, 3, 4, 5}
		fmt.Println(tesOne)

		var parseOne int = 3
		var tesTwo [][3]typeTwo = [][3]typeTwo{{parseOne, "A", nil}, { nil, 1}}
		fmt.Println(tesTwo)

		number3 := 2;
		var result int;

		result = number1 * number2 + number3;
		
		university := "Untag";

		fmt.Println("NAMA 		: "+firstName+ " " +lastName);
		fmt.Println("RESULT 		: ",result);
		fmt.Println("UNIVERSITY 	: ",university);

		var1 := 12;
		// var1 = "DEDE";
		fmt.Println("VAR 		: ",var1);
		println()

	// A.9.4 Deklarasi Multi Variabel
	fmt.Println("# - A.9.4 Deklarasi Multi Variabel");
		var first, second, third string;
		first, second, third = "satu", "dua", "tiga";
		fmt.Println(first+" | "+second+" | "+third);

		var fourth, fifth, sixth string = "empat", "lima", "enam";
		fmt.Println(fourth+" | "+fifth+" | "+sixth);

		seventh, eight, ninth := "tujuh", "delapan", "sembilan";
		fmt.Println(seventh+" | "+eight+" | "+ninth);

		one, isFriday, twoPointTwo, say, arrayInt := 1, true, 2.2, "hello", []int{1, 2, 3};
		fmt.Println(one," | ",isFriday," | ",twoPointTwo," | ",say," | ",arrayInt);

		arrayInteger, 
		arrayString, 
		arrayBool, 
		arrayFloat, 
		arrayInterface := 
		[]int{1, 2, 3}, 
		[]string{"A", "B"}, 
		[]bool{true, false}, 
		[]float64{1.2, 1.5}, 
		[]interface{}{1, 1.2, "A", true, []int{1, 2, 3}}
		
		fmt.Println(
			arrayInteger," | ",
			arrayString," | ",
			arrayBool," | ",
			arrayFloat," | ",
			arrayInterface,
		);

		a, b, c := one, isFriday, say;
		fmt.Println(a, b, c);
		println()

	// A.9.5 Variabel Underscore _
	fmt.Println("# - A.9.5 Variabel Underscore _");
		_ = "belajar golang";
		_ = "Golang itu mudah";
		name, _ := "JODH", "WICK";
		fmt.Println(name);

	// A.9.6 Deklarasi Variabel Menggunakan Keyword

	// A.9.7 Deklarasi Variabel Menggunakan Keyword 
}
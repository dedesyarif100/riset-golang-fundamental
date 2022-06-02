package main

import "fmt"

func main() {
	// A.9.2 Deklarasi Variabel Menggunakan Keyword var
	// A.9.3 Deklarasi Variabel Tanpa Tipe Data
	fmt.Println("# - A.9.2 Deklarasi Variabel Menggunakan Keyword var");
	fmt.Println("# - A.9.3 Deklarasi Variabel Tanpa Tipe Data");
		var firstName string = "DEDE";
		// var lastName string;
		var lastName = "SYARIF";

		var number1 int = 12;
		var number2 int = 2

		number3 := 2;
		var result int;

		result = number1 * number2 + number3;
		
		university := "Untag";

		fmt.Println("NAMA : "+firstName+ " " +lastName);
		fmt.Println("RESULT : ",result);
		fmt.Println("UNIVERSITY : ",university);

		var1 := 12;
		// var1 = "DEDE";
		fmt.Println("VAR : ",var1);

	// A.9.4 Deklarasi Multi Variabel
	fmt.Println("# - A.9.4 Deklarasi Multi Variabel");
		var first, second, third string;
		first, second, third = "satu", "dua", "tiga";
		fmt.Println(first+" | "+second+" | "+third);

		var fourth, fifth, sixth string = "empat", "lima", "enam";
		fmt.Println(fourth+" | "+fifth+" | "+sixth);

		seventh, eight, ninth := "tujuh", "delapan", "sembilan";
		fmt.Println(seventh+" | "+eight+" | "+ninth);

		one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello";
		fmt.Println(one," | ",isFriday," | ",twoPointTwo," | ",say);

		a, b, c := one, isFriday, say;
		fmt.Println(a, b, c);

	// A.9.5 Variabel Underscore _
	fmt.Println("# - A.9.5 Variabel Underscore _");
		_ = "belajar golang";
		_ = "Golang itu mudah";
		name, _ := "JODH", "WICK";
		fmt.Println(name);

	// A.9.6 Deklarasi Variabel Menggunakan Keyword

	// A.9.7 Deklarasi Variabel Menggunakan Keyword 
}
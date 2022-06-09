package main

import (
	. "A.26.Properti-Public-dan-Private/library"
	f "fmt"
)

func main() {
	// A.26.3. Penggunaan Package, Import, Dan Hak Akses Exported dan Unexported
	f.Println("# - A.26.3. Penggunaan Package, Import, Dan Hak Akses Exported dan Unexported");
		ApelMalang("SAYA MEMBELI APEL DI MALANG");
		// apelHijau("ethan")
		f.Println();

	// A.26.4. Penggunaan Hak Akses Exported dan Unexported pada Struct dan Propertinya
	// A.26.5. Import Dengan Prefix Tanda Titik
	// A.26.6. Pemanfaatan Alias Ketika Import
	f.Println("# - A.26.4. Penggunaan Hak Akses Exported dan Unexported pada Struct dan Propertinya");
		f.Println();
	
	f.Println("# - A.26.5. Import Dengan Prefix Tanda Titik");
		f.Println();
	
	f.Println("# - A.26.6. Pemanfaatan Alias Ketika Import");
		var s1 = Student{"ethan", 21};
		f.Println("NAME 	:", s1.Name);
		f.Println("GRADE	:", s1.Grade);
		f.Println();

	// A.26.7. Mengakses Properti Dalam File Yang Package-nya Sama
	f.Println("# - A.26.7. Mengakses Properti Dalam File Yang Package-nya Sama");
		exampleSeven("FEBBY SAILOLIN");
		f.Println();

	// A.26.7.1. Fungsi init()
	f.Println("# - A.26.7.1. Fungsi init()");
		f.Printf("NAME  	: %s\n", ExampleSix.Name);
		f.Printf("GRADE 	: %d\n", ExampleSix.Grade);
		f.Println("VALUE A :", ValA);
		f.Println("VALUE B :", valB);
		f.Println();
}
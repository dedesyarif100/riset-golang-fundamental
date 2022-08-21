package main

import (
	"A.26.Properti-Public-dan-Private/library"
	f "fmt"
)

func main() {
	// A.26.3. Penggunaan Package, Import, Dan Hak Akses Exported dan Unexported
	f.Println("# - A.26.3. Penggunaan Package, Import, Dan Hak Akses Exported dan Unexported");
		library.ApelMalang("SAYA MEMBELI APEL DI MALANG");
		// library.apelHijau("ethan")
		f.Println();

	// A.26.4. Penggunaan Hak Akses Exported dan Unexported pada Struct dan Propertinya
	// A.26.5. Import Dengan Prefix Tanda Titik
	// A.26.6. Pemanfaatan Alias Ketika Import
	f.Println("# - A.26.4. Penggunaan Hak Akses Exported dan Unexported pada Struct dan Propertinya");
		f.Println();
	
	f.Println("# - A.26.5. Import Dengan Prefix Tanda Titik");
		f.Println();
	
	f.Println("# - A.26.6. Pemanfaatan Alias Ketika Import");
		var s1 = library.Student{"ethan", 21, "SBY", []int{12, 3}};
		f.Println("NAME 	:", s1.Name);
		f.Println("GRADE	:", s1.Grade);
		f.Println("ADDRESS	:", s1.Address);
		f.Println("DATA	:", s1.Data);
		f.Println();

	// A.26.7. Mengakses Properti Dalam File Yang Package-nya Sama
	f.Println("# - A.26.7. Mengakses Properti Dalam File Yang Package-nya Sama");
		exampleSeven("FEBBY SAILOLIN");
		f.Println();

	// A.26.7.1. Fungsi init()
	f.Println("# - A.26.7.1. Fungsi init()");
		f.Printf("NAME  	: %s\n", library.ExampleSix.Name);
		f.Printf("GRADE 	: %d\n", library.ExampleSix.Grade);
		f.Println("VALUE A :", library.ValA);
		f.Println("VALUE B :", valB);
		f.Println();
}
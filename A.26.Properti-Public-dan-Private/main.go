package main

import (
	// "A.26.Properti-Public-dan-Private/library"
	. "A.26.Properti-Public-dan-Private/library"
	f "fmt"
)

func main() {
	// A.26.3. Penggunaan Package, Import, Dan Hak Akses Exported dan Unexported
	f.Println("# - A.26.3. Penggunaan Package, Import, Dan Hak Akses Exported dan Unexported");
		ApelMalang("SAYA MEMBELI APEL DI MALANG");
		// introduce("ethan")

	// A.26.4. Penggunaan Hak Akses Exported dan Unexported pada Struct dan Propertinya
	// A.26.5. Import Dengan Prefix Tanda Titik
	// A.26.6. Pemanfaatan Alias Ketika Import
	f.Println("# - A.26.4. Penggunaan Hak Akses Exported dan Unexported pada Struct dan Propertinya");
	f.Println("# - A.26.5. Import Dengan Prefix Tanda Titik");
	f.Println("# - A.26.6. Pemanfaatan Alias Ketika Import");
		var s1 = Student{"ethan", 21};
		f.Println("name ", s1.Name);
		f.Println("grade", s1.Grade);

	// A.26.7. Mengakses Properti Dalam File Yang Package-nya Sama
	f.Println("# - A.26.7. Mengakses Properti Dalam File Yang Package-nya Sama");
    	// ExampleSeven("ethan");

	// A.26.7.1. Fungsi init()
	f.Println("# - A.26.7.1. Fungsi init()");
		f.Printf("Name  : %s\n", ExampleSix.Name);
		f.Printf("Grade : %d\n", ExampleSix.Grade);
}
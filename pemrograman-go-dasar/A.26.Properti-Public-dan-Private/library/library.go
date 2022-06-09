package library

import "fmt"


// A.26.3. Penggunaan Package, Import, Dan Hak Akses Exported dan Unexported
func ApelMalang(name string) {
    fmt.Println("hello")
	apelHijau(name)
}
func apelHijau(name string) {
    fmt.Println("nama saya", name)
}


// A.26.4. Penggunaan Hak Akses Exported dan Unexported pada Struct dan Propertinya
type Student struct {
    Name  	string
    Grade 	int
}


// A.26.7.1. Fungsi init()
var ExampleSix = struct {
	Name 	string
	Grade 	int
}{}

var ValA = "VALUE A";

func init() {
	ExampleSix.Name = "DEDE SYARIFUDIN"
	ExampleSix.Grade = 2
    fmt.Println("--> library/library.go imported")
	fmt.Println("NAME 	:",ExampleSix.Name);
	fmt.Println("GRADE 	:",ExampleSix.Grade);
	fmt.Println();
}
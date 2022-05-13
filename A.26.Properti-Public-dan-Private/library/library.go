package library

import "fmt"

// A.26.3. Penggunaan Package, Import, Dan Hak Akses Exported dan Unexported
func SayHello(name string) {
    fmt.Println("hello")
	introduce(name)
}
func introduce(name string) {
    fmt.Println("nama saya", name)
}

// A.26.4. Penggunaan Hak Akses Exported dan Unexported pada Struct dan Propertinya
type Student struct {
    Name  string
    Grade int
}

// A.26.7.1. Fungsi init()
var ExampleSix = struct {
	Name string
	Grade int
}{}

func init() {
	ExampleSix.Name = "DEDE SYARIFUDIN"
	ExampleSix.Grade = 2

    fmt.Println("--> library/library.go imported")
}
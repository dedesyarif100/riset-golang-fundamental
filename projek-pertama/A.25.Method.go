package main

import (
	"fmt"
	"strings"
)

type exampleOne struct {
	name  string
	grade int
}

type exampleTwo struct {
	name  string
	grade int
}

func main() {
	// A.25.1. Penerapan Method
	fmt.Println("# - A.25.1. Penerapan Method");
    var exampleOne = exampleOne{"john wick", 21}
    exampleOne.sayHello()
    var name = exampleOne.getNameAt(2)
    fmt.Println("nama panggilan :", name)
    fmt.Println("------------------------------")
	// var result = &exampleOne{"tes 123", 21}
	// result.sayHello()

	// A.25.2. Method Pointer
	fmt.Println("# - A.25.2. Method Pointer");
	var s1 = exampleTwo{"DEDE", 22}
	fmt.Println("before", s1.name)
    s1.changeName1("jason bourne")
    fmt.Println("s1 after changeName1", s1.name) // john wick
    s1.changeName2("ethan hunt")
    fmt.Println("s1 after changeName2", s1.name) // ethan hunt

	// pengaksesan method dari variabel objek biasa
	// var s1 = student{"john wick", 21}
	// s1.sayHello()

	// pengaksesan method dari variabel objek pointer
	// var s2 = &student{"ethan hunt", 22}
	// s2.sayHello()
}

// A.25.1. Penerapan Method
func (ex exampleOne) sayHello() {
    fmt.Println("halo", ex.name)
}
func (ex exampleOne) getNameAt(i int) string {
	return strings.Split(ex.name, " ")[i-1]
}

// A.25.2. Method Pointer
func (s exampleTwo) changeName1(name string) {
    fmt.Println("on changeName1, name changed to", name)
    s.name = name
}
func (s *exampleTwo) changeName2(name string) {
    fmt.Println("on changeName2, name changed to", name)
    s.name = name
}
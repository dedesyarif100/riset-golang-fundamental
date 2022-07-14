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
		var exOne = exampleOne{"DEDE SYARIFUDIN HIDAYAT", 21}
		exOne.sayHello()
		var name = exOne.getNameAt(3)
		fmt.Println("nama panggilan 	:", name)
		fmt.Println("--------------------------")
		var resultOne = &exampleOne{"TES 123", 21}
		resultOne.sayHello()
		fmt.Println("--------------------------")
		var resultTwo = &exampleOne{"ONE", 22}
		resultTwo.sayHello()
        fmt.Println();

	// A.25.2. Method Pointer
	fmt.Println("# - A.25.2. Method Pointer");
		var s1 = exampleTwo{"DEDE", 22}
		fmt.Println("before", s1.name)
		s1.changeName1("jason bourne")
		fmt.Println("S1 AFTER CHANGE NAME 1 :", s1.name) // john wick
		s1.changeName2("ethan hunt")
		fmt.Println("S1 AFTER CHANGE NAME 2 :", s1.name) // ethan hunt
		fmt.Println("--------------------------")

		// pengaksesan method dari variabel objek biasa
		var s2 = exampleTwo{"DEDE SYARIFUDIN HIDAYAT", 21}
		s2.running("FEBBY SAILOLIN", 24)
		fmt.Println("NAME  	:",s2.name)
		fmt.Println("GRADE  	:",s2.grade)
		fmt.Println("--------------------------")

		// pengaksesan method dari variabel objek pointer
		var s3 = &exampleTwo{"ethan hunt", 22}
		fmt.Println("NAME  	:",s3.name)
		fmt.Println("GRADE  	:",s3.grade)
		fmt.Println("--------------------------")
		s3.running("ELIZABETH SAILOLIN", 25)
		fmt.Println("NAME  	:",s3.name)
		fmt.Println("GRADE  	:",s3.grade)
        fmt.Println();
}

// A.25.1. Penerapan Method
func (ex exampleOne) sayHello() {
    fmt.Println("halo 		:", ex.name);
	fmt.Println("my grade is 	:", ex.grade);
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
func (val *exampleTwo) running(name string, grade int) {
	val.name = name
	val.grade = grade
}
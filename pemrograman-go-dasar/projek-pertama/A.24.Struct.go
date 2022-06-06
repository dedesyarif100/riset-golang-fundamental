package main

import (
	"fmt"
)

// A.24.1. Deklarasi Struct
type exampleOne struct {
	name string
	grade int
}

// A.24.3. Inisialisasi Object Struct
type exampleThree struct {
	name string
	grade int
}

// A.24.4. Variabel Objek Pointer
type exampleFour struct {
	name string
	grade int
}

// A.24.5. Embedded Struct
type exampleFive struct {
	name string
	age  int
}
type derivativeExampleFive struct {
	grade int
	exampleFive
}

// A.24.6. Embedded Struct Dengan Nama Property Yang Sama
type exampleSix struct {
	name string
	age  int
}
type derivativeExampleSix struct {
	exampleSix
	age   int
	grade int
}

// A.24.7. Pengisian Nilai Sub-Struct
type exampleSeven struct {
	name string
	age  int
}
type derivativeExampleSeven struct {
	exampleSeven
	grade int
}

// A.24.8. Anonymous Struct
type exampleEight struct {
	name string
	age  int
}

// A.24.9. Kombinasi Slice & Struct
type exampleNine struct {
	name string
	age  int
}

// A.24.10. Inisialisasi Slice Anonymous Struct
type exampleTen struct {
	name string
	age  int
}

// A.24.11. Deklarasi Anonymous Struct Menggunakan Keyword var
type exampleEleven struct {
	name string
	age  int
}

// A.24.12. Nested struct
type exampleTwelve struct {
	derivativeExampleTwelve struct {
		name string
		age int
	}
	grade int
	hobbies []string
}

// A.24.13. Deklarasi Dan Inisialisasi Struct Secara Horizontal
type exampleThirTeen struct { name string; age int; hobbies []string }

// A.24.14. Tag property dalam struct
type exampleFourTeen struct {
    name string `nama`
    age  int    `umur`
}

// A.24.15. Type Alias
type exampleFiveTeen struct {
	name string
	age  int
}
type derivativeExampleFiveTeen = exampleFiveTeen


func main() {
	// A.24.1. Deklarasi Struct
	fmt.Println("# - A.24.1. Deklarasi Struct");
		fmt.Println();

	// A.24.2. Penerapan Struct
	fmt.Println("# - A.24.2. Penerapan Struct");
		var resultExampleOne exampleOne
		resultExampleOne.name = "john wick"
		resultExampleOne.grade = 2
		fmt.Println("name  :", resultExampleOne.name)
		fmt.Println("grade :", resultExampleOne.grade)
        fmt.Println();

	// A.24.3. Inisialisasi Object Struct
	fmt.Println("# - A.24.3. Inisialisasi Object Struct");
		var resultExampleThree_SectionOne = exampleThree{}
		resultExampleThree_SectionOne.name = "wick"
		resultExampleThree_SectionOne.grade = 2
		var resultExampleThree_SectionTwo = exampleThree{"ethan", 11}
		var resultExampleThree_SectionThree = exampleThree{name: "jason"}
		fmt.Println("resultExampleThree_SectionOne name :", resultExampleThree_SectionOne.name)
		fmt.Println("resultExampleThree_SectionOne grade :", resultExampleThree_SectionOne.grade)
		fmt.Println("resultExampleThree_SectionTwo name :", resultExampleThree_SectionTwo.name)
		fmt.Println("resultExampleThree_SectionThree name :", resultExampleThree_SectionThree.name)

		fmt.Println("-------------------------------------------------");
		var resultExampleThree_SectionFour = exampleThree{name: "wayne", grade: 2}
		var resultExampleThree_SectionFive = exampleThree{grade: 11, name: "bruce"}
		fmt.Println("resultExampleThree_SectionFour name :", resultExampleThree_SectionFour.name)
		fmt.Println("resultExampleThree_SectionFour grade :", resultExampleThree_SectionFour.grade)

		fmt.Println("resultExampleThree_SectionFive name :", resultExampleThree_SectionFive.name)
		fmt.Println("resultExampleThree_SectionFive grade :", resultExampleThree_SectionFive.grade)
        fmt.Println();

	// A.24.4. Variabel Objek Pointer
	fmt.Println("# - A.24.4. Variabel Objek Pointer");
		var resultExampleFour_SectionOne = exampleFour{name: "DEDE", grade: 12}
		var resultExampleFour_SectionTwo *exampleFour = &resultExampleFour_SectionOne

		fmt.Println("resultExampleFour_SectionOne, name :", resultExampleFour_SectionOne.name)
		fmt.Println("resultExampleFour_SectionTwo, name :", resultExampleFour_SectionTwo.name)

		resultExampleFour_SectionTwo.name = "ethan"
		fmt.Println("resultExampleFour_SectionOne, name :", resultExampleFour_SectionOne.name)
		fmt.Println("resultExampleFour_SectionTwo, name :", resultExampleFour_SectionTwo.name)
        fmt.Println();

	// A.24.5. Embedded Struct
	fmt.Println("# - A.24.5. Embedded Struct");
		var resultExampleFive_SectionOne = derivativeExampleFive{}
		resultExampleFive_SectionOne.name = "wick"
		resultExampleFive_SectionOne.age = 21
		resultExampleFive_SectionOne.grade = 2

		fmt.Println("name  :", resultExampleFive_SectionOne.name)
		fmt.Println("age   :", resultExampleFive_SectionOne.age)
		fmt.Println("name  :", resultExampleFive_SectionOne.exampleFive.name)
		fmt.Println("age   :", resultExampleFive_SectionOne.exampleFive.age)
		fmt.Println("grade :", resultExampleFive_SectionOne.grade)
        fmt.Println();

	// A.24.6. Embedded Struct Dengan Nama Property Yang Sama
	fmt.Println("# - A.24.6. Embedded Struct Dengan Nama Property Yang Sama");
		var resultExampleSix_SectionOne = derivativeExampleSix{}
		resultExampleSix_SectionOne.name = "DEDE"
		resultExampleSix_SectionOne.age = 21
		resultExampleSix_SectionOne.exampleSix.age = 22

		fmt.Println(resultExampleSix_SectionOne.name)
		fmt.Println(resultExampleSix_SectionOne.age)
		fmt.Println(resultExampleSix_SectionOne.exampleSix.age)
        fmt.Println();

	// A.24.7. Pengisian Nilai Sub-Struct
	fmt.Println("# - A.24.7. Pengisian Nilai Sub-Struct");
		var resultExampleSeven_SectionOne = exampleSeven{name: "DEDE", age: 25}
		var derivativeExampleSeven = derivativeExampleSeven{exampleSeven: resultExampleSeven_SectionOne, grade: 22}

		fmt.Println("name  :", derivativeExampleSeven.exampleSeven.name)
		fmt.Println("age   :", derivativeExampleSeven.exampleSeven.age)
		fmt.Println("grade :", derivativeExampleSeven.grade)
        fmt.Println();

	// A.24.8. Anonymous Struct
	fmt.Println("# - A.24.8. Anonymous Struct");
		var derivativeExampleEight = struct {
			exampleEight
			grade int
		}{}
		derivativeExampleEight.exampleEight = exampleEight{"wick", 21}
		derivativeExampleEight.grade = 2

		fmt.Println("name  :", derivativeExampleEight.exampleEight.name)
		fmt.Println("age   :", derivativeExampleEight.exampleEight.age)
		fmt.Println("grade :", derivativeExampleEight.grade)

		fmt.Println("-------------------------------------------------");
		// anonymous struct tanpa pengisian property
		var exEightSectionOne = struct {
			exampleEight
			grade int
		}{}

		// anonymous struct dengan pengisian property
		var exEightSectionTwo = struct {
			exampleEight
			grade int
		}{
			exampleEight: exampleEight{"wick", 21},
			grade:  2,
		}
		fmt.Println(exEightSectionOne.exampleEight)
		fmt.Println("name : ", exEightSectionTwo.exampleEight.name)
		fmt.Println("age  : ", exEightSectionTwo.exampleEight.age)
        fmt.Println();

	// A.24.9. Kombinasi Slice & Struct
	fmt.Println("# - A.24.9. Kombinasi Slice & Struct");
		var exampleNine = []exampleNine{
			{name: "Wick", age: 23},
			{name: "Ethan", age: 23},
			{name: "Bourne", age: 22},
		}

		for _, student := range exampleNine {
			fmt.Println(student.name, "age is", student.age)
		}
        fmt.Println();

	// A.24.10. Inisialisasi Slice Anonymous Struct
	fmt.Println("# - A.24.10. Inisialisasi Slice Anonymous Struct");
		var exampleTen = []struct{
			exampleTen
			grade int
		} {
			{exampleTen: exampleTen{"wick", 21}, grade: 2},
			{exampleTen: exampleTen{"ethan", 22}, grade: 3},
			{exampleTen: exampleTen{"bond", 21}, grade: 3},
		}
		for _, student := range exampleTen {
			fmt.Println(student)
		}
        fmt.Println();

	// A.24.11. Deklarasi Anonymous Struct Menggunakan Keyword var
	fmt.Println("# - A.24.11. Deklarasi Anonymous Struct Menggunakan Keyword var");
		// hanya deklarasi
		var derivativeExampleEleven_SectionOne struct {
			exampleEleven
			grade  int
		}
		derivativeExampleEleven_SectionOne.exampleEleven = exampleEleven{"DEDE", 21}
		derivativeExampleEleven_SectionOne.grade = 2
		fmt.Println(derivativeExampleEleven_SectionOne)
		fmt.Println("-------------------------------------------------");
		// deklarasi sekaligus inisialisasi
		var derivativeExampleEleven_SectionTwo = struct {
			exampleEleven
			grade  int
		} {
			exampleEleven{"HEND", 25},
			21,
		}
		fmt.Println(derivativeExampleEleven_SectionTwo)
        fmt.Println();

	// A.24.12. Nested struct
	fmt.Println("# - A.24.12. Nested struct");
		// var resultExampleTwelve_SectionOne = exampleTwelve{
		// 										{
		// 											derivativeExampleTwelve.name: "dede", 
		// 											derivativeExampleTwelve.age: 21,
		// 										},
		// 										grade: 21,
		// 										hobbies: {"a", "b"},
		// 									}
		// fmt.Println(resultExampleTwelve_SectionOne)

		exampleTwelve := exampleTwelve{
							grade : 21,
							hobbies: []string{"a", "b"},
						}
		exampleTwelve.derivativeExampleTwelve.name = "DEDE"
		exampleTwelve.derivativeExampleTwelve.age = 21
		fmt.Println("GRADE 		:", exampleTwelve.grade)
		fmt.Println("HOBBIES 	:", exampleTwelve.hobbies)
		fmt.Println("NAME 		:", exampleTwelve.derivativeExampleTwelve.name)
		fmt.Println("AGE 		:", exampleTwelve.derivativeExampleTwelve.age)
		fmt.Println();

	// A.24.13. Deklarasi Dan Inisialisasi Struct Secara Horizontal
	fmt.Println("# - A.24.13. Deklarasi Dan Inisialisasi Struct Secara Horizontal");
		var resultExampleThirTeen_SectionOne = struct { name string; age int } { age: 22, name: "DEDE" }
		var resultExampleThirTeen_SectionTwo = struct { name string; age int } { "RIAN", 24 }
		fmt.Println(resultExampleThirTeen_SectionOne)
		fmt.Println(resultExampleThirTeen_SectionTwo)
        fmt.Println();

	// A.24.14. Tag property dalam struct
	fmt.Println("# - A.24.14. Tag property dalam struct");
		fmt.Println();

	// A.24.15. Type Alias
	fmt.Println("# - A.24.15. Type Alias");
		var p1 = exampleFiveTeen{"DEDE", 22}
		var p2 = derivativeExampleFiveTeen{"HEND", 22}
		fmt.Println(p1)
		fmt.Println(p2)
		fmt.Println("-------------------------------------------------");
		derivativeExampleFiveTeen := exampleFiveTeen{"DEDE", 22}
		fmt.Println(exampleFiveTeen(derivativeExampleFiveTeen))
		// exampleFiveTeen := derivativeExampleFiveTeen{"HEND", 22}
		// fmt.Println(derivativeExampleFiveTeen(exampleFiveTeen))

		fmt.Println("-------------------------------------------------");
		type People1 struct {
			name string
			age  int
		}
		type People2 = struct {
			name string
			age  int
		}
		var a = People1{"DEDE", 11}
		var b = People2{"HEND", 22}
		fmt.Println(a)
		fmt.Println(b)

		fmt.Println("-------------------------------------------------");
		// TYPE NUMBER ADALAH ALIAS DARI TYPE INTEGER
		type Number = int
		var num Number = 12
		fmt.Println(num)

		type Huruf = string
		var valString Huruf = "HALLO"
		fmt.Println(valString);

		type ArrayInt = []int
		var arrayInt ArrayInt = []int{1, 2, 3}
		fmt.Println(arrayInt);
		fmt.Println();
}
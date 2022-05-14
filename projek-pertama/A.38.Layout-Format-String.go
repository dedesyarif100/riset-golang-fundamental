package main

import (
	"fmt"
)

func main() {
	// A.38.1. Persiapan
	fmt.Println("# - A.38.1. Persiapan");
	type student struct {
		name        string
		height      float64
		age         int32
		isGraduated bool
		hobbies     []string
	}

	var data = student{
		name:        "wick",
		height:      182.5,
		age:         26,
		isGraduated: false,
		hobbies:     []string{"eating", "sleeping"},
	}

	// A.38.2. Layout Format %b
	fmt.Println("# - A.38.2. Layout Format %b");
	fmt.Printf("%b\n", data.age)

	// A.38.3. Layout Format %c
	fmt.Println("# - A.38.3. Layout Format %c");
	fmt.Printf("%c\n", 1400)
	// ո
	fmt.Printf("%c\n", 1235)
	// ӓ

	// A.38.4. Layout Format %d
	fmt.Println("# - A.38.4. Layout Format %d");
	fmt.Printf("%d\n", data.age)
	// 26

	// A.38.5. Layout Format %e atau %E
	fmt.Println("# - A.38.5. Layout Format %e atau %E");
	fmt.Printf("%e\n", data.height)
	// 1.825000e+02

	fmt.Printf("%E\n", data.height)
	// 1.825000E+02

	// A.38.6. Layout Format %f atau %F
	fmt.Println("# - A.38.6. Layout Format %f atau %F");
	fmt.Printf("%f\n", data.height)
	// 182.500000
	fmt.Printf("%.9f\n", data.height)
	// 182.500000000
	fmt.Printf("%.2f\n", data.height)
	// 182.50
	fmt.Printf("%.f\n", data.height)
	// 182

	// A.38.7. Layout Format %g atau %G
	fmt.Println("# - A.38.7. Layout Format %g atau %G");
	fmt.Printf("%e\n", 0.123123123123)
	// 1.231231e-01
	fmt.Printf("%f\n", 0.123123123123)
	// 0.123123
	fmt.Printf("%g\n", 0.123123123123)
	// 0.123123123123

	fmt.Printf("%g\n", 0.12)
	// 0.12
	fmt.Printf("%.5g\n", 0.12)
	// 0.12

	// A.38.8. Layout Format %o
	fmt.Println("# - A.38.8. Layout Format %o");
	fmt.Printf("%o\n", data.age)
	// 32

	// A.38.9. Layout Format %p
	fmt.Println("# - A.38.9. Layout Format %p");
	fmt.Printf("%p\n", &data.name)
	// 0x2081be0c0

	// A.38.10. Layout Format %q
	fmt.Println("# - A.38.10. Layout Format %q");
	fmt.Printf("%q\n", `" name \ height "`)
	// "\" name \\ height \""

	// A.38.11. Layout Format %s
	fmt.Println("# - A.38.11. Layout Format %s");
	fmt.Printf("%s\n", data.name)
	// wick

	// A.38.12. Layout Format %t
	fmt.Println("# - A.38.12. Layout Format %t");
	fmt.Printf("%t\n", data.isGraduated)
	// false

	// A.38.13. Layout Format %T
	fmt.Println("# - A.38.13. Layout Format %T");
	fmt.Printf("%T\n", data.name)
	// string
	fmt.Printf("%T\n", data.height)
	// float64
	fmt.Printf("%T\n", data.age)
	// int32
	fmt.Printf("%T\n", data.isGraduated)
	// bool
	fmt.Printf("%T\n", data.hobbies)
	// []string

	// A.38.14. Layout Format %v
	fmt.Println("# - A.38.14. Layout Format %v");
	fmt.Printf("%v\n", data)
	// {wick 182.5 26 false [eating sleeping]}

	// A.38.15. Layout Format %+v
	fmt.Println("# - A.38.15. Layout Format %+v");
	fmt.Printf("%+v\n", data)
	// {name:wick height:182.5 age:26 isGraduated:false hobbies:[eating sleeping]}

	// A.38.16. Layout Format %#v
	fmt.Println("# - A.38.16. Layout Format %#v");
	fmt.Println("%#v\n", data)
	var exampleSixTeen = struct {
		name   string
		height float64
	} {
		name:   "wick",
		height: 182.5,
	}

	fmt.Printf("%#v\n", exampleSixTeen);
	// struct { name string; height float64 }{name:"wick", height:182.5}

	// A.38.17. Layout Format %x atau %X
	fmt.Println("# - A.38.17. Layout Format %x atau %X");
	fmt.Printf("%x\n", data.age)
	// 1a
	var d = data.name
	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])
	// 7769636b
	fmt.Printf("%x\n", d)
	// 7769636b

	// A.38.18. Layout Format %%
	fmt.Println("# - A.38.18. Layout Format %%");
	fmt.Printf("%%\n")
	// %
}
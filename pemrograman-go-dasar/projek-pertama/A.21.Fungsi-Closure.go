package main

import (
	"reflect"
	"fmt"
)

func main() {
	// A.21.1. Closure Disimpan Sebagai Variabel >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	fmt.Println("# - A.21.1. Closure Disimpan Sebagai Variabel");
		var exampleOne = func(n []int) (int, int) {
			var min, max int
			for i, e := range n {
				switch {
					case i == 0:
						max, min = e, e
					case e > max:
						max = e
					case e < min:
						min = e
				}
			}
			return min, max
		}

		var numbersExampleOne = []int{2, 3, 4, 3, 4, 2, 3}
		var minExampleOne, maxExampleOne = exampleOne(numbersExampleOne)
		fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbersExampleOne, minExampleOne, maxExampleOne)

		// A.21.1.1. Penggunaan Template String %v
		// fmt.Println("# - A.21.1.1. Penggunaan Template String %v");
		println()

	// A.21.2. Immediately-Invoked Function Expression (IIFE) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	fmt.Println("# - A.21.2. Immediately-Invoked Function Expression (IIFE)");
		var numbersExampleTwo = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
		var newNumbers = func(min int) []int {
			// fmt.Println(min);
			var value []int
			for _, val := range numbersExampleTwo {
				if val < min {
					continue
				}
				value = append(value, val)
			}
			return value
		}(3)

		fmt.Println("original number :", numbersExampleTwo)
		fmt.Println("filtered number :", newNumbers)

		// Closure bisa juga dengan gaya manifest typing, caranya dengan menuliskan skema closure-nya sebagai tipe data. Contoh:
		var closure (func (string, int, []string) int)
		closure = func (a string, b int, c []string) int {
			val := len(a) + b + len(c[0])
			return val
		}
		val := closure("TES", 5, []string{"DEDE"})

		fmt.Println("CLOSURE :", val)
		println()

	// A.21.3. Closure Sebagai Nilai Kembalian >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	fmt.Println("# - A.21.3. Closure Sebagai Nilai Kembalian");
		var maxExampleThree = 3
		var numbersExampleThree = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
		var howMany, getNumbers = exampleThree(numbersExampleThree, maxExampleThree)
		fmt.Println(reflect.TypeOf(howMany))
		fmt.Println(reflect.TypeOf(getNumbers))
		// var theNumbers = getNumbers()

		fmt.Println("numbers\t:", numbersExampleThree)
		fmt.Printf("find \t: %d\n", maxExampleThree)
		fmt.Println("found \t:", howMany)    // 9
		fmt.Println("value \t:", getNumbers()) // [2 3 0 3 2 0 2 0 3]
		println()

	// EXAMPLE FOUR >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	fmt.Println("# - EXAMPLE FOUR")
		numbersExampleFour := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		maxExampleFour := 8
		valA, valB := exampleFour(numbersExampleFour, maxExampleFour)
		fmt.Println("FOUND \t:", valA)
		fmt.Println("VALUE \t:", valB())
		println()
}

func exampleThree(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, val := range numbers {
		if val <= max {
			res = append(res, val)
		}
	}
	return len(res), func() []int {
		return res
	}
}




func exampleFour(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, val := range numbers {
		if val <= max {
			res = append(res, val)
		}
	}
	return len(res), func() []int {
		return res
	}
}
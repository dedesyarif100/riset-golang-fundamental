package main

import (
	"fmt"
	"reflect"
)

func main() {
	// A.23.1. Penerapan Pointer
	fmt.Println("# - A.23.1. Penerapan Pointer");
		var numberA int = 4
		var numberB *int = &numberA
		fmt.Println(reflect.TypeOf(numberB))
		fmt.Println("numberA (value)   :", numberA)  // 4
		fmt.Println("numberA (address) :", &numberA) // 0xc20800a220
		fmt.Println("numberB (value)   :", *numberB) // 4
		fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

		fmt.Println("-------------------------------")
		numberA = 5

		fmt.Println("numberA (value)   :", numberA)
		fmt.Println("numberA (address) :", &numberA)
		fmt.Println("numberB (value)   :", *numberB)
		fmt.Println("numberB (address) :", numberB)
		println()

	// A.23.3. Parameter Pointer
	fmt.Println("# - A.23.3. Parameter Pointer");
		var number = 4
		fmt.Println("before :", number) // 4
		exampleThree(&number, 10)
		fmt.Println("after  :", number) // 10
		println()

	// EXAMPLE ANY
	fmt.Println("# - EXAMPLE ANY");
		var valExOne interface{} = "DEDE"
		var valExTwo *any = &valExOne
		fmt.Println("VAL EX ONE 		:", valExOne)
		fmt.Println("VAL EX ONE ADDRESS 	:", &valExOne)
		fmt.Println("VAL EX TWO 		:", *valExTwo)
		fmt.Println("VAL EX TWO ADDRESS 	:", valExTwo)
		fmt.Println(reflect.TypeOf(valExTwo))
		fmt.Println("--------------------------------------")
		valExOne = []int{1, 2, 3, 4}
		fmt.Println("VAL EX ONE 		:", valExOne)
		fmt.Println("VAL EX ONE ADDRESS 	:", &valExOne)
		fmt.Println("VAL EX TWO 		:", *valExTwo)
		fmt.Println("VAL EX TWO ADDRESS 	:", valExTwo)
		fmt.Println("--------------------------------------")
		ExampleAny(&valExOne, "TES")
		fmt.Println("VAL EX ONE 		:", valExOne)
		fmt.Println("VAL EX ONE ADDRESS 	:", &valExOne)
		println()

	// EXAMPLE ARRAY
	fmt.Println("# - EXAMPLE ARRAY")
		arrNumbersOne := []interface{}{1, 2}
		arrNumbersTwo := &arrNumbersOne
		fmt.Println(arrNumbersOne)
		fmt.Println(&arrNumbersOne)
		fmt.Println(*arrNumbersTwo)
		fmt.Println(arrNumbersTwo)
		fmt.Println("------------------------------------")
		arrNumbersOne = []any{3, 4}
		fmt.Println(arrNumbersOne)
		fmt.Println(&arrNumbersOne)
		fmt.Println(*arrNumbersTwo)
		fmt.Println(arrNumbersTwo)
		fmt.Println("------------------------------------")
		println()

	// EXAMPLE MULTIPLE ARRAY
	fmt.Println("# - EXAMPLE MULTIPLE ARRAY")
		arrMultipleOne := [][3]any{
							{1, 2, 3},
							{4, 5, 6},
							{"A", 1, 2.5},
						}
		arrMultipleTwo := &arrMultipleOne
		fmt.Println(arrMultipleOne)
		fmt.Println(&arrMultipleOne)
		fmt.Println(*arrMultipleTwo)
		fmt.Println(arrMultipleTwo)

		fmt.Println("-------------------------------")
		arrMultipleOne = [][3]any{
							{"A", "B", "C"},
							{11, 22, 33},
						}
		fmt.Println(arrMultipleOne)
		fmt.Println(&arrMultipleOne)
		fmt.Println(*arrMultipleTwo)
		fmt.Println(arrMultipleTwo)
		println()
}

func exampleThree(original *int, value int) {
    *original = value
}

func ExampleAny(original *interface{}, value string) {
	*original = value
}
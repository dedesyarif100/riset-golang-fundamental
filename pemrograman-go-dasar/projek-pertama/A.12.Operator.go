package main

import (
	"fmt"
	// "reflect"
	// "math"
)

func main() {
	// A.12.1. Operator Aritmatika
	fmt.Println("# - A.12.1. Operator Aritmatika");
		var value1 = (((2 + 6) % 3) * 4 - 2) / 3;
		fmt.Println(value1);
		println()

	// A.12.2. Operator Perbandingan
	fmt.Println("# - A.12.2. Operator Perbandingan");
		var value2 = (((2 + 6) % 3) * 4 - 2) / 3;
		var isEqual = (value2 == 2);
		fmt.Printf("nilai %d (%t) \n", value2, isEqual);
		println()

	// A.12.3. Operator Logika
	fmt.Println("# - A.12.3. Operator Logika");
		var resultThreeSecOne = false && true;
		fmt.Printf("false && true : \t(%t) \n", resultThreeSecOne);

		var resultThreeSecTwo = false || true;
		fmt.Printf("false || true : \t(%t) \n", resultThreeSecTwo);

		var resultThreeSecThree = !false;
		fmt.Printf("!false	: \t\t(%t) \n", resultThreeSecThree);
		println()
}

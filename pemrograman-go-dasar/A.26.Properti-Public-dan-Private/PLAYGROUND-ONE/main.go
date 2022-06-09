package main

import (
	"fmt"
	SubOne "PLAYGROUND-ONE/subfolderone"
)


func main() {
	fmt.Println("# - EXAMPLE ONE")
		var ExampleOne = SubOne.ExampleOne{"NAME", 22, []int{1, 2, 3}, []string{"A", "B"}}
		SubOne.ApplyAllSubFolderOne([]int{1, 2, 3})
		fmt.Println("EX ONE 	:", SubOne.ExOne)
		fmt.Println("NAME 	:", ExampleOne.Name)
		fmt.Println("AGE 	:", ExampleOne.Age)
		fmt.Println("NUMBER 	:", ExampleOne.Number)
		fmt.Println("ABJAD 	:", ExampleOne.Abjad)
		println()

	fmt.Println("# - EXAMPLE TWO")
		var ExampleTwo = SubOne.ExampleTwo{ExampleOne, "DEDE", 22}
		fmt.Println("EX TWO 			:", SubOne.ExTwo)
		fmt.Println("EXAMPLE ONE (NAME) 	:", ExampleTwo.ExampleOne.Name)
		fmt.Println("EXAMPLE ONE (AGE) 	:", ExampleTwo.ExampleOne.Age)
		fmt.Println("EXAMPLE ONE (NUMBER) 	:", ExampleTwo.ExampleOne.Number)
		fmt.Println("EXAMPLE ONE (ABJAD) 	:", ExampleTwo.ExampleOne.Abjad)
		println()

	fmt.Println("# - EXAMPLE THREE")
		fmt.Println("EX THREE 	:", SubOne.ExThree)
		fmt.Printf("NAME 		:%s\n", SubOne.ExampleThree.Name)
		fmt.Printf("AGE  		:%d\n", SubOne.ExampleThree.Age)

}
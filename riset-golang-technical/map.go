package main

import (
	"fmt"
)

func main() {
	case6()
}

type dataTypeInt int

type dataTypeString string

func case1() {
	map1 := map[dataTypeString]dataTypeInt{"A": 10, "B": 10, "C": 10}
	map2 := map[dataTypeString]dataTypeInt{
		"A": 10,
		"B": 10,
	}
	map3 := map[dataTypeInt]dataTypeInt{
		1: 10,
		2: 10,
	}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	for key, val := range map1 {
		fmt.Println(key, ":\t", val)
	}
	println()
}

func case2() {
	map1 := make(map[string]int)
	// map1 := *new(map[string]int)
	// map1 := map[string]int{}
	map1["A"] = 1
	map1["B"] = 2
	map1["C"] = 3
	map1["D"] = 4
	map1["E"] = 5
	fmt.Println(map1)
}

func case3() {
	var jumlahMap int
	fmt.Print("JUMLAH MAP : ")
	fmt.Scanln(&jumlahMap)

	arr := []any{}

	var i = 1
	for i <= jumlahMap {
		arr = append(arr, i)
		i++
	}

	var variable []interface{}
	var text string
	for j := 1; j <= len(arr); j++ {
		fmt.Println(j)
		fmt.Print("INPUT", j, " : ")
		fmt.Scanln(&text)
		variable = append(variable, text)
	}
	
	fmt.Println(variable...)
	// fmt.Println(variable)
	// return

	arrayObject := []map[string]any{}
	for k := 0; k < len(variable); k++ {
		arrayObject = append(arrayObject, map[string]any{"name": variable[k]})
	}

	fmt.Println("---------------------------")
	for _, val := range arrayObject {
		for _, val2 := range val {
			fmt.Println(val2)
		}
	}
}

func case4() {
	map1 := []map[string]int{
		map[string]int{"A": 1, "B": 4},
		map[string]int{"A": 2, "B": 5},
		map[string]int{"A": 3, "B": 6},
	}
	for _, val1 := range map1 {
		for _, val2 := range val1 {
			fmt.Print(val2, ", ")
		}
	}
	println()
}

func case5() {
	var exampleFive = []map[string]int{
		{"A":1, "B":2, "C":3},
		{"A":4, "B":5, "C":6},
	}
	var value, isExist = exampleFive[1]["A"]
	fmt.Println(isExist)
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}
	fmt.Println();
}

func case6() {
	map1 := [1][2][2]map[int]int{
				{
					{
						{
							1: 11,
							2: 22,
						},
						{
							3: 33,
							4: 44,
						},
					},
					{
						{
							5: 55,
							6: 66,
						},
						{
							7: 77,
							8: 88,
						},
					},
				},
			}
	for _, val1 := range map1 {
		for _, val2 := range val1 {
			for _, val3 := range val2 {
				for _, val4 := range val3 {
					fmt.Print(val4, ", ")
				}
			}
		}
	}
	println()
}

func init() {
	fmt.Println("MASUK PERTAMA KALI")
}

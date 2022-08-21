package main

import "fmt"

func main() {
	case3()
}

func case1() {
	var input string
	fmt.Print("INPUT : ")
	fmt.Scanln(&input)
	if a := "1"; a == input || input == "satu" {
		fmt.Println("SATU")
	} else if a := "2"; a == input || input == "dua" {
		fmt.Println("DUA")
	} else if a := "3"; a == input || input == "tiga" {
		fmt.Println("TIGA")
	} else if a := "4"; a == input || input == "empat" {
		fmt.Println("EMPAT")
	} else {
		fmt.Println(false)
	}
}

func case2() {
	var start int
	var end int
	fmt.Print("START 	: ")
	fmt.Scanln(&start)
	fmt.Print("END 	: ")
	fmt.Scanln(&end)
	if start < end {
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		fmt.Println("--------------------------")
		for key, val := range a {
			if a[key] >= start && a[key] <= end {
				fmt.Print(val, ", ")
			}
		}
		println()
	} else {
		fmt.Println("NUMBER START MUST MORE LARGE")
	}
}

func case3() {
	var input int
	fmt.Print("INPUT : ")
	fmt.Scanln(&input)
	arr := []int{}
	var i = 1
	for i <= input {
		arr = append(arr, i)
		i++
	}
	fmt.Println(arr)
}

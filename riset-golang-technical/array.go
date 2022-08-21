package main

import "fmt"

func main() {
	case7()
}

type tipeDataInt int

func case1() {
	// arr := [...]int{1, 2, 3, 4}
	// fmt.Println(arr)
	// a := 3
	// b := 3

	// arr := [3][3]int{
	// 		[3]int{1, 2, 3},
	// 		[3]int{1, 2, 3},
	// 	}
	// fmt.Println(arr)

	arr := []tipeDataInt{1, 2, 3, 4, 5}
	fmt.Println(arr)
}

func case2() {
	// a := []any{1, 2, 3, 4, 5}
	// fmt.Println(a)
	
	a := []tipeDataInt{1, 2, 3, 4, 5}
	b := &a[1]
	fmt.Println(a)
	fmt.Println(*b)
}

func case3() {
	a := 1
	b := 2.2
	c := "c"
	d := []any{a, b, c}
	arr := []interface{}{a, b, c, d}
	fmt.Println(arr)
	fmt.Println("PANJANG ARRAY :", len(arr))
}

func case4() {
	arr := [5]tipeDataInt{1, 2, 3, 4, 5}
	for _, val := range arr {
		fmt.Print(val,", ")
	}
	println()
}

func case5() {
	var size = 4
	arr := make([]tipeDataInt, size)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	fmt.Println(arr)
}

func case6() {
	data := [][][]int{
			{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
			},
			{
				[]int{7, 8, 9},
				[]int{10, 11, 12},
			},
		}

	tmp := data[0][0]
	data[0][0] = data[0][1]
	data[0][1] = tmp
	for _, val1 := range data {
		for _, val2 := range val1 {
			for _, val3 := range val2 {
				fmt.Print(val3, ", ")
			}
		}
	}
	println()
}

func case7() {
	data := [][][][][]int{
		{
			{
				{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
				},
			},
		},
		{
			{
				{
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
			},
		},
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			for k := 0; k < len(data[i][j]); k++ {
				for l := 0; l < len(data[i][j][k]); l++ {
					for m := 0; m < len(data[i][j][k][l]); m++ {
						fmt.Print(data[i][j][k][l][m], " | ")
					}
				}
			}
		}
	}
	println()
}

func init() {
	fmt.Println("MASUK PERTAMA KALI")
}
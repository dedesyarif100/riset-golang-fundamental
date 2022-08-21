package main

import "fmt"

func main() {
	case2()
}

type dataTypeInt int

func case1() {
	var start int
	var end int
	var input int
	var arr = []int{}
	fmt.Print("START : ")
	fmt.Scanln(&start)
	fmt.Print("END : ")
	fmt.Scanln(&end)

	fmt.Print("INPUT : ")
	fmt.Scanln(&input)
	
	i := 1
	for i <= input {
		arr = append(arr, i)
		i++
	}

	split := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(arr[start:end])
	fmt.Println(arr[:])
	fmt.Println(arr[split[2]:])
	fmt.Println(arr[:split[2]])
}

func case2() {
	var data int
	fmt.Print("JUMLAH DATA : ")
	fmt.Scanln(&data)

	var start int
	fmt.Print("START	: ")
	fmt.Scanln(&start)

	var end int
	fmt.Print("END	: ")
	fmt.Scanln(&end)

	arr := []int{}
	for i := 1; i <= data; i++ {
		arr = append(arr, i)
	}
	if end < start {
		fmt.Println("ENDING TIDAK BOLEH LEBIH BESAR DARI START")
		return
	}
	if end > len(arr) {
		fmt.Println("ENDING TERLALU BESAR")
		return
	}

	fmt.Println(arr[start:end])
	fmt.Println(len(arr[start:end]))
	fmt.Println(cap(arr[start:end]))
}

func case3() {
	var input int
	fmt.Print("INPUT : ")
	fmt.Scanln(&input)
	make := make([]int, 4)
	arr := []int{1, 2, 3, 4, 5}
	data := copy(make, arr)
	fmt.Println(data)
}

func case4() {
	var jmlData int
	fmt.Print("JUMLAH DATA : ")
	fmt.Scanln(&jmlData)

	var start int
	fmt.Print("START 	: ")
	fmt.Scanln(&start)

	var end int
	fmt.Print("END 	: ")
	fmt.Scanln(&end)

	arr := []int{}
	for i := 1; i <= jmlData; i++ {
		arr = append(arr, i)
	}
	if end < start {
		fmt.Println("ENDING TIDAK BOLEH LEBIH BESAR DARI START")
		return
	}
	if end > len(arr) {
		fmt.Println("ENDING TERLALU BESAR")
		return
	}

	val := arr[start:end:jmlData]
	fmt.Println(val)
}

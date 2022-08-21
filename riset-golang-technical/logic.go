package main

import "fmt"

func main() {
	case4()
}

func case1() {
	a := 1
	b := 12
	for c := a; c <= b; c++ {
		if c == 3 || c == 5 {
			fmt.Println("----")
		} else {
			fmt.Println(c)
		}
	}
}

func case2() {
	var input int
	fmt.Print("INPUT : ")
	fmt.Scanln(&input)
	fmt.Println("OUTPUT :", input)
	i := 1
	for i <= input {
		fmt.Print(i,", ")
		i++
	}
	println()
}

func case3() {
	var input int
	fmt.Print("INPUT : ")
	fmt.Scanln(&input)
	var text string
	switch input {
		case 1:
			fmt.Print("INPUT : ")
			fmt.Scan(&text)
			switch text {
				case "a":
					fmt.Println("A")
				case "b":
					fmt.Println("B")
				case "c":
					fmt.Println("C")
			}
		case 2:
			fmt.Print("INPUT : ")
			fmt.Scan(&text)
			switch text {
				case "d":
					fmt.Println("D")
				case "e":
					fmt.Println("E")
				case "f":
					fmt.Println("F")
			}
		case 3:
			fmt.Print("INPUT : ")
			fmt.Scan(&text)
			switch text {
				case "g":
					fmt.Println("G")
				case "h":
					fmt.Println("H")
				case "i":
					fmt.Println("I")
			}
		case 4:
			fmt.Print("INPUT : ")
			fmt.Scan(&text)
			switch text {
				case "j":
					fmt.Println("J")
				case "k":
					fmt.Println("K")
				case "l":
					fmt.Println("L")
			}
		case 5:
			fmt.Print("INPUT : ")
			fmt.Scan(&text)
			switch text {
				case "m":
					fmt.Println("M")
				case "n":
					fmt.Println("N")
				case "o":
					fmt.Println("O")
			}
		case 6:
			fmt.Print("INPUT : ")
			fmt.Scan(&text)
			switch text {
				case "p":
					fmt.Println("P")
				case "q":
					fmt.Println("Q")
				case "r":
					fmt.Println("R")
			}
		case 7:
			fmt.Print("INPUT : ")
			fmt.Scan(&text)
			switch text {
				case "s":
					fmt.Println("S")
				case "t":
					fmt.Println("T")
				case "u":
					fmt.Println("U")
			}
		case 8:
			fmt.Print("INPUT : ")
			fmt.Scan(&text)
			switch text {
				case "v":
					fmt.Println("V")
				case "w":
					fmt.Println("W")
				case "x":
					fmt.Println("X")
			}
		default:
			fmt.Println("BILANGAN TIDAK ADA")
	}
}

func case4() {
	var input int
	fmt.Print("INPUT : ")
	fmt.Scanln(&input)
	switch input {
		case 1, 2, 3, 4:
			{
				fmt.Println("1 - 4")
				fmt.Println("1 - 4")
			}
			fallthrough
		case 5, 6, 7, 8:
			{
				fmt.Println("5 - 8")
				fmt.Println("5 - 8")
			}
		case 9, 10, 11, 12:
			fmt.Println("9 - 12")
		case 13, 14, 15, 16:
			fmt.Println("13 - 16")
		default:
			fmt.Println("OUT RANGE")
	}
}

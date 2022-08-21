package main

import "fmt"

func main() {
	case2()
}

func case1() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			start:
			for k := 1; k <= 5; k++ {
				for l := 1; l <= 5; l++ {
					if i == 3 {
						fmt.Println("----------------------------------------")
						break start
					}
					fmt.Print("matriks [", i, "][", j, "][", k, "][", l, "]", "\n")
				}
			}
		}
	}
	println()
}

func case2() {
	start:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				fmt.Println("----------------------------------------")
				continue start
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}

func case3() {
	i := 1;
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
}

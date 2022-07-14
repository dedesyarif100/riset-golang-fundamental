package main

import "fmt"

func main() {
	fmt.Println("# - EXAMPLE ONE")
		var a int = 5
		SKIP:
		for a < 10 {
			if a == 7 {
					a = a + 1
					goto SKIP
			}
			fmt.Printf("A : %d\n", a)
			a++
		}
		println()

	fmt.Println("# - EXAMPLE TWO")
		b := 10
		LOOP:
		for b <= 20 {
			if b == 15 {
				b += 1
				fmt.Println("CONTINUE")
				goto LOOP
			}
			fmt.Printf("VALUE : %d\n", b)
			b++
		}
		println()

	fmt.Println("# - EXAMPLE THREE")
		i := 0
		START:
		fmt.Println("here")
		for i < 10 {
			if i%2 == 0 {
				i++
				goto START
			}  
			fmt.Println(i)
			i++
		}
		println()
}

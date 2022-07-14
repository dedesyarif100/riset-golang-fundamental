package main

import (
	"fmt"
	// "math"
)

type Circle struct {
	a, b, c int `json:"tes2"`
	d, e, f func(int) (int, float64) `json:"tes3"`
	x, y, z float64 `json:"tes1"`
}

func main() {
	// PANGKAT
		fmt.Println("============ PANGKAT ===============")
		// a := 2.0
		// // b := 3
		// b := math.Pow(a, 4)
		// fmt.Println(b)

	// TES LOOPING
		fmt.Println("============ LOOPING ===============")
		a := 1
		for ; a <= 10; {
			fmt.Print(a, ", ")
			a++
		}
		println()

		b := 10
		for ; b >= 1; {
			fmt.Print(b, ", ")
			b--
		}
		println()

		i := 1
		for ; i <= 10; i++ {
			fmt.Print(i, ", ")
		}
		println()

		j := 10
		for ; j >= 1; j-- {
			fmt.Print(j, ", ")
		}
		println()

		for i := 1; i <= 10; {
			fmt.Print(i, ", ")
			i++
		}
		println()

		for i := 10; i >= 1; {
			fmt.Print(i, ", ")
			i--
		}
		println()

		for i := 1.5; i <= 10.5; i += 1 {
			fmt.Print(i, ", ")
		}
		println()

		for i, j := 1, 20; i <= j; i, j = i+1, j-1 {
			fmt.Print(j, ", ")
		}
		println()

		for a, b := 20, 1; a >= b; a, b = a-1, b+1 {
			fmt.Print(b, ", ")
		}
		println()

		for i := 1; ; i++ {
			fmt.Print(i, ", ")
			if i == 5 {
				break
			}
		}
		println()

		for i, j := 1, 5; i <= j; i++ {
			fmt.Print(i, ", ")
		}
		println()

		// Title("FEBY")
		fmt.Println(Plus())

	fmt.Println("============ SWITCH CASE ===============")
		valA := 1
		valB := 10
		switch {
		case valA < 5 && valB > 5:
			fmt.Println("TRUE")
		case valA > 5 && valB < 5:
			fmt.Println("FALSE")
		default :
			fmt.Println("DEFAULT")
		}

	fmt.Println("============ DEFINITION STRUCT ===============")
		// c := app.pool.Get().(*Ctx)
		CustomStruct()
}


// STRUCT
type DataA struct {
	Name DataB
}
type DataB struct {
	Age DataC
}
type DataC struct {
	Job DataD
}
type DataD struct {
	location string
	wife string
	friends []string
}
func (d *DataA) Method() {
	d.Name.Age.Job.location = "SURABAYA"
	d.Name.Age.Job.wife = "FEBBY"
	d.Name.Age.Job.friends = []string{"RIAN", "TRIS", "HEND"}
	fmt.Println(d.Name.Age.Job.friends)
}
func CustomStruct() {
	data := DataA{}
	data.Method()
}


// TES FUNCTION
type TypeInt int
func Map(val string, elem ...int) string {
	// fmt.Println(elem)
	return "tes"
}
// func tes(val string, elem ...TypeInt) string
func Title(a string) string {
	// fmt.Println(string(a[2]))
	b := []int{1, 2, 3, 4}
	return Map(a, b...)
}
func Plus() (a, b int) {
	c := 3
	d := 3

	a = c + 1
	b = d + 1
	return
}

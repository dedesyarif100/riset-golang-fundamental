package main

import (
	apply "fmt"
)


type Hitung interface {
	// luas() 		float64
	// keliling()	float64
}



type Persegi struct {
	sisi 		float64
}
func (p Persegi) luas() float64 {
	return p.sisi* 2
}
func (p Persegi) keliling() float64 {
	return p.sisi* 2
}
func (l Persegi) jariJari() float64 {
	return l.luas() + l.keliling()
}



type Lingkaran struct {
	diameter 	float64
}
func (l Lingkaran) luas() float64 {
	return l.diameter * 2
}
func (l Lingkaran) keliling() float64 {
	return l.diameter * 2
}
func (l Lingkaran) jariJari() float64 {
	return l.luas() + l.keliling()
}


func main() {
	apply.Println("-----------------------------")
		var bangunDatar Hitung
		bangunDatar = Persegi{10.0}
		apply.Println("LUAS 	:", bangunDatar.(Persegi).luas())
		apply.Println("LUAS 	:", bangunDatar.(Persegi).keliling())
		apply.Println("LUAS 	:", bangunDatar.(Persegi).jariJari())

		bangunDatar = Lingkaran{14.0}
		apply.Println("LUAS 	:", bangunDatar.(Lingkaran).luas())
		apply.Println("KELILING :", bangunDatar.(Lingkaran).keliling())
		apply.Println("KELILING :", bangunDatar.(Lingkaran).jariJari())
}
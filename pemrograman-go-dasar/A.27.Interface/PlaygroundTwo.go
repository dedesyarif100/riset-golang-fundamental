package main

import (
	"fmt"
	"math"
)


type Hitung2D interface {
	luas() float64
	keliling() float64
}
type Hitung3D interface {
	volume() float64	
}
type Calculate interface {
	Hitung2D
	Hitung3D
}


type Kubus struct {
	sisi float64
}
func (k *Kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}
func (k *Kubus) keliling() float64 {
	return k.sisi * 12
}
func (k *Kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}


func main() {
	var bangunDatar Calculate = &Kubus{6}
	fmt.Println("============= KUBUS")
	fmt.Println("LUAS		:", bangunDatar.luas())
	fmt.Println("KELILING	:", bangunDatar.keliling())
	fmt.Println("VOLUME		:", bangunDatar.volume())
	println()
}
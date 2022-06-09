package main

import (
	"fmt"
	"math"
)


// EMBEDED INTERFACE
type Hitung2d interface {
    luas() float64
    keliling() float64
}
type Hitung3d interface {
    volume() float64
}
type Calculate interface {
    Hitung2d
    Hitung3d
}
// EMBEDED INTERFACE


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
    var bangunRuang Calculate = &Kubus{6}
    fmt.Println("===== kubus")
    fmt.Println("luas      :", bangunRuang.luas())
    fmt.Println("keliling  :", bangunRuang.keliling())
    fmt.Println("volume    :", bangunRuang.volume())
}
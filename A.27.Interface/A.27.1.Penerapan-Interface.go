package main

import (
	f "fmt"
	"math"
)

// INTERFACE HITUNG
type hitung interface {
    luas() float64
    keliling() float64
}

// STRUCT LINGKARAN
type lingkaran struct {
    diameter float64
}
func (l lingkaran) jariJari() float64 {
    return l.diameter / 2
}
func (l lingkaran) luas() float64 {
	// f.Println("CEK DATA : ", math.Pow(5, 3))
    return math.Pi * math.Pow(l.jariJari(), 2)
}
func (l lingkaran) keliling() float64 {
    return math.Pi * l.diameter
}

// STRUCT PERSEGI
type persegi struct {
    sisi float64
}
func (p persegi) luas() float64 {
	// f.Println("CEK DATA : ", math.Pow(p.sisi, 2))
    return math.Pow(p.sisi, 2)
}
func (p persegi) keliling() float64 {
	// f.Println("CEK DATA : ", p.sisi * 4)
    return p.sisi * 4
}

func main() {
	// A.27.1. Penerapan Interface
	f.Println("# - A.27.1. Penerapan Interface");
    var bangunDatar hitung
    bangunDatar = persegi{10.0}
    f.Println("PERSEGI ----------- ")
    f.Println("luas      :", bangunDatar.luas())
    f.Println("keliling  :", bangunDatar.keliling())

    bangunDatar = lingkaran{14.0}
    f.Println("LINGKARAN ----------- ")
    f.Println("luas      :", bangunDatar.luas())
    f.Println("keliling  :", bangunDatar.keliling())
    f.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())
}
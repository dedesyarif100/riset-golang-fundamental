package main

import (
	f "fmt"
	"math"
)


// INTERFACE HITUNG
type Hitung interface {
    luas()      float64
    keliling()  float64
    tes()       int
}


// STRUCT PERSEGI
type Persegi struct {
    sisi    float64
    val     int
}
func (p Persegi) luas() float64 {
	// f.Println("CEK DATA : ", math.Pow(p.sisi, 2))
    return math.Pow(p.sisi, 2)
}
func (p Persegi) keliling() float64 {
	// f.Println("CEK DATA : ", p.sisi * 4)
    return p.sisi * 4
}
func (p Persegi) tes() int {
    return p.val
}


// STRUCT LINGKARAN
type Lingkaran struct {
    diameter float64
    val2     string
}
func (l Lingkaran) jariJari() float64 {
    return l.diameter / 2
}
func (l Lingkaran) luas() float64 {
	// f.Println("CEK DATA : ", math.Pow(5, 3))
    return math.Pi * math.Pow(l.jariJari(), 2)
}
func (l Lingkaran) keliling() float64 {
    return math.Pi * l.diameter
}
func (l Lingkaran) tes() int {
    return len(l.val2)
}


func main() {
	// A.27.1. Penerapan Interface
	f.Println("# - A.27.1. Penerapan Interface");
        var bangunDatar Hitung
        bangunDatar = Persegi{10.0, 1}
        f.Println("PERSEGI ----------- ")
        f.Println("luas      :", bangunDatar.luas())
        f.Println("keliling  :", bangunDatar.keliling())
        f.Println();

        bangunDatar = Lingkaran{14.0, "tes"}
        f.Println("LINGKARAN ----------- ")
        f.Println("luas      :", bangunDatar.luas())
        f.Println("keliling  :", bangunDatar.keliling())
        f.Println("jari-jari :", bangunDatar.(Lingkaran).jariJari())
        f.Println();
}
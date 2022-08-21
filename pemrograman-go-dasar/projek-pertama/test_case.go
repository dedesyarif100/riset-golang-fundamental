package main

import "fmt"

type Orang struct {
	namaDepan string
	namaBelakang string
}

func main() {
	var dia Orang = Orang{namaDepan: "Agung", namaBelakang: "Wibowo"}
	fmt.Println(dia.namaLengkap())
}

func (f *Orang) namaLengkap() string {
	return f.namaDepan +" "+ f.namaBelakang
}
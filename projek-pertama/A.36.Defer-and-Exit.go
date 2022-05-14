package main

import (
	"fmt"
	"os"
)

func main() {
	// A.36.1. Penerapan keyword defer
	fmt.Println("# - A.36.1. Penerapan keyword defer");
    defer fmt.Println("INI DIJALANKAN PALING AKHIR");
    fmt.Println("selamat datang");

    orderSomeFood("pizza");
    orderSomeFood("burger");

	// A.36.2. Kombinasi defer dan IIFE
	fmt.Println("# - A.36.2. Kombinasi defer dan IIFE");
    numberTwoSecOne := 3;
    if numberTwoSecOne == 3 {
        fmt.Println("halo 1");
        defer fmt.Println("halo 3 -----------------");
    }
    fmt.Println("halo 2");
	// ----------------------------------------------------
	numberTwoSecTwo := 3;
    if numberTwoSecTwo == 3 {
        fmt.Println("hai 1");
        func() {
			defer fmt.Println("hai 3 -----------------");
        }();
    }
    fmt.Println("hai 2");

	// A.36.3. Penerapan Fungsi os.Exit();
	fmt.Println("# - A.36.3. Penerapan Fungsi os.Exit()");
    defer fmt.Println("halo");
    os.Exit(1);
    fmt.Println("DEDE");
}

// A.36.1. Penerapan keyword defer
func orderSomeFood(menu string) {
    defer fmt.Println("Terimakasih, silakan tunggu");
    if menu == "pizza" {
        fmt.Print("Pilihan tepat!", " ");
        fmt.Print("Pizza ditempat kami paling enak!", "\n");
        fmt.Println("----------------");
        return;
    }
    fmt.Println("Pesanan anda:", menu);
}
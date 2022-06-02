package main

import (
    "fmt"
    "time"
)

func main() {
	// A.42.1. Praktek
	fmt.Println("# - A.42.1. Praktek");
	fmt.Println("# - A.42.2. Kalkulasi Durasi Menggunakan time.Since()");
	fmt.Println("# - A.42.3. Method time.Duration");
		start := time.Now();
		time.Sleep(3 * time.Second);
		durationOne := time.Since(start);
		fmt.Println("time elapsed in seconds	: ", durationOne.Seconds());
		fmt.Println("time elapsed in minutes	: ", durationOne.Minutes());
		fmt.Println("time elapsed in hours	: ", durationOne.Hours());
		fmt.Println();

	// A.42.4. Kalkulasi Durasi Antara 2 Objek Waktu
	fmt.Println("# - A.42.4. Kalkulasi Durasi Antara 2 Objek Waktu");
		t1 := time.Now();
		time.Sleep(3 * time.Second);
		t2 := time.Now();

		durationFour := t2.Sub(t1);

		fmt.Println("time elapsed in seconds 	:", durationFour.Seconds());
		fmt.Println("time elapsed in minutes	:", durationFour.Minutes());
		fmt.Println("time elapsed in hours	:", durationFour.Hours());
		fmt.Println();

	// A.42.5. Konversi Angka ke time.Duration
	fmt.Println("# - A.42.5. Konversi Angka ke time.Duration");
		fmt.Println(12 * time.Minute);
		fmt.Println(65 * time.Hour);
		fmt.Println(150000 * time.Millisecond);
		fmt.Println(45 * time.Microsecond);
		fmt.Println(233 * time.Nanosecond);

		var n time.Duration = 5;
		durationFive := n * time.Microsecond;
		fmt.Println(durationFive);
}
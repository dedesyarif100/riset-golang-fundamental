package main

import (
	"fmt"
	"encoding/base64"
)

func main() {
	fmt.Println("# - A.46.1. Penerapan Fungsi EncodeToString() & DecodeString()");
		DataOne := "DEDE SYARIFUDIN HIDAYAT"
		EncodedStringExOne := base64.StdEncoding.EncodeToString([]byte(DataOne))
		fmt.Println("ENCODED :", EncodedStringExOne)

		DecodedByteExOne, _ := base64.StdEncoding.DecodeString(EncodedStringExOne)
		DecodedStringExOne := string(DecodedByteExOne)
		fmt.Println("DECODED :", DecodedStringExOne)
		println()

	fmt.Println("# - A.46.2. Penerapan Fungsi Encode() & Decode()");
		DataTwo := "DEDE SYARIFUDIN HIDAYAT"
		EncodedExTwo := make([]byte, base64.StdEncoding.EncodedLen(len(DataTwo)))
		base64.StdEncoding.Encode(EncodedExTwo, []byte(DataTwo))
		EncodedStringExTwo := string(EncodedExTwo)
		fmt.Println("ENCODED :", EncodedStringExTwo)

		DecodedExTwo := make([]byte, base64.StdEncoding.DecodedLen(len(EncodedExTwo)))
		base64.StdEncoding.Decode(DecodedExTwo, EncodedExTwo)
		DecodedStringExTwo := string(EncodedExTwo)
		fmt.Println("DECODED :", DecodedStringExTwo)
		println()

	fmt.Println("# - A.46.3. Encode & Decode Data URL");
		DataThree := "https://dasarpemrogramangolang.novalagung.com/A-channel-select.html"
		EncodedStringExThree := base64.URLEncoding.EncodeToString([]byte(DataThree))
		fmt.Println("ENCODED URL :", EncodedStringExThree)

		DecodedByteExThree, _ := base64.URLEncoding.DecodeString(EncodedStringExThree)
		DecodedStringExThree := string(DecodedByteExThree)
		fmt.Println("DECODED URL :", DecodedStringExThree)
		println()
}
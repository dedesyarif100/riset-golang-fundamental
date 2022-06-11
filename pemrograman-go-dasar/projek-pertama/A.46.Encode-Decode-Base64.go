package main

import (
    "fmt"
	"encoding/base64"
)

func main() {
	// A.46.1. Penerapan Fungsi EncodeToString() & DecodeString()
	fmt.Println("# - A.46.1. Penerapan Fungsi EncodeToString() & DecodeString()");
		var dataExOne = "john wick"
		var encodedStringExOne = base64.StdEncoding.EncodeToString([]byte(dataExOne))
		fmt.Println("ENCODED :", encodedStringExOne)

		var decodedByteExOne, _ = base64.StdEncoding.DecodeString(encodedStringExOne)
		var decodedStringExOne = string(decodedByteExOne)
		fmt.Println("DECODED :", decodedStringExOne)
		fmt.Println();

	// A.46.2. Penerapan Fungsi Encode() & Decode()
	fmt.Println("# - A.46.2. Penerapan Fungsi Encode() & Decode()");
		var dataExTwo = "john wick"
		var encodedExTwo = make([]byte, base64.StdEncoding.EncodedLen(len(dataExTwo)))
		base64.StdEncoding.Encode(encodedExTwo, []byte(dataExTwo))
		var encodedStringExTwo = string(encodedExTwo)
		fmt.Println("ENCODED :", encodedStringExTwo)

		var decodedExTwo = make([]byte, base64.StdEncoding.DecodedLen(len(encodedExTwo)))
		var _, errExTwo = base64.StdEncoding.Decode(decodedExTwo, encodedExTwo)
		if errExTwo != nil {
			fmt.Println(errExTwo.Error())
		}
		var decodedStringExTwo = string(encodedExTwo)
		fmt.Println("DECODED :", decodedStringExTwo)
		fmt.Println();

	// A.46.3. Encode & Decode Data URL
	fmt.Println("# - A.46.3. Encode & Decode Data URL");
		var dataExThree = "https://kalipare.com/"
		var encodedStringExThree = base64.URLEncoding.EncodeToString([]byte(dataExThree))
		fmt.Println("ENCODED URL :", encodedStringExThree)

		var decodedByteExThree, _ = base64.URLEncoding.DecodeString(encodedStringExThree)
		var decodedStringExThree = string(decodedByteExThree)
		fmt.Println("DECODED URL :", decodedStringExThree)
		fmt.Println();
}
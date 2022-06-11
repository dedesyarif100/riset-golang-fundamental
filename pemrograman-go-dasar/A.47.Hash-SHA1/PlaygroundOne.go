package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	fmt.Println("# - A.47.1. Penerapan Hash SHA1");
		TextExOne := "DEDE SYARIF"
		sha := sha1.New()
		sha.Write([]byte(TextExOne))
		Encrypted := sha.Sum(nil)
		EncryptedString := fmt.Sprintf("%x", Encrypted)
		fmt.Println(EncryptedString)
		println()

	fmt.Println("# - A.47.2. Metode Salting Pada Hash SHA1");
		TextExTwo := "DEDE SYARIF"
		fmt.Printf("ORIGINAL : %s\n\n", TextExTwo)

		Hashed1, salt1 := DoHashUsingSalt(TextExTwo)
		fmt.Printf("HASHED 1 : %s\n\n", Hashed1)

		Hashed2, salt2 := DoHashUsingSalt(TextExTwo)
		fmt.Printf("HASHED 2 : %s\n\n", Hashed2)

		Hashed3, salt3 := DoHashUsingSalt(TextExTwo)
		fmt.Printf("HASHED 3 : %s\n\n", Hashed3)

		_, _, _ = salt1, salt2, salt3
}

func DoHashUsingSalt(text string) (string, string) {
	Salt := fmt.Sprintf("%d", time.Now().UnixNano())
	SaltedText := fmt.Sprintf("TEXT : %s, SALT : %s", text, Salt)
	fmt.Println(SaltedText)
	sha := sha1.New()
	sha.Write([]byte(SaltedText))
	Encrypted := sha.Sum(nil)Z
	return fmt.Sprintf("%x", Encrypted), Salt
}
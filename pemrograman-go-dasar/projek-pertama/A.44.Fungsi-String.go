package main

import (
    "fmt"
	"strings"
)

func main() {
	// A.44.1. Fungsi strings.Contains()
	fmt.Println("# - A.44.1. Fungsi strings.Contains()");
    var isExists = strings.Contains("john wick", "wick");
    fmt.Println(isExists);

	// A.44.2. Fungsi strings.HasPrefix()
	fmt.Println("# - A.44.2. Fungsi strings.HasPrefix()");
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false

	// A.44.3. Fungsi strings.HasSuffix()
	fmt.Println("# - A.44.3. Fungsi strings.HasSuffix()");
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false

	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true

	// A.44.4. Fungsi strings.Count()
	fmt.Println("# - A.44.4. Fungsi strings.Count()");
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2

	// A.44.5. Fungsi strings.Index()
	fmt.Println("# - A.44.5. Fungsi strings.Index()");
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2

	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2) // 4

	// A.44.6. Fungsi strings.Replace()
	fmt.Println("# - A.44.6. Fungsi strings.Replace()");
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	// A.44.7. Fungsi strings.Repeat()
	fmt.Println("# - A.44.7. Fungsi strings.Repeat()");
	var strExSeven = strings.Repeat("na", 4)
	fmt.Println(strExSeven) // "nananana"

	// A.44.8. Fungsi strings.Split()
	fmt.Println("# - A.44.8. Fungsi strings.Split()");
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1) // output: ["the", "dark", "knight"]

	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]

	// A.44.9. Fungsi strings.Join()
	fmt.Println("# - A.44.9. Fungsi strings.Join()");
	var data = []string{"banana", "papaya", "tomato"}
	var strExNine = strings.Join(data, "-")
	fmt.Println(strExNine) // "banana-papaya-tomato"

	// A.44.10. Fungsi strings.ToLower()
	fmt.Println("# - A.44.10. Fungsi strings.ToLower()");
	var strExTen = strings.ToLower("aLAy")
	fmt.Println(strExTen) // "alay"

	// A.44.11. Fungsi strings.ToUpper()
	fmt.Println("# - A.44.11. Fungsi strings.ToUpper()");
	var strExEleven = strings.ToUpper("eat!")
	fmt.Println(strExEleven) // "EAT!"
}
package main

import (
    "fmt"
	"regexp"
)

func main() {
	// A.45.1. Penerapan Regexp
	fmt.Println("# - A.45.1. Penerapan Regexp");
		var textExOne = "banana burger soup"
		var regexExOne, errExOne = regexp.Compile(`[a-z]+`)

		if errExOne != nil {
			fmt.Println(errExOne.Error())
		}

		var res1 = regexExOne.FindAllString(textExOne, 2)
		fmt.Printf("%#v \n", res1)
		// []string{"banana", "burger"}

		var res2 = regexExOne.FindAllString(textExOne, -1)
		fmt.Printf("%#v \n", res2)
		// []string{"banana", "burger", "soup"}

	// A.45.2. Method MatchString()
	fmt.Println("# - A.45.2. Method MatchString()");
		var textExTwo = "banana burger soup"
		var regexExTwo, _ = regexp.Compile(`[a-z]+`)

		var isMatch = regexExTwo.MatchString(textExTwo)
		fmt.Println(isMatch)
		// true

	// A.45.3. Method FindString()
	fmt.Println("# - A.45.3. Method FindString()");
		var textExThree = "banana burger soup"
		var regexExThree, _ = regexp.Compile(`[a-z]+`)

		var strExThree = regexExThree.FindString(textExThree)
		fmt.Println(strExThree)
		// "banana"

	// A.45.4. Method FindStringIndex()
	fmt.Println("# - A.45.4. Method FindStringIndex()");
		var textExFour = "banana burger soup"
		var regexExFour, _ = regexp.Compile(`[a-z]+`)

		var idxExFour = regexExFour.FindStringIndex(textExFour)
		fmt.Println(idxExFour)
		// [0, 6]

		var strExFour = textExFour[0:6]
		fmt.Println(strExFour)
		// "banana"

	// A.45.5. Method FindAllString()
	fmt.Println("# - A.45.5. Method FindAllString()");
		var textExFive = "banana burger soup"
		var regexExFive, _ = regexp.Compile(`[a-z]+`)

		var str1 = regexExFive.FindAllString(textExFive, -1)
		fmt.Println(str1)
		// ["banana", "burger", "soup"]

		var str2 = regexExFive.FindAllString(textExFive, 1)
		fmt.Println(str2)
		// ["banana"]

	// A.45.6. Method ReplaceAllString()
	fmt.Println("# - A.45.6. Method ReplaceAllString()");
		var textExSix = "banana burger soup"
		var regexExSix, _ = regexp.Compile(`[a-z]+`)

		var strExSix = regexExSix.ReplaceAllString(textExSix, "potato")
		fmt.Println(strExSix)
		// "potato potato potato"

	// A.45.7. Method ReplaceAllStringFunc()
	fmt.Println("# - A.45.7. Method ReplaceAllStringFunc()");
		var textExSeven = "banana burger soup"
		var regexExSeven, _ = regexp.Compile(`[a-z]+`)

		var strExSeven = regexExSeven.ReplaceAllStringFunc(textExSeven, func(each string) string {
			if each == "burger" {
				return "potato"
			}
			return each
		})
		fmt.Println(strExSeven)
		// "banana potato soup"

	// A.45.8. Method Split()
	fmt.Println("# - A.45.8. Method Split()");
		var textExEight = "banana burger soup"
		var regexExEight, _ = regexp.Compile(`[a-b]+`) // split dengan separator adalah karakter "a" dan/atau "b"

		var strExEight = regexExEight.Split(textExEight, -1)
		fmt.Printf("%#v \n", strExEight)
		// []string{"", "n", "n", " ", "urger soup"}
}
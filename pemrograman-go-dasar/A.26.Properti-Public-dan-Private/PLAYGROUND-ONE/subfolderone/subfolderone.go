package subfolderone

import "fmt"

func init() {
	ExampleThree.Name = "DEDE"
	ExampleThree.Age  = 22
}

// # - EXAMPLE ONE
var ExOne = "EXAMPLE ONE";
type ExampleOne struct {
	Name 	string
	Age		int
	Number  []int
	Abjad   []string
}
func subfolderOne(num []int) {
	fmt.Println("SUB FOLDER ONE ",num[0]);
}
func subfolderTwo(num []int) {
	fmt.Println("SUB FOLDER TWO ",num[1]);
}
func subfolderThree(num []int) {
	fmt.Println("SUB FOLDER THREE ",num[2]);
}
func ApplyAllSubFolderOne(num []int) {
	subfolderOne(num)
	subfolderTwo(num)
	subfolderThree(num)
}


// # - EXAMPLE TWO
var ExTwo = "EXAMPLE TWO";
type ExampleTwo struct {
	ExampleOne
	Name	string
	Age 	int
}


// # - EXAMPLE THREE
var ExThree = "EXAMPLE THREE";
var ExampleThree = struct {
	Name 	string
	Age		int
}{}
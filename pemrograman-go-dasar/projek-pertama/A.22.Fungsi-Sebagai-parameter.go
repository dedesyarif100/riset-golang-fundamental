package main

import (
	"fmt"
	"strings"
)

func main() {
	// A.22.1. Penerapan Fungsi Sebagai Parameter
	fmt.Println("# - A.22.1. Penerapan Fungsi Sebagai Parameter");
        var data = []string{"wick", "jason", "ethan"}
        var dataContainsO = filter(data, func(each string) bool {
            return strings.Contains(each, "o")
        })
        var dataLenght5 = filter(data, func(each string) bool {
            return len(each) == 5
        })

        fmt.Println("data asli \t\t:", data)
        // data asli : [wick jason ethan]
        fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
        // filter ada huruf "o" : [jason]
        fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
        // filter jumlah huruf "5" : [jason ethan]
        println()

    // GET MIN VALUE INTEGER
    fmt.Println("# - GET MIN VALUE INTEGER");
        var numberOne = []int{1, 2, 3, 4, 5, 6, 7, 8}
        var dataMinNumber = getMinNumber(numberOne, func(each int) bool {
            return each <= 5;
        });
        fmt.Println("MIN VALUE : ", dataMinNumber);
        println()

    // GET MAX VALUE INTEGER
    fmt.Println("# - GET MAX VALUE INTEGER");
        numberTwo := []int{1, 2, 3, 4, 5, 6, 7, 8}
        dataMaxNumber := getMaxNumber(numberTwo, func (each int) bool {
            return each > 5;
        });
        fmt.Println("MAX VALUE :", dataMaxNumber)
        println()
}

type FilterCallback func(string) bool
func filter(data []string, callback FilterCallback) []string {
	// fmt.Println(data)
    var result []string
    for _, each := range data {
        if filtered := callback(each); filtered {
			// fmt.Println(callback(each))
            result = append(result, each)
        }
    }
    return result
}


type FilterMinNumber func(int) bool
func getMinNumber(data []int, callback FilterMinNumber) []int {
    var result []int;
    for _, each := range data {
        if filtered := callback(each); filtered {
            result = append(result, each);
        }
    }
    return result;
}


type FilterMaxNumber func(int) bool
func getMaxNumber(numbers []int, callback FilterMaxNumber) []int {
    var result []int
    for _, val := range numbers {
        if filter := callback(val); filter {
            result = append(result, val)
        }
    }
    return result
}
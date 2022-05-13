package main

import (
	"fmt"
	"strings"
)

func main() {
	// A.22.1. Penerapan Fungsi Sebagai Parameter
	fmt.Println("A.22.1. Penerapan Fungsi Sebagai Parameter");
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
}

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

type FilterCallback func(string) bool
package main

import "fmt"

// func Sum(numbers []int) int {
//     var total int
//     for _, e := range numbers {
//         total += e
//     }
//     return total
// }


// A.65.2. Penerapan Generic pada Fungsi
func Sum[V int](numbers []V) V {
    var total V
    for _, val := range numbers {
        total += val
    }
    return total
}


// A.65.3. Comparable Data Type pada Fungsi Generic
func ExampleThree[V int | float32 | float64](numbers []V) V {
    var total V
    for _, val := range numbers {
        total += val
    }
    return total
}


// A.65.5. Keyword comparable
func ExampleFive[K comparable, V int64 | float64](mapping map[K]V) V {
    var data V
    for _, val := range mapping {
        data += val
    }
    return data
}


// A.65.6. Generic Type Constraint
type Number interface {
    int64 | float64
}
func ExampleSix[K comparable, V Number](mapping map[K]V) V {
    var data V
    for _, val := range mapping {
        data += val
    }
    return data
}


// A.65.7. Struct Generic
type UserModel[T int | float64] struct {
    Name string
    Scores []T
}
func (mapping *UserModel[int]) SetScoresA(scores []int) {
    mapping.Scores = scores
}
func (mapping *UserModel[float64]) SetScoresB(scores []float64) {
    mapping.Scores = scores
}


func main() {
	// A.65.2. Penerapan Generic pada Fungsi
	fmt.Println("# - A.65.2. Penerapan Generic pada Fungsi");
        total1 := Sum([]int{1, 2, 3, 4, 5})
        fmt.Println("total:", total1)
        fmt.Println();

	// A.65.3. Comparable Data Type pada Fungsi Generic
	fmt.Println("# - A.65.3. Comparable Data Type pada Fungsi Generic");
        totalExTwoSecOne := ExampleThree([]int{1, 2, 3, 4, 5})
        fmt.Println("total:", totalExTwoSecOne);

        totalExTwoSecTwo := ExampleThree([]float32{2.5, 7.2})
        fmt.Println("total:", totalExTwoSecTwo);

        totalExTwoSecThree := ExampleThree([]float64{1.23, 6.33, 12.6})
        fmt.Println("total:", totalExTwoSecThree);
        fmt.Println();

	// A.65.5. Keyword comparable
	fmt.Println("# - A.65.5. Keyword comparable");
        ints := map[string]int64{ "first": 34, "second": 12 }
        floats := map[string]float64{ "first": 35.98, "second": 26.99 }

        fmt.Printf("Generic Sums with Constraint: %v and %v\n", ExampleFive(ints), ExampleFive(floats))
        fmt.Println();

    // A.65.6. Generic Type Constraint
    fmt.Println("# - A.65.6. Generic Type Constraint");
        valA := map[string]int64{ "first": 34, "second": 12 }
        valB := map[string]float64{ "first": 35.98, "second": 26.99 }
        exampleSixSecOne := ExampleSix(valA);
        exampleSixSecTwo := ExampleSix(valB);

        fmt.Println("CEK DATA : ", exampleSixSecOne);
        fmt.Println("CEK DATA : ", exampleSixSecTwo);
        fmt.Println();

    // A.65.7. Struct Generic
    fmt.Println("# - A.65.7. Struct Generic");
        var m1 UserModel[int]
        m1.Name = "Noval"
        m1.Scores = []int{1, 2, 3}
        fmt.Println("NAME       :", m1.Name)
        fmt.Println("SCORES     :", m1.Scores)
        fmt.Println("----------------------");

        var m2 UserModel[int]
        m2.Name = "DEDE"
        m2.SetScoresA([]int{1, 2, 3})
        fmt.Println("NAME       :",m2.Name)
        fmt.Println("SCORES     :",m2.Scores)
        fmt.Println("----------------------");

        var m3 UserModel[float64] 
        m3.Name = "Noval"
        m3.SetScoresB([]float64{10.1, 11.2})
        fmt.Println("NAME       :", m3.Name)
        fmt.Println("SCORES     :", m3.Scores)
        fmt.Println();
}
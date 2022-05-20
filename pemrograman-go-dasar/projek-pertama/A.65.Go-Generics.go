package main

import "fmt"

// func Sum(numbers []int) int {
//     var total int
//     for _, e := range numbers {
//         total += e
//     }
//     return total
// }

func Sum[V int](numbers []V) V {
    var total V
    for _, e := range numbers {
        total += e
    }
    return total
}

func ExampleTwo[V int | float32 | float64](numbers []V) V {
    var total V
    for _, e := range numbers {
        total += e
    }
    return total
}

func SumNumbers3[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// A.65.6. Generic Type Constraint
type Number interface {
    int64 | float64
}

func ExampleSix[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// A.65.7. Struct Generic
type UserModel[T int | float64] struct {
    Name string
    Scores []T
}

// func (m *UserModel[int]) SetScoresA(scores []int) {
//     m.Scores = scores
// }

func (m *UserModel[float64]) SetScoresB(scores []float64) {
    m.Scores = scores
}



func main() {
	// A.65.2. Penerapan Generic pada Fungsi
	fmt.Println("# - A.65.2. Penerapan Generic pada Fungsi");
    total1 := Sum([]int{1, 2, 3, 4, 5})
    fmt.Println("total:", total1)
	fmt.Println();

	// A.65.3. Comparable Data Type pada Fungsi Generic
	fmt.Println("# - A.65.3. Comparable Data Type pada Fungsi Generic");
	totalExTwoSecOne := ExampleTwo([]int{1, 2, 3, 4, 5})
    fmt.Println("total:", totalExTwoSecOne);

	totalExTwoSecTwo := ExampleTwo([]float32{2.5, 7.2})
    fmt.Println("total:", totalExTwoSecTwo);

	totalExTwoSecThree := ExampleTwo([]float64{1.23, 6.33, 12.6})
    fmt.Println("total:", totalExTwoSecThree);
	fmt.Println();

	// A.65.5. Keyword comparable
	fmt.Println("# - A.65.5. Keyword comparable");
	ints := map[string]int64{ "first": 34, "second": 12 }
    floats := map[string]float64{ "first": 35.98, "second": 26.99 }

    fmt.Printf("Generic Sums with Constraint: %v and %v\n", SumNumbers3(ints), SumNumbers3(floats))
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
    fmt.Println("scores:", m1.Scores)

    var m2 UserModel[float64] 
    m2.Name = "Noval"
    m2.SetScoresB([]float64{10, 11})
    fmt.Println("scores:", m2.Scores)
}
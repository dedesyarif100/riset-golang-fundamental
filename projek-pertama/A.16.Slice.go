package main

import (
	"fmt"
)

func main() {
	// A.16.1. Inisialisasi Slice
	fmt.Println("A.16.1. Inisialisasi Slice");
	var exampleOne = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(exampleOne[1]) // "apple"

	var fruitsA = []string{"apple", "grape"}      // slice
	var fruitsB = [2]string{"banana", "melon"}    // array
	var fruitsC = [...]string{"papaya", "grape"}  // array
	fmt.Println(fruitsA, fruitsB, fruitsC)

	// A.16.2. Hubungan Slice Dengan Array & Operasi Slice
	fmt.Println("A.16.2. Hubungan Slice Dengan Array & Operasi Slice");
	var exampleTwo = []string{"A", "B", "C", "D"}
	fmt.Println(exampleTwo[0:2])
	fmt.Println(exampleTwo[0:4])
	fmt.Println(exampleTwo[0:0])
	fmt.Println(exampleTwo[4:4])
	// fmt.Println(exampleTwo[4:0])
	fmt.Println(exampleTwo[:])
	fmt.Println(exampleTwo[2:])
	fmt.Println(exampleTwo[:2])

	// A.16.3. Slice Merupakan Tipe Data Reference
	fmt.Println("A.16.3. Slice Merupakan Tipe Data Reference");
	var exampleThree = []string{"A", "B", "C", "D"}
	fmt.Println(exampleThree[0:3])
	fmt.Println(exampleThree[1:4])
	fmt.Println(exampleThree[1:2])
	fmt.Println(exampleThree[0:1])
	setExampleThree := exampleThree[0:1]
	setExampleThree[0] = "pinnaple"
	fmt.Println(exampleThree[0:3])
	fmt.Println(exampleThree[1:4])
	fmt.Println(exampleThree[1:2])
	fmt.Println(exampleThree[0:1])

	// A.16.4. Fungsi len()
	fmt.Println("A.16.4. Fungsi len()");
	var exampleFour = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(exampleFour)) // 4

	// A.16.5. Fungsi cap()
	fmt.Println("A.16.5. Fungsi cap()");
	var exampleFive = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(exampleFive))  // len: 4
	fmt.Println(cap(exampleFive))  // cap: 4
	fmt.Println("--------------")
	fmt.Println(len(exampleFive[0:3])) // len: 3
	fmt.Println(cap(exampleFive[0:3])) // cap: 4
	fmt.Println("--------------")
	fmt.Println(len(exampleFive[1:4])) // len: 3
	fmt.Println(cap(exampleFive[1:4])) // cap: 3

	// A.16.6. Fungsi append()
	fmt.Println("A.16.6. Fungsi append()");
	var exampleSix = []string{"A", "B", "C"}
	exampleSix = append(exampleSix, "D")
	fmt.Println(exampleSix)

	var exampleFruitsSixA = []string{"apple", "grape", "banana"}
	var exampleFruitsSixB = exampleFruitsSixA[0:2]
	fmt.Println("--------------------")
	fmt.Println(cap(exampleFruitsSixB))
	fmt.Println(len(exampleFruitsSixB))
	fmt.Println(exampleFruitsSixA)
	fmt.Println(exampleFruitsSixB)
	fmt.Println("--------------------")
	var exampleFruitsSixC = append(exampleFruitsSixB, "papaya")
	fmt.Println(exampleFruitsSixA)
	fmt.Println(exampleFruitsSixB)
	fmt.Println(exampleFruitsSixC)

	// A.16.7. Fungsi copy()
	fmt.Println("A.16.7. Fungsi copy()");
	exampleSevenA := make([]string, 3)
	srcA := []string{"A", "B", "C", "D"}
	lengthA := copy(exampleSevenA, srcA)
	fmt.Println(exampleSevenA)
	fmt.Println(srcA)
	fmt.Println(lengthA)
	fmt.Println("--------------------")
	exampleSevenB := []string{"A", "A", "A"}
	srcB := []string{"B", "C"}
	lengthB := copy(exampleSevenB, srcB)
	fmt.Println(exampleSevenB)
	fmt.Println(srcB)
	fmt.Println(lengthB)

	// A.16.8. Pengaksesan Elemen Slice Dengan 3 Indeks
	fmt.Println("A.16.8. Pengaksesan Elemen Slice Dengan 3 Indeks");
	var exampleEight = []string{"A", "B", "C"}
	var resultExampleEightA = exampleEight[0:2]
	var resultExampleEightB = exampleEight[0:2:2]

	fmt.Println(exampleEight)
	fmt.Println(len(exampleEight))
	fmt.Println(cap(exampleEight))
	fmt.Println("--------------------")
	fmt.Println(resultExampleEightA)
	fmt.Println(len(resultExampleEightA))
	fmt.Println(cap(resultExampleEightA))
	fmt.Println("--------------------")
	fmt.Println(resultExampleEightB)
	fmt.Println(len(resultExampleEightB))
	fmt.Println(cap(resultExampleEightB))
}
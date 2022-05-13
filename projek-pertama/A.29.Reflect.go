package main

import (
	"fmt"
    "reflect"
)

func main() {
	// A.29.1. Mencari Tipe Data & Value Menggunakan Reflect
	fmt.Println("# - A.29.1. Mencari Tipe Data & Value Menggunakan Reflect");
    var numberOne = 23;
    var exOne = reflect.ValueOf(numberOne);
	fmt.Println(reflect.Int);
    fmt.Println("tipe  variabel :", exOne.Type());
    if exOne.Kind() == reflect.Int {
        fmt.Println("nilai variabel :", exOne.Int());
    }

	fmt.Println("--------------------------------------------------");
	var numberTwo = 23;
	var exTwo = reflect.ValueOf(numberTwo);
	fmt.Println("tipe  variabel :", exTwo.Type());
	fmt.Println("nilai variabel :", exTwo.Interface());
	fmt.Println("nilai variabel :", exTwo.Interface().(int));
	fmt.Println();

	// A.29.2. Pengaksesan Informasi Property Variabel
	fmt.Println("# - A.29.2. Pengaksesan Informasi Property Variabel");
    var exampleTwo = &exampleTwo{Name: "wick", Grade: 2, ValA: "A", ValB: "B"}
    exampleTwo.getPropertyInfo()

	// A.29.3. Pengaksesan Informasi Method Variabel
	fmt.Println("# - A.29.3. Pengaksesan Informasi Method Variabel");

    var exampleThree = &exampleThree{Name: "john wick", Grade: 2}
    fmt.Println("nama :", exampleThree.Name)

    var reflectValue = reflect.ValueOf(exampleThree)
    var method = reflectValue.MethodByName("SetName")
    method.Call([]reflect.Value{
        reflect.ValueOf("wick"),
    })

    fmt.Println("nama :", exampleThree.Name)
}

// A.29.2. Pengaksesan Informasi Property Variabel
type exampleTwo struct {
	Name  string
	Grade int
	ValA  string
	ValB  string
}
func (s *exampleTwo) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	var reflectType = reflectValue.Type()
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

// A.29.3. Pengaksesan Informasi Method Variabel
type exampleThree struct {
	Name  string
	Grade int
}
func (s *exampleThree) SetName(name string) {
    s.Name = name
}
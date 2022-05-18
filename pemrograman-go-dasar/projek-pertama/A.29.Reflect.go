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
		fmt.Println();

	// EXAMPLE FOUR
	fmt.Println("# - EXAMPLE FOUR");
		var exFour = []int{1, 2, 3, 4, 5, 6, 7, 8};
		val := reflect.ValueOf(exFour);
		fmt.Println("CEK DATA EXAMPLE FOUR : ",val.Slice(2, 4), "",val.Type());

		var exFourSecTwo = "DEDE SYARIFUDIN";
		valExSecTwo := reflect.ValueOf(exFourSecTwo);
		setExampleFour(valExSecTwo);

		var exFourSecThree = 12;
		valExSecThree := reflect.ValueOf(exFourSecThree);
		fmt.Println(valExSecThree.Type());
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
	// fmt.Println("CEK DATA :",reflectValue.NumField());
	fmt.Println("CEK DATA :",reflect.Ptr);
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

func setExampleFour(v any) {
	fmt.Println("CEK DATA EXAMPLE FOUR : ",v);
}

// A.29.3. Pengaksesan Informasi Method Variabel
type exampleThree struct {
	Name  string
	Grade int
}
func (s *exampleThree) SetName(name string) {
    s.Name = name
}
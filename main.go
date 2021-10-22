package main

import (
	"fmt"
)

var nilaiPi = 3.14

type Person struct {
	Name string
	age  int
}

func tambahFunc() func(int, int) int {
	return func(a, b int) int {
		return a + b
	}
}

func main() {

	hasilTambah := tambahFunc()(1, 6)

	nilaiPi = 10

	fmt.Println("hasil", hasilTambah, nilaiPi)

	// var greet string = "Hello Redant!"

	// greet := "Idam"
	// sayHello(greet)

	// c := printers.NewChild("Umar", 18)
	// fmt.Println("Child name:", c.Name)

	// var people []Person = []Person{
	// 	{
	// 		Name: "Idam",
	// 		age:  1,
	// 	},
	// 	{
	// 		Name: `"Amri"
	// 		123`,
	// 	},
	// }

	// for _, p := range people {
	// 	fmt.Println(p.Name)
	// }
}

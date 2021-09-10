package main

import "fmt"

func main() {
	var name1 = "Johar"
	var name2 = "Johar"

	var result = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2) // false
	fmt.Println(value1 < value2) // true
	fmt.Println(value1 == value2) // false
	fmt.Println(value1 != value2) // true
}
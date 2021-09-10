package main

import "fmt"

func main() {
	var names [3] string

	names[0] = "Jauhar"
	names[1] = "Juned"
	names[2] = "Agung"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var value = [3]int {10, 20, 30}

	fmt.Println(value)
	fmt.Println(value[0])
	fmt.Println(value[1])
	fmt.Println(value[2])

	fmt.Println(len(value))
	fmt.Println(len(names))

	var nominal [10] int
	fmt.Println(len(nominal))
}
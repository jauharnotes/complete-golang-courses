package main

import "fmt"

func main() {
	var fruits = []string{"apple", "orange", "banana"}
	var bFruit = fruits[0:2]

	fmt.Println(cap(bFruit))
	fmt.Println(len(bFruit))
	
	fmt.Println(fruits)
	fmt.Println(bFruit)

	var cFruit = append(fruits, "papaya")

	fmt.Println(fruits)
	fmt.Println(bFruit)
	fmt.Println(cFruit)

	// funngsi copy()
	dst := make([]string, 3)
	src := []string{"watermelon", "pipnnaplel", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}
package main

import (
	"basic/calculation"
	"fmt"
)

func main() {
	fmt.Println("Halo, yuk belajar golang")

	result := calculation.Add(10, 20)
	fmt.Println(result)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(7.3))
	fmt.Println(math.Round(7.8))
	fmt.Println(math.Floor(7.3))
	fmt.Println(math.Floor(7.8))

	fmt.Println(math.Ceil(7.3))
	fmt.Println(math.Ceil(7.8))
	fmt.Println(math.Min(10, 3))
	fmt.Println(math.Max(10, 3))
}
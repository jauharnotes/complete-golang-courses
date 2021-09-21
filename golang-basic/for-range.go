package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "orange", "banana", "cipluan"}

	for i, fruit := range fruits {
		fmt.Println("element :", i, fruit)
	}

	// Penggunaan Variabel Underscore _ Dalam for - range
	var fruits2 = [4]string{"apple", "orange", "banana", "cipluan"}

	for _, fruit := range fruits2 {
		fmt.Println("nama buah :", fruit)
	}
}
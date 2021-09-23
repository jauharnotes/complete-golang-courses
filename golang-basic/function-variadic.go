package main

import (
	"fmt"
	"strings"
)

func sumAll(number ...int) int {
	total := 0
	for _,value := range number {
		total += value
	}
	return total
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("hello, my name is: %s\n", name)
	fmt.Printf("my hobbies are: %s\n", hobbiesAsString)
}

func main() {
	total := sumAll(10,10,10,10)
	fmt.Println(total)

	slice := []int{10,20,30,40}
	total = sumAll(slice...)
	fmt.Println(total)

	// Pengisian Parameter Fungsi Variadic Menggunakan Data Slice
	var numbers = []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	var avg = sumAll(numbers...)
	fmt.Println("Rata-rata :", avg)

	yourHobbies("johar", "reading", "traveler")
}
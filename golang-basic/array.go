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

	// inisialisasi nilai awal array
	var value = [3]int {10, 20, 30}

	fmt.Println(value)
	fmt.Println(value[0])
	fmt.Println(value[1])
	fmt.Println(value[2])

	fmt.Println(len(value))
	fmt.Println(len(names))

	var nominal [10] int
	fmt.Println(len(nominal))

	// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	var number = [...]int {2, 3, 4, 5}

	fmt.Println("data array \t\t:", number)
	fmt.Println("jumlah elelmen \t:", len(number))

	// array multidimensi
	var number1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var number2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("Number1", number1)
	fmt.Println("Number2", number2)

	// perulangan element array menggunakan keyword for
	var fruits = [4]string{"orange", "banana", "apple", "cipluan"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println("Element : ", i, fruits[i])
	}
}
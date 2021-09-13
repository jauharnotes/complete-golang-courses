package main

import "fmt"


func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// for with statment
	for counter1 := 1; counter1 <= 10; counter1++ {
		fmt.Println("Perulangan ke ", counter1)
	}

	slice := []string{"jauhar", "juned", "agung"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, value := range slice {
		fmt.Println("Index", index, "=", value)
	}

	person := make(map[string]string)

	person["name"] = "johar"
	person["title"] = "programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
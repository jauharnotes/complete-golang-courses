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

	// for argument only conditions
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	// fro without argument
	var foo = 0

	for {
		fmt.Println("Angka", foo)

		foo++
		if foo == 5 {
			break
		}
	}

	// keyword break & countinue
	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// memanfaatkan label dalam perulangan
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
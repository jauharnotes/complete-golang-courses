package main

import "fmt"

func main() {
	name := "johar"
	counter := 0

	increment := func(){
		name := "agung"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}
	increment()
	fmt.Println(name)
}


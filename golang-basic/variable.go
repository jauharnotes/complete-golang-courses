package main

import "fmt"

func main(){
	var name string

	name = "Jauhar"
	fmt.Println(name)

	name = "Juned"
	fmt.Println(name)

	var friendName = "Agung"
	fmt.Println(friendName)

	var age = 22
	fmt.Println(age)

	county := "Indonesia"
	fmt.Println(county)

	// multiple variable
	var (
		lastName = "Jauhar"
		firstName = "Uddin"
	)

	fmt.Println(lastName)
	fmt.Println(firstName)

	// Deklarasi multi Variable
	var first, secon, third string
	first, secon, third = "one", "two", "three"

	fmt.Println(first, secon, third)
}
package main

import "fmt"

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "Johar"
	middleName = "Uddin"
	lastName = "Batubara"

	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a, b, c)
}


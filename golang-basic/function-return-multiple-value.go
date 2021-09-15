package main

import "fmt"

func getFullName() (string, string, string) {
	return "Johar", "Uddin", "Batubara"
}

func main() {
	firsName, lastName,_ := getFullName()
	fmt.Println(firsName, lastName)
}
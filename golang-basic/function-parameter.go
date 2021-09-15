package main

import "fmt"

func sayHello(firsName string, lastName string){
	fmt.Println("Hello", firsName, lastName)
}

func main() {
	sayHello("Johar", "uddin")
	sayHello("agung", "podomoro")
}
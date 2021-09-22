package main

import (
	"fmt"
	"strings"
)

func sayHello(){
	fmt.Println("Hello johar")
}

func main() {
	var names = []string{"johar", "uddin"}
	printMessages("halo", names)
	// sayHello();
	// sayHello();
	// sayHello();

	for i := 0; i < 10; i++ {
		sayHello();
	}
}

func printMessages(messege string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(messege, nameString)
}
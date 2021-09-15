package main

import "fmt"

func sayHello(){
	fmt.Println("Hello johar")
}

func main() {
	// sayHello();
	// sayHello();
	// sayHello();

	for i := 0; i < 10; i++ {
		sayHello();
	}
}
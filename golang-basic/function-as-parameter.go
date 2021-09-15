package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string{
	if name == "Jancuk" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Johar", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Jancuk", filter)
}
package main

import "fmt"

type Blaclist func(string) bool

func registerUser(name string, blaclist Blaclist) {
	if blaclist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blaclist := func(name string) bool {
		return name == "admin"
	}

	registerUser("johar", blaclist)
	registerUser("admin", blaclist)
	registerUser("anjing", func(name string)bool {
		return name == "anjing"
	})
}
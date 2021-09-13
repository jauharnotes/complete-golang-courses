package main

import "fmt"

func main() {
	var name = "joko anwar"

	if name == "johar" {
		fmt.Println("Hello " + name)
	} else if name == "juned" {
		fmt.Println("Hello " + name)
	} else {
		fmt.Println("Hi, Kenalan yuk!")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
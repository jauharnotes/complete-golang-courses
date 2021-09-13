package main

import "fmt"

func main() {
	name := "joharuddinn"

	switch name {
	case "johar":
		fmt.Println("Hello " + name)
		fmt.Println("Apakabar")
	case "joko":
		fmt.Println("Hello " + name)
		fmt.Println("Apakabar")
	default:
		fmt.Println("Kenalan yuk!")
	}

	switch lenght := len(name); lenght > 4 {
	case true:
		fmt.Println("Nama sudah benar")
	case false:
		fmt.Println("Nama terlalu panjang")
	}

	lenght2 := len(name)
	switch {
	case lenght2 > 10:
		fmt.Println("Nama terlalu panjanag")
	case lenght2 > 6:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
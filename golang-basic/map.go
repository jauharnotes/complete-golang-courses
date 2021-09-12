package main

import "fmt"

func main() {
	persons := map[string]string {
		"name": "Johar",
		"address": "Depok",
	}

	persons["Title"] = "Programmer"
	
	fmt.Println(persons)
	fmt.Println(persons["name"])
	fmt.Println(persons["address"])
	fmt.Println(persons["Title"])

	var book = make(map[string]string)
	book["Title"] = "Belajar Golang"
	book["Audhor"] = "Johar"
	book["Terbit"] = "2021"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "Terbit")

	fmt.Println(book)
	fmt.Println(len(book))
}